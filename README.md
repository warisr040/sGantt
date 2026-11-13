# sTrack — simply handle IT Track · version 2.1

Offline project tracking for small teams: Gantt, agile board, sprints with burndown, timesheet, team sheet and reports.
Nothing is uploaded. A team can share one database file on a network drive or server folder.
Same look and themes as [sUML](https://github.com/warisr040/sUML). (This repository started as sGantt.)

| Folder | What |
| --- | --- |
| `web/` | The whole app: `index.html` (single file, no dependencies) + `netlify.toml` for Netlify hosting |
| `desktop/` | Windows exe wrapper (Go + WebView2). Adds the shared-path team database. `web_index.html` is a copy of `web/index.html` embedded at build time — see `BUILD.md` |
| `logo/` | App logo as SVG and PNG (16–512 px); `desktop/icon.ico` is the Windows icon |

Edit `web/index.html`, then copy it to `desktop/web_index.html` before rebuilding the exe.
