# sGantt (smart Gantt) — release notes

## Version 1.0

Offline Gantt chart and project monitoring tool. Single-file web app + Windows exe. Nothing is uploaded.

### Plan
- Task list with up to 8 levels: indent / outdent, move up / down, collapse. Parent rows add up dates, hours and progress (weighted by days).
- Project start date and an optional target in working days (drawn as a target line).
- Headline figures: plan end, duration in working days, forecast end, target end, overall progress ring and status (on target / over target / behind plan).

### Calendar
- Saturday and Sunday can each be a working day or not.
- Holidays: add them in the Calendar panel, or click a date in the timeline header.
- Hours per day; task days are calculated from hours (rounded up).

### Tasks
- Start date (empty = starts automatically after its predecessors, or on the project start).
- Days or hours: type either one, the other is calculated. 0 days = milestone.
- After: predecessor row numbers (finish → start). Arrows are drawn between bars; a red arrow means the task starts too early.
- Progress %, actual start (filled in automatically when progress starts) and actual days.

### Views
- Plan, Actual, or Plan vs Actual (plan bar above, actual bar below). Late actual bars turn red with "N d late".
- Zoom, jump to today, three column sets (all / plan / name only).
- 5 themes: Basic, Dark, Starry Night, Forest, Burning.

### Files
- Autosave in the browser (or `%LOCALAPPDATA%\sGantt` for the exe).
- Save / Open plan as `.sgantt.json`.
- Export PNG with title block, key figures and legend, on a white background or in the theme colours.
- Undo / redo. Keyboard shortcuts work on any keyboard layout (Thai, English, Chinese, Japanese, …). Dates can be typed with the Thai year (16/11/2569).
