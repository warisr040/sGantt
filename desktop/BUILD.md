# Building sGantt.exe (smart Gantt desktop)

The exe is a ~6.5 MB native Windows window (Microsoft Edge WebView2) around the same `index.html`
used for the website. It serves the page only on 127.0.0.1 — nothing goes to the internet.
The autosaved plan lives in `%LOCALAPPDATA%\sGantt`. Requires WebView2 Runtime (built into Windows 10/11).

To rebuild after editing the site:
1. Copy `web/index.html` here as `web_index.html`.
2. Install Go (https://go.dev). Resources (icon, manifest, version info 1.0) are in
   `rsrc_windows_amd64.syso`. To regenerate it after changing `versioninfo.json`, `icon.ico` or `app.manifest`:
       go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
       goversioninfo -64 -o rsrc_windows_amd64.syso versioninfo.json
3. Build:
       go build -trimpath -ldflags "-s -w -H windowsgui" -o sGantt.exe .
   (From Linux/macOS add: GOOS=windows GOARCH=amd64 CGO_ENABLED=0)
