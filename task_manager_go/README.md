# Practice: Task Manager (CodeSignal ICA style)

**Format:** 90 minutes, 4 levels, one codebase. Each level unlocks only after the
previous level passes. Levels 3 and 4 treat an assignment (task + user + finish
time) as its own thing, so plan to store assignments separately from tasks.

Implement `TaskManager` in `task_manager.go`. Run tests one level at a time:

```bash
cd task_manager_go
go test -run Level1 -v
go test -run Level2 -v
go test -run Level3 -v
go test -run Level4 -v
```

**General rules**
- Every method takes a `timestamp`. Timestamps across calls are strictly
  increasing.
- A `nil` return means "not found". Never panic.

---

## Level 1: Tasks (target: ~10 min)

- `AddTask(timestamp int, name string, priority int) string`
  Creates a task and returns its ID: `"task_1"`, `"task_2"`, … in creation
  order. Names don't have to be unique.
- `UpdateTask(timestamp int, taskID, name string, priority int) bool`
  Replaces the task's name and priority. Returns `false` if the task doesn't exist.
- `GetTask(timestamp int, taskID string) *string`
  Returns `"name(priority)"`, e.g. `"write docs(5)"`, or `nil` if the task
  doesn't exist.

## Level 2: Sorting and search (target: ~15 min)

Task order everywhere: priority descending, then creation order (so
`"task_2"` comes before `"task_10"`; compare the numbers, not the strings).

- `ListTasks(timestamp int, limit int) []string`
  Returns up to `limit` task IDs in task order.
- `SearchTasks(timestamp int, nameFilter string, maxResults int) []string`
  Returns up to `maxResults` IDs of tasks whose name contains `nameFilter`
  (case-sensitive), in task order.
- Both return an empty slice (not `nil`) when there's nothing to return.

## Level 3: Users and assignments (target: ~25 min)

- `AddUser(timestamp int, userID string, quota int) bool`
  Adds a user who can have at most `quota` active assignments at once. Returns
  `false` if the user already exists.
- `AssignTask(timestamp int, taskID, userID string, finishTime int) bool`
  Assigns the task to the user. The assignment is **active** during
  `[timestamp, finishTime)`; `finishTime > timestamp` is guaranteed.
  Returns `false` if the task or user doesn't exist, the user already has an
  active assignment of this task, or the user's quota is full.
  A task can be assigned to several users.
- `GetUserTasks(timestamp int, userID string) []string`
  Task IDs of the user's active assignments, sorted by finish time ascending,
  ties by creation order of the task. Empty slice if the user doesn't exist.

## Level 4: Completion and overdue (target: ~30 min)

- `CompleteTask(timestamp int, taskID, userID string) bool`
  Marks the user's active assignment of this task as done. It stops being active,
  which frees one slot of the user's quota. Returns `false` if there's no such
  active assignment (including one whose finish time has already passed).
- `GetOverdueAssignments(timestamp int, userID string) []string`
  Task IDs of the user's assignments whose finish time is `<= timestamp` and that
  were never completed, sorted by finish time ascending, ties by creation order
  of the task. Empty slice if the user doesn't exist.
