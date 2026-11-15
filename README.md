# sTrack — simply handle IT Track · version 2.7.2

Offline project tracking for small teams: Gantt, agile board, sprints with burndown, timesheet, team sheet and reports.
Nothing is uploaded. A team can share one database file on a network drive or server folder.
Same look and themes as [sUML](https://github.com/warisr040/sUML). (This repository started as sGantt.)

| Folder | What |
| --- | --- |
| `web/` | The whole app: `index.html` (single file, no dependencies) + `netlify.toml` for Netlify hosting |
| `desktop/` | Windows exe wrapper (Go + WebView2). Adds the shared-path team database. `web_index.html` is a copy of `web/index.html` embedded at build time — see `BUILD.md` |
| `sutora/` | Original artwork for Sutora’s second look (the app embeds small cropped copies) |
| `logo/` | App logo as SVG and PNG (16–512 px); `desktop/icon.ico` is the Windows icon |

## Keyboard
Left-hand shortcuts: **Ctrl+1…5** run the main buttons of the current page, **Alt+1…6** switch pages, **F1** lists everything. Arrow keys move between rows, columns and Board cards (not while the Japanese IME is converting). Ctrl/Shift+click selects several tasks for batch edits.

Edit `web/index.html`, then copy it to `desktop/web_index.html` before rebuilding the exe.
