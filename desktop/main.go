// sTrack (simply handle IT Track) desktop — wraps the offline index.html in a native Windows window (WebView2).
// The page is served only on the loopback interface (127.0.0.1); nothing is sent to the internet.
// The team database is a plain JSON file at a path the team shares (e.g. \\server\share\project.strack.json);
// the page reads and writes it through the st* functions below, with a .lock file guarding each merge-write.
package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
	"unsafe"

	"github.com/jchv/go-webview2"
)

//go:embed web_index.html
var page []byte

func listen() net.Listener {
	// A fixed port keeps browser storage (the local copy of the plan) stable between runs; fall back if busy.
	if l, err := net.Listen("tcp", "127.0.0.1:47393"); err == nil {
		return l
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	return l
}

func msg(text string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	t, _ := syscall.UTF16PtrFromString(text)
	c, _ := syscall.UTF16PtrFromString("sTrack")
	user32.NewProc("MessageBoxW").Call(0, uintptr(unsafe.Pointer(t)), uintptr(unsafe.Pointer(c)), 0x40)
}

func out(v map[string]any) string { b, _ := json.Marshal(v); return string(b) }
func fail(err error) string        { return out(map[string]any{"ok": false, "err": err.Error()}) }

// fingerprint changes whenever the file is rewritten (time + size), so pollers notice other people's saves.
func fingerprint(st os.FileInfo) string { return fmt.Sprintf("%d-%d", st.ModTime().UnixMilli(), st.Size()) }

func main() {
	ln := listen()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: blob:; connect-src 'none'")
		w.Write(page)
	})
	go http.Serve(ln, mux)

	dataDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "sTrack")
	_ = os.MkdirAll(dataDir, 0o755)

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     false,
		AutoFocus: true,
		DataPath:  dataDir,
		WindowOptions: webview2.WindowOptions{
			Title:  "sTrack 2.0 — simply handle IT Track",
			Width:  1480,
			Height: 920,
			IconId: 2,
			Center: true,
		},
	})
	if w == nil {
		msg("sTrack needs the Microsoft Edge WebView2 Runtime.\n\nInstall it from:\nhttps://developer.microsoft.com/microsoft-edge/webview2/")
		return
	}
	defer w.Destroy()

	// ---- team database on a shared path ----
	_ = w.Bind("stDbStat", func(p string) string {
		st, err := os.Stat(p)
		if err != nil {
			return fail(err)
		}
		return out(map[string]any{"ok": true, "mtime": fingerprint(st)})
	})
	_ = w.Bind("stDbRead", func(p string) string {
		b, err := os.ReadFile(p)
		if err != nil {
			return fail(err)
		}
		st, err := os.Stat(p)
		if err != nil {
			return fail(err)
		}
		return out(map[string]any{"ok": true, "text": string(b), "mtime": fingerprint(st)})
	})
	_ = w.Bind("stDbWrite", func(p, text string) string {
		tmp := p + ".tmp-" + strconv.Itoa(os.Getpid())
		if err := os.WriteFile(tmp, []byte(text), 0o644); err != nil {
			return fail(err)
		}
		if err := os.Rename(tmp, p); err != nil { // replaces the old file in one step
			_ = os.Remove(tmp)
			return fail(err)
		}
		st, err := os.Stat(p)
		if err != nil {
			return fail(err)
		}
		return out(map[string]any{"ok": true, "mtime": fingerprint(st)})
	})
	_ = w.Bind("stDbLock", func(p string) string {
		lock := p + ".lock"
		host, _ := os.Hostname()
		deadline := time.Now().Add(6 * time.Second)
		for {
			f, err := os.OpenFile(lock, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
			if err == nil {
				fmt.Fprintf(f, "%s %d %s", host, os.Getpid(), time.Now().Format(time.RFC3339))
				f.Close()
				return out(map[string]any{"ok": true})
			}
			if st, e := os.Stat(lock); e == nil && time.Since(st.ModTime()) > 30*time.Second {
				_ = os.Remove(lock) // left behind by a computer that crashed
				continue
			}
			if !os.IsExist(err) {
				return fail(err)
			}
			if time.Now().After(deadline) {
				return out(map[string]any{"ok": false, "err": "the database is busy saving from another computer; it will retry"})
			}
			time.Sleep(150 * time.Millisecond)
		}
	})
	_ = w.Bind("stDbUnlock", func(p string) string {
		_ = os.Remove(p + ".lock")
		return out(map[string]any{"ok": true})
	})

	w.SetSize(1100, 700, webview2.HintMin)
	w.SetSize(1480, 920, webview2.HintNone)
	w.Navigate("http://" + ln.Addr().String() + "/")
	w.Run()
}
