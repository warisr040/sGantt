# sTrack — static site, no build step

Single `index.html`, no dependencies. Nothing is uploaded — the plan lives in your browser, in a saved `.strack.json` file, or in a team database file on a shared folder (Edge or Chrome).

## Deploy to Netlify
- **Drag & drop:** open https://app.netlify.com/drop and drop this `web` folder.
- **Git:** "Add new site → Import" this repo, leave the build command empty, base directory `web`, publish directory `web`.

`netlify.toml` adds a Content-Security-Policy with `connect-src 'none'`, so the page cannot send data anywhere.
