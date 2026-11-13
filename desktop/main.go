// smart Gantt (sGantt) desktop — wraps the offline index.html in a native Windows window (WebView2).
// The page is served only on the loopback interface (127.0.0.1); nothing is sent to the internet.
package main

import (
	_ "embed"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/jchv/go-webview2"
)

//go:embed web_index.html
var page []byte

func listen() net.Listener {
	// A fixed port keeps the browser storage (the autosaved plan) stable between runs; fall back if busy.
	if l, err := net.Listen("tcp", "127.0.0.1:47392"); err == nil {
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
	c, _ := syscall.UTF16PtrFromString("smart Gantt")
	user32.NewProc("MessageBoxW").Call(0, uintptr(unsafe.Pointer(t)), uintptr(unsafe.Pointer(c)), 0x40)
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

	dataDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "sGantt")
	_ = os.MkdirAll(dataDir, 0o755)

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     false,
		AutoFocus: true,
		DataPath:  dataDir,
		WindowOptions: webview2.WindowOptions{
			Title:  "smart Gantt (sGantt) 1.0",
			Width:  1440,
			Height: 900,
			IconId: 2,
			Center: true,
		},
	})
	if w == nil {
		msg("smart Gantt needs the Microsoft Edge WebView2 Runtime.\n\nInstall it from:\nhttps://developer.microsoft.com/microsoft-edge/webview2/")
		return
	}
	defer w.Destroy()
	w.SetSize(1100, 700, webview2.HintMin)
	w.SetSize(1440, 900, webview2.HintNone)
	w.Navigate("http://" + ln.Addr().String() + "/")
	w.Run()
}
