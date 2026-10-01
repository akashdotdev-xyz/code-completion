package tasks

import "testing"

func setupLevel3() *TaskManager {
	m := NewTaskManager()
	m.AddTask(1, "a", 1) // task_1
	m.AddTask(2, "b", 1) // task_2
	m.AddTask(3, "c", 1) // task_3
	m.AddUser(4, "u1", 2)
	return m
}

func TestLevel3AddUser(t *testing.T) {
	m := setupLevel3()
	eqBool(t, m.AddUser(5, "u1", 3), false)
	eqBool(t, m.AddUser(6, "u2", 3), true)
}

func TestLevel3QuotaAndExpiry(t *testing.T) {
	m := setupLevel3()
	eqBool(t, m.AssignTask(10, "task_1", "u1", 100), true)
	eqBool(t, m.AssignTask(11, "task_2", "u1", 50), true)
	eqBool(t, m.AssignTask(12, "task_3", "u1", 200), false) // quota full
	eqList(t, m.GetUserTasks(13, "u1"), []string{"task_2", "task_1"})
	// task_2 was active during [11, 50), so at 50 there's room again
	eqBool(t, m.AssignTask(50, "task_3", "u1", 200), true)
	eqList(t, m.GetUserTasks(51, "u1"), []string{"task_1", "task_3"})
}

func TestLevel3Invalid(t *testing.T) {
	m := setupLevel3()
	eqBool(t, m.AssignTask(10, "task_9", "u1", 100), false)
	eqBool(t, m.AssignTask(11, "task_1", "nobody", 100), false)
	eqList(t, m.GetUserTasks(12, "nobody"), []string{})
	eqList(t, m.GetUserTasks(13, "u1"), []string{})
}

func TestLevel3NoDuplicateActive(t *testing.T) {
	m := setupLevel3()
	eqBool(t, m.AssignTask(10, "task_1", "u1", 100), true)
	eqBool(t, m.AssignTask(11, "task_1", "u1", 200), false)
	eqBool(t, m.AssignTask(100, "task_1", "u1", 300), true) // old one expired
}

func TestLevel3SameTaskManyUsers(t *testing.T) {
	m := setupLevel3()
	m.AddUser(5, "u2", 1)
	eqBool(t, m.AssignTask(10, "task_1", "u1", 100), true)
	eqBool(t, m.AssignTask(11, "task_1", "u2", 100), true)
	eqList(t, m.GetUserTasks(12, "u2"), []string{"task_1"})
}

func TestLevel3TieOnFinishTime(t *testing.T) {
	m := setupLevel3()
	m.AssignTask(10, "task_3", "u1", 100)
	m.AssignTask(11, "task_1", "u1", 100)
	eqList(t, m.GetUserTasks(12, "u1"), []string{"task_1", "task_3"})
}
