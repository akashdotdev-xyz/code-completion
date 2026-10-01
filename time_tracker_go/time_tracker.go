package tracker

// TimeTracker: see README.md for the full spec.
// A nil pointer return means "invalid".
type TimeTracker struct {
	// TODO: your state here
}

func NewTimeTracker() *TimeTracker {
	return &TimeTracker{}
}

// ---------- Level 1 ----------

func (t *TimeTracker) AddWorker(timestamp int, workerID, position string, compensation int) bool {
	return false
}

func (t *TimeTracker) Register(timestamp int, workerID string) string {
	return ""
}

func (t *TimeTracker) GetTime(timestamp int, workerID string) *int {
	return nil
}

// ---------- Level 2 ----------

func (t *TimeTracker) TopNWorkers(timestamp int, n int, position string) []string {
	return nil
}

// ---------- Level 3 ----------

func (t *TimeTracker) Promote(timestamp int, workerID, newPosition string, newCompensation, startTimestamp int) string {
	return ""
}

func (t *TimeTracker) CalcSalary(timestamp int, workerID string, startTimestamp, endTimestamp int) *int {
	return nil
}

// ---------- Level 4 ----------

func (t *TimeTracker) SetDoublePaid(timestamp int, startTimestamp, endTimestamp int) {
}
