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
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"syscall"
	"time"
	"unicode/utf16"
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
func fail(err error) string       { return out(map[string]any{"ok": false, "err": err.Error()}) }

// fingerprint changes whenever the file is rewritten (time + size), so pollers notice other people's saves.
func fingerprint(st os.FileInfo) string {
	return fmt.Sprintf("%d-%d", st.ModTime().UnixMilli(), st.Size())
}

// ---- native "choose file" dialog for the team database (comdlg32, no cgo) ----
type openFileNameW struct {
	lStructSize       uint32
	hwndOwner         uintptr
	hInstance         uintptr
	lpstrFilter       *uint16
	lpstrCustomFilter *uint16
	nMaxCustFilter    uint32
	nFilterIndex      uint32
	lpstrFile         *uint16
	nMaxFile          uint32
	lpstrFileTitle    *uint16
	nMaxFileTitle     uint32
	lpstrInitialDir   *uint16
	lpstrTitle        *uint16
	Flags             uint32
	nFileOffset       uint16
	nFileExtension    uint16
	lpstrDefExt       *uint16
	lCustData         uintptr
	lpfnHook          uintptr
	lpTemplateName    *uint16
	pvReserved        uintptr
	dwReserved        uint32
	FlagsEx           uint32
}

func u16(s string) []uint16 { return append(utf16.Encode([]rune(s)), 0) }

// a mapped drive (Z:\...) becomes \\server\share\... so every PC can use the same location
func toUNC(p string) string {
	if len(p) < 3 || p[1] != ':' {
		return p
	}
	proc := syscall.NewLazyDLL("mpr.dll").NewProc("WNetGetConnectionW")
	local := u16(p[:2])
	buf := make([]uint16, 1024)
	n := uint32(len(buf))
	r, _, _ := proc.Call(uintptr(unsafe.Pointer(&local[0])), uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&n)))
	if r != 0 {
		return p // a local disk, or not connected
	}
	return syscall.UTF16ToString(buf) + p[2:]
}

func pickFile(owner uintptr, start, title string) (string, bool) {
	buf := make([]uint16, 4096)
	if start != "" {
		copy(buf[:len(buf)-1], utf16.Encode([]rune(start)))
	}
	filter := u16("sTrack database (*.json)\x00*.json\x00All files (*.*)\x00*.*\x00")
	ttl := u16(title)
	ext := u16("json")
	ofn := openFileNameW{
		hwndOwner:    owner,
		lpstrFilter:  &filter[0],
		nFilterIndex: 1,
		lpstrFile:    &buf[0],
		nMaxFile:     uint32(len(buf)),
		lpstrTitle:   &ttl[0],
		lpstrDefExt:  &ext[0],
		// OFN_EXPLORER | OFN_PATHMUSTEXIST | OFN_NOCHANGEDIR | OFN_HIDEREADONLY (no FILEMUSTEXIST: a new name may be typed)
		Flags: 0x00080000 | 0x00000800 | 0x00000008 | 0x00000004,
	}
	ofn.lStructSize = uint32(unsafe.Sizeof(ofn))
	r, _, _ := syscall.NewLazyDLL("comdlg32.dll").NewProc("GetOpenFileNameW").Call(uintptr(unsafe.Pointer(&ofn)))
	if r == 0 {
		return "", false
	}
	return toUNC(syscall.UTF16ToString(buf)), true
}

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
			Title:  "sTrack 3.4 — simply handle IT Track",
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
	_ = w.Bind("stPickFile", func(start string) string {
		p, ok := pickFile(uintptr(w.Window()), start, "Choose the team database file (pick one, or type a new name)")
		if !ok {
			return out(map[string]any{"ok": false, "cancel": true})
		}
		_, err := os.Stat(p)
		return out(map[string]any{"ok": true, "path": p, "exists": err == nil})
	})
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

	// Classic Outlook on this PC: read calendar items (start, end, busy status, all-day, response) between two dates.
	// Uses Outlook's own automation through PowerShell. Nothing leaves the computer; titles are not read.
	_ = w.Bind("stOutlook", func(from, to string) string {
		day := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
		if !day.MatchString(from) || !day.MatchString(to) {
			return out(map[string]any{"ok": false, "err": "bad dates"})
		}
		ps := `$ErrorActionPreference='Stop';
try { $ol = New-Object -ComObject Outlook.Application } catch { Write-Output '{"ok":false,"err":"noOutlook"}'; exit }
$ns = $ol.GetNamespace('MAPI'); $cal = $ns.GetDefaultFolder(9)
$a = [datetime]::ParseExact('` + from + `','yyyy-MM-dd',$null); $b = [datetime]::ParseExact('` + to + `','yyyy-MM-dd',$null).AddDays(1)
$items = $cal.Items; $items.IncludeRecurrences = $true; $items.Sort('[Start]')
$r = $items.Restrict("[Start] < '" + $b.ToString('g') + "' AND [End] > '" + $a.ToString('g') + "'")
$list = @(); foreach ($i in $r) { if ($i.Start -ge $b) { break }; $list += [pscustomobject]@{ s = $i.Start.ToString('yyyy-MM-ddTHH:mm'); e = $i.End.ToString('yyyy-MM-ddTHH:mm'); b = [int]$i.BusyStatus; a = [bool]$i.AllDayEvent; r = [int]$i.ResponseStatus } }
@{ ok = $true; items = $list } | ConvertTo-Json -Depth 3 -Compress`
		cmd := exec.Command("powershell.exe", "-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-Command", ps)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000}
		b, err := cmd.Output()
		if err != nil {
			return fail(err)
		}
		return string(b)
	})

	w.SetSize(1100, 700, webview2.HintMin)
	w.SetSize(1480, 920, webview2.HintNone)
	w.Navigate("http://" + ln.Addr().String() + "/")
	w.Run()
}
