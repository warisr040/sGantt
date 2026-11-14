# sTrack (simply handle IT Track) — release notes

## Version 2.3
- **Shortcut keys on every page.** The most used actions work with the left hand only: **Ctrl+1 … Ctrl+5** do the main buttons of the page you are on (the same keys mean different things on different pages). **Alt+1 … Alt+6** switch pages. Hover any button to see its key; press **F1** for the full list.
  - Gantt: Ctrl+1 task, Ctrl+2 milestone, Ctrl+3 event, Ctrl+4 details, Ctrl+5 capture.
  - Board: Ctrl+1 To do, Ctrl+2 In progress, Ctrl+3 Review, Ctrl+4 Blocked, Ctrl+5 Done, Ctrl+6 On hold (moves the selected cards).
  - Sprints: Ctrl+1 new sprint, Ctrl+2 add tasks, Ctrl+3/4 previous/next sprint, Ctrl+5 whole project.
  - Timesheet: Ctrl+1/2/3 previous/this/next week, Ctrl+4 add task, Ctrl+5 today's cell.
  - Team: Ctrl+1 add member. Reports: Ctrl+1 copy, Ctrl+2 save.
  - Everywhere: Alt+Q names 1/2/both, Alt+W database, Alt+A Sutora, Alt+S settings, Alt+T theme, Alt+D customer view, Alt+X lock, Alt+N new project, Ctrl+S/O save/open, Ctrl+Z/Y undo/redo. In dialogs Ctrl+Enter confirms, Esc closes.
- **Select several tasks** with Ctrl+click or Shift+click (Ctrl+A for all, Esc to clear). A batch bar appears to set status, PICs, sprint, type, days, MH, done % or start for all of them at once, move them N working days earlier or later, indent/outdent, or delete. On the Board, click to select cards (double-click opens details) and drag or press Ctrl+1…6 to move them together.

## Version 2.2
- Sutora can now look like the **Donut girl**, shown as her full picture (no crop, about 96 px tall, background removed) (choose with ⚙ on her card, or in Theme). Classic stays the default. She uses the calm pose when idle and the pointing pose while giving a tip. Size 1×, 1.5× or 2× (⚙ on her card, or Theme). Your choices are remembered on your computer.
- In Customer view or while locked, Team, Board and Timesheet now say why editing is off and offer "Back to editing" (or Unlock); Sutora says so too instead of asking for edits. The empty team list has its own Add member button.
- Hide Sutora with the small × at her bottom right; bring her back with the small S in the corner (or Theme → Show Sutora).
- Fix: text typed in the guided setup is no longer lost when Sutora refreshes her card.

## Version 2.1
- **Sutora**, the guide (bottom right). Click her to turn her on or off. She suggests what to fill in next and what is missing, with a "Show me" button that jumps to the place. Viewpoints: Auto (default), Setup, Project manager, Resources, My work, Agile, Before sharing. Speaks English, ไทย or 日本語.
- **New project** asks: start with Sutora's guide, or start empty. The guide has 8 steps in the order projects really start (basics, calendar, team, scope, schedule, owners and effort, sprints, baseline and sharing). Stop at any moment.
- **Templates**: V-model (ASPICE-style, tests linked to their design steps, optional HIL and vehicle tests), Waterfall, Iterative and Blank. Change everything before inserting: tick or untick, rename in both languages, add steps inside any phase, days, MH, gates. Save your own templates; they are shared through the team database.

## Version 2.0
Renamed from smart Gantt (sGantt) to **sTrack**. New logo.

### Team work
- **Team database**: one plan file on a shared drive or server folder. Desktop app: connect by path (e.g. `\\fileserver\PM\Project.strack.json`) and share the path; browser (Edge/Chrome): pick the file. Syncs every few seconds.
- **Field-by-field merge**: two people can edit different fields of the same task at the same time; if both change the same field, the newer value wins and the other person is told.
- **Who is editing**: rows and cells being edited by someone else are highlighted in that person's colour.
- **Lock**: type `lock` to freeze the plan (and `unlock` to release it). Edits made by others after the lock are kept aside on their computer — copy them as text, save them as a file, or apply them again after unlocking.
- **Team sheet**: names in two languages, role, contact, MH per day, open tasks, remaining MH, hours logged this week.
- **Timesheet**: each PIC logs man-hours per day on their tasks (several PICs per task never overwrite each other) and can update progress.

### Planning
- Several PICs per task; Status column (To do, In progress, Review, Blocked, On hold, Done); Planning state for tasks without a start date (no bar).
- **MH** (planned man-hours) is independent from Days — waiting tasks can be long with few MH.
- Children start from their parent's start date when they have no own start or predecessor.
- **Events**: star, pin or important marks, next to Task and Milestone.
- **Two names per task** (e.g. Japanese and Thai) with a switch: first only, second only, or both.
- **Baselines**: save the plan, see it as a grey bar under each task and a "vs baseline" column in working days.
- **Customer view**: one name language, no man-hours, read only.
- Columns: drag to reorder, drag edges to resize, right-click to choose. Resizable table/chart splitter.
- Shortcuts: Alt+arrows to indent, outdent and move; F2 task details; right-click menu on rows.

### Agile
- **Board**: Kanban by status with drag and drop, filter by sprint and PIC.
- **Sprints**: dates, goal, backlog and a burndown linked to progress (ideal, actual, forecast).

### Output
- **Capture to clipboard**: whole Gantt, a date range, or draw an area like the Snipping Tool (Ctrl+Shift+S). Save as PNG too.
- **Simple report**: tick the tasks, get a clean one-page view to copy.
- **Smart report**: progress against plan, schedule index, effort, needs-attention list, next 14 days, workload per person, status mix.
- **Excel / CSV**: export Tasks, Team and Timesheet to .xlsx, tasks to CSV; import from .xlsx or .csv (people and sprints are created when missing).

## Version 1.0
First release as smart Gantt: Gantt with 8 levels, working calendar and holidays, predecessors, Plan / Actual / Plan vs Actual, 5 themes, PNG export.
