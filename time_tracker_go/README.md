# Practice: Employee Time Tracker (CodeSignal ICA style)

**Format:** 90 minutes, 4 levels, one codebase. Each level unlocks only after the
previous level passes. Levels 3 and 4 compute pay over time ranges, so keep a
list of each worker's office sessions from the start, not just a running total.

Implement `TimeTracker` in `time_tracker.go`. Run tests one level at a time:

```bash
cd time_tracker_go
go test -run Level1 -v
go test -run Level2 -v
go test -run Level3 -v
go test -run Level4 -v
```

**General rules**
- Every method takes a `timestamp`. Timestamps across calls are strictly
  increasing.
- A **session** is the time from when a worker enters the office to when they
  leave. Only finished sessions count toward time or pay; a worker still in the
  office contributes nothing until they leave.
- A `nil` return means "invalid". Never panic.

---

## Level 1: Clocking in and out (target: ~10 min)

- `AddWorker(timestamp int, workerID, position string, compensation int) bool`
  Adds a worker with a position and pay per time unit. Returns `false` if the
  worker already exists.
- `Register(timestamp int, workerID string) string`
  The worker enters the office at `timestamp` if they're outside, or leaves if
  they're inside. Returns `"registered"`, or `"invalid_request"` if the worker
  doesn't exist.
- `GetTime(timestamp int, workerID string) *int`
  Total length of all the worker's finished sessions, or `nil` if the worker
  doesn't exist.

## Level 2: Ranking (target: ~15 min)

- `TopNWorkers(timestamp int, n int, position string) []string`
  Returns up to `n` workers whose **current** position is `position`, formatted
  `"workerID(time)"`, where time is their finished-session time **in that
  position** (see level 3). Sort by time descending, ties by ID ascending.
  Returns an empty slice (not `nil`) if nobody matches.

## Level 3: Promotions and salary (target: ~25 min)

- `Promote(timestamp int, workerID, newPosition string, newCompensation, startTimestamp int) string`
  Schedules a promotion. It takes effect the first time the worker **enters** the
  office at or after `startTimestamp`. A session that's already running when
  `startTimestamp` passes keeps the old position and pay.
  Returns `"success"`, or `"invalid_request"` if the worker doesn't exist or
  already has a promotion that hasn't taken effect yet.
- Once a promotion takes effect, `TopNWorkers` counts only time spent **since**
  the promotion. `GetTime` still counts everything.
- `CalcSalary(timestamp int, workerID string, startTimestamp, endTimestamp int) *int`
  Pay earned during `[startTimestamp, endTimestamp)`: for each finished session,
  the part of it that falls inside that range, times the compensation the
  session was worked at. Returns `nil` if the worker doesn't exist.

## Level 4: Double pay periods (target: ~30 min)

- `SetDoublePaid(timestamp int, startTimestamp, endTimestamp int)`
  For every worker, time worked during `[startTimestamp, endTimestamp)` is paid at
  twice their compensation. This applies to past and future sessions alike.
  Double pay periods never overlap each other.
- `CalcSalary` must include double pay. `GetTime` and `TopNWorkers` are not
  affected.
