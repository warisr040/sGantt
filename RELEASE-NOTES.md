# sTrack (simply handle IT Track) — release notes

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
