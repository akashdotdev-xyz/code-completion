package tasks

// TaskManager: see README.md for the full spec.
// A nil pointer return means "not found".
type TaskManager struct {
	// TODO: your state here
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// ---------- Level 1 ----------

func (m *TaskManager) AddTask(timestamp int, name string, priority int) string {
	return ""
}

func (m *TaskManager) UpdateTask(timestamp int, taskID, name string, priority int) bool {
	return false
}

func (m *TaskManager) GetTask(timestamp int, taskID string) *string {
	return nil
}

// ---------- Level 2 ----------

func (m *TaskManager) ListTasks(timestamp int, limit int) []string {
	return nil
}

func (m *TaskManager) SearchTasks(timestamp int, nameFilter string, maxResults int) []string {
	return nil
}

// ---------- Level 3 ----------

func (m *TaskManager) AddUser(timestamp int, userID string, quota int) bool {
	return false
}

func (m *TaskManager) AssignTask(timestamp int, taskID, userID string, finishTime int) bool {
	return false
}

func (m *TaskManager) GetUserTasks(timestamp int, userID string) []string {
	return nil
}

// ---------- Level 4 ----------

func (m *TaskManager) CompleteTask(timestamp int, taskID, userID string) bool {
	return false
}

func (m *TaskManager) GetOverdueAssignments(timestamp int, userID string) []string {
	return nil
}
