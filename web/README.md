# sGantt — static site, no build step

Single `index.html`, no dependencies. Nothing is uploaded — plans live in your browser until you **Save** (`.sgantt.json`) or **Export PNG**.

## Deploy to Netlify
- **Drag & drop:** open https://app.netlify.com/drop and drop this `web` folder.
- **Git:** "Add new site → Import" this repo, leave the build command empty, base directory `web`, publish directory `web`.

`netlify.toml` adds a Content-Security-Policy with `connect-src 'none'`, so the page cannot send data anywhere.
