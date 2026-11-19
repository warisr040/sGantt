# sTrack (simply handle IT Track) — release notes

## Version 3.1
- **Sprint rows on the Gantt** (replace the Sprint column): click **Sprint ▾** next to Milestone (or Ctrl+6) and pick a sprint from the Sprints tab, or "New sprint".
  - The row shows the sprint's name, dates and goal, and its bar spans the sprint.
  - Tasks under a sprint row belong to that sprint. Tasks already assigned to the sprint move under its row when you add it.
  - Select the sprint row and press Insert to add a task under it; a task with no start date begins on the sprint's start.
  - Assigning a task to a sprint anywhere else (task details, "Add tasks" in the Sprints tab, batch edit, templates) moves it under that sprint's row; removing it moves it out.
  - The sprint row's **Start** is editable: type a date or use ▲ ▼ (one working day). The whole sprint moves and keeps its length; tasks under it follow.
  - Deleting a sprint keeps its tasks (one level up). Double-clicking a sprint row opens it in the Sprints tab.
- The guide is now called **アシス・トラ**.
- Her guided setup has a new **Roles** step before Team: the default roles are filled in; add more from a list of common team roles or type your own. In the Team step, the role is a drop-down.
- **Drag bars on the Gantt chart**: drag a bar to move its start date, drag its right end to change its length (working days). Milestones and events move too. In Plan vs Actual, the upper bar is the plan and the lower bar is the actual dates; in the Actual view you move actual dates. Moves snap to working days, a tip shows the new dates, and Ctrl+Z undoes. Group and sprint rows follow their tasks.
- **Milestones and events: ▲ ▼ buttons** in the Days column move the date one working day later or earlier (the cell shows the date), and in the MH column add or take away 1 MH.
- **Backups** (Alt+K, or Open ▸ Back up now / Backups…, or the Team database window): a full copy of the plan (tasks, team, sprints, timesheet, baselines…) saved with a name. When connected, backups are stored in the team database, so everyone can restore them; the newest 20 are kept. **Restore** first backs up the current plan, then replaces it (for everyone connected). **Save as file** exports a backup as a plan file.
- Connecting to a database replaces the plan on this computer, but your previous local plan is no longer lost: it is saved as a "Before connecting" backup in the team database.
- **Opening a plan file or starting a new project disconnects you from the team database** instead of replacing the team plan. Changes you had not synced yet are saved to the team database first; the opened or new plan stays on this computer only, and the team plan is unchanged. To share it, connect it again with Team database (Alt+W).
- **When a PIC can't be assigned, アシス・トラ says why** and offers the fix: the row has sub-tasks (its PIC comes from them; shown as "↳ sub-tasks"), it is a sprint row, Customer view is on, the plan is locked, there are no team members, or everyone is inactive. If she is hidden or off, the reason shows as a message instead.
- Typing task names: **Enter** saves and moves to the next row (no new task); **Shift+Enter** saves and adds a new task below. Shift+Enter on a selected row also adds a task below it.
- Multi-select: a plain click selects only that row and clears the old selection, and becomes the start for Shift+click; Shift+click selects exactly the rows between (also across sprint rows); on the Board it follows the card order on screen. Batch "Sprint…" leaves sprint rows alone.
- Team: **Add role** is a drop-down of common team roles (plus "Custom role…"). A member's Role list also offers the common roles, so you can add and pick one in one go.

## Version 3.0
- Milestone release: everything from 2.x in one stable version — Gantt with Plan vs Actual, agile Board and sprints, Timesheet (Actual / Burndown MH), reports, team roles and availability (built-in Thai/Japanese holidays, personal schedules, Outlook/.ics import), templates, and the Sutora guide.
- Same features and file format as 2.7.5; files saved by 2.x open unchanged.
- A User Manual with screenshots and the full hotkey list is available.

## Version 2.7.5
- The full sample's project manager is now **John Smith** (จอห์น สมิธ). A saved copy of the sample that still shows the old name is updated when it opens.

## Version 2.7.4
- **Easy sample** (New project → "Open the easy sample"): a small *To-do List App* project in English and Thai. Four people (owner, team leader, two PICs), five short steps (Plan, Design, Build, Test, Launch), each ending with a small milestone M1–M5. Today always falls in the Build step, so you see finished work, work in progress and work still to come. The full ECU sample is still there ("Open the full sample").
- **Simple app template** (Template → Simple app): the same five steps with a milestone after each, in English, Japanese and Thai, ready to insert into your own plan.

## Version 2.7.3
- **Timesheet MH mode**, chosen when you create a project (and changeable in Settings): **Actual MH**, **Burndown MH**, or **Actual and burndown MH**.
  - Actual MH = hours spent. Burndown MH = hours that really moved the work forward (lower when someone is stuck on a problem or blocked).
  - In "Actual and burndown" each day has two boxes: actual on top, burndown below. Fill one and the other copies it automatically (shown faded); fill both and each keeps its own value.
  - A switch beside Export Excel shows Actual, Burndown or Both.
  - In the burndown modes, Done % (and so the sprint burndown) follows burndown MH ÷ planned MH; "MH spent" follows actual. The Excel export has both columns.

## Version 2.7.2
- **Gantt filter** (next to Plan / Actual): All tasks, Needs attention (late, behind, not started yet, blocked), Late only, My open tasks. Each choice shows how many tasks match. Parent rows stay, dimmed, so you can see where a task belongs. Alt+F switches between them.
- **Board: change Done % on the card**: drag or click the progress bar (steps of 5%). Moving above 0% starts a To do task; 100% marks it Done.
- **Sutora**: when tasks are late or behind, she offers to highlight them on the Board (and opens it), or, on the Gantt, to show only those tasks.

## Version 2.7.1
- Sutora's look picker shows just the pictures, without names.
- New third choice: no picture. Sutora then sits as a small speech-bubble button and still gives tips.

## Version 2.7: team availability
- **National holidays, built in (no internet)**: Settings → Holidays → "Add national holidays". Pick Thailand or Japan and the years, tick or untick each day, rename any of them, Apply.
  - Thailand: fixed days, weekend substitution days, and the Buddhist days (official lists for 2026 and 2027; 2028 marked as estimate). Khao Phansa is offered but unticked (government only).
  - Japan: calculated for any year, including equinoxes, substitute holidays and citizens' holidays; the year-end company break is offered unticked.
- **Personal schedule** (Team → Schedule, or the Timesheet): fill once per person.
  - MH per day and which country's public holidays apply to that person.
  - Every week: classes, rest time, fixed meetings, with days and hours.
  - One-off: leave (all day or hours), training, visits.
  - Import busy time from **classic Outlook** on the same PC (desktop app, one click, no internet) or from an **.ics file** (new Outlook, Google, any calendar). Recurring meetings, time zones, out-of-office, free, declined and cancelled items are handled. Only busy hours within 09:00–18:00 are kept, never meeting titles.
- **Available MH** everywhere: the Timesheet shows available hours per day and logged / available (red when over), the Team sheet shows this week's logged / available, and Sutora's overload warning uses real availability.
- **Timesheet: Show → Everyone** (or one person): everyone's hours grouped by person, with a team total.
- **Edit rights per role** (Team → Roles → "Edit others' hours"): by default Project owner, Manager and Team leader may change other people's hours and schedules. The person gets a notice from Sutora. sTrack has no passwords, so this prevents mistakes; it is not security.

## Version 2.6
### Sutora
- **Finds the delay**: when the forecast passes the target, Sutora follows the chain of tasks that sets the end date and tells you which ones are late (and who owns them), which should have started but haven't, and whether the plan itself is simply longer than the target. "Show me" outlines that chain in the Gantt.
- **Shows exactly where**: "Show me" now draws a clear pulsing ring, labelled Sutora, around the field or row she means.
- **First use**: Sutora asks which look and size you'd like, and reminds you she can be hidden or changed anytime.
- Suggests setting up roles before adding members, and points out batch editing once a plan has many tasks.
- **Templates per milestone**: choose where a template goes: leading to a milestone (the milestone waits for it), after a milestone, or inside a phase. You can also put all its tasks into a sprint. The guided setup now suggests milestones or sprints first.
### Gantt
- Drag a row by its number to move it (sub-steps move with it).
- PIC list: click PIC again to close it, or press OK.
- Date boxes open the calendar on the first click.
### Board
- **Done % per column** (Board settings): moving a card sets Done % to at least that number; To do sets 0. Defaults: To do 0, In progress 30, Review 70, Done 100, Blocked and On hold unchanged.
- **Card layout**: choose what a card shows and in which order (WBS, dates, days, MH, Done %, PIC, second name, sprint, status, progress bar, note, attention tag).
- Colour cards by status (on/off) and mark tasks that need attention: should have started but haven't, behind the plan's pace, or late.
- Sort each column by plan order, PIC, start, end or Done %.
### Timesheet
- MH boxes are always outlined so it's clear where to type.
- **Note after MH**: after typing hours, a small box asks what you did (Ctrl+Enter saves, Esc skips). A dot marks cells with a note; hover to read it. Switch it off in the bar.
- **Export Excel** for the week in view, with a Notes sheet (Ctrl+6).
### Team
- **Members | Roles** tabs. Roles can be added, renamed (members follow), ordered and removed.
### Fixes
- Saving an Excel or PNG file no longer offers the wrong file type in the save dialog.

## Version 2.5
- **Template preview (V-model)**: unticking a design step also unticks its test, and ticking it again brings the test back (works for whole phases too). Click any box in the picture to tick or untick that step; unticked steps stay in the picture, faded and dashed, so you can bring them back. The list no longer jumps to the top after each click.
- **Arrow keys** on the Gantt: ↑ ↓ select the row above or below (Shift+↑ ↓ adds rows to the selection), → opens a group, ← closes it or goes to its parent, Enter edits the name. While editing a cell, ↑ ↓ move to the same column in the next row and ← → at the start or end of the text move to the neighbouring column.
- On the Board, the arrow keys move between cards and columns; Enter opens the card.
- **Japanese / Thai input**: while the IME is converting (candidate list open), the arrow keys and Enter stay with the IME, so they no longer jump to another row or add a task.

## Version 2.4
- **Team roles**: the Role column is now a list. Default: Project owner, Manager, Team leader, PIC. Change the list with **Roles** (Ctrl+2 on the Team page); it is shared with everyone on the project. Roles that are already in use are kept.
- **Member photos**: click the picture circle on the Team page (or drop an image onto the row) to add a photo. It is cropped to a square and made small (about 10 KB) so the shared file stays light. Photos appear in the PIC column, the Board, the PIC picker and in captured PNGs. × removes it.
- **Nicknames**: new optional Nickname column. Turn on **Call by nickname** (Ctrl+3 on the Team page) to show nicknames everywhere; people without a nickname are shown by the first word of their full name. Customer view always keeps full names.
- Excel export: the Team sheet has a Nickname column.
- Fix: the version under the app name now always matches the real version.
- Fix: in the V-model template preview, design and test boxes no longer overlap in the middle; hover a box to read its full name.

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
- Sutora has a **second look**, shown as a full picture (no crop, about 96 px tall, background removed) (choose with ⚙ on her card, or in Theme). Classic stays the default. She uses the calm pose when idle and the pointing pose while giving a tip. Size 1×, 1.5× or 2× (⚙ on her card, or Theme). Your choices are remembered on your computer.
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
