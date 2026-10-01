package tasks

import "testing"

func TestLevel4Complete(t *testing.T) {
	m := NewTaskManager()
	m.AddTask(1, "a", 1) // task_1
	m.AddTask(2, "b", 1) // task_2
	m.AddUser(3, "u1", 1)
	eqBool(t, m.AssignTask(10, "task_1", "u1", 20), true)
	eqBool(t, m.CompleteTask(15, "task_1", "u1"), true)
	eqBool(t, m.CompleteTask(16, "task_1", "u1"), false) // already done
	eqList(t, m.GetUserTasks(17, "u1"), []string{})
	eqBool(t, m.AssignTask(18, "task_2", "u1", 30), true) // quota freed
	eqBool(t, m.CompleteTask(30, "task_2", "u1"), false)  // expired at 30
	eqList(t, m.GetOverdueAssignments(31, "u1"), []string{"task_2"})
}

func TestLevel4CompleteInvalid(t *testing.T) {
	m := NewTaskManager()
	m.AddTask(1, "a", 1)
	m.AddUser(2, "u1", 1)
	m.AddUser(3, "u2", 1)
	m.AssignTask(4, "task_1", "u1", 100)
	eqBool(t, m.CompleteTask(5, "task_1", "u2"), false) // not u2's
	eqBool(t, m.CompleteTask(6, "task_9", "u1"), false)
	eqBool(t, m.CompleteTask(7, "task_1", "nobody"), false)
}

func TestLevel4OverdueOrdering(t *testing.T) {
	m := NewTaskManager()
	m.AddTask(1, "a", 1) // task_1
	m.AddTask(2, "b", 1) // task_2
	m.AddTask(3, "c", 1) // task_3
	m.AddUser(4, "u1", 5)
	m.AssignTask(10, "task_1", "u1", 40)
	m.AssignTask(11, "task_2", "u1", 30)
	m.AssignTask(12, "task_3", "u1", 60)
	eqList(t, m.GetOverdueAssignments(29, "u1"), []string{})
	eqList(t, m.GetOverdueAssignments(35, "u1"), []string{"task_2"})
	eqList(t, m.GetOverdueAssignments(45, "u1"), []string{"task_2", "task_1"})
	m.CompleteTask(50, "task_3", "u1")
	eqList(t, m.GetOverdueAssignments(100, "u1"), []string{"task_2", "task_1"})
	eqList(t, m.GetOverdueAssignments(101, "nobody"), []string{})
}

func TestLevel4OverdueTieOnFinishTime(t *testing.T) {
	m := NewTaskManager()
	m.AddTask(1, "a", 1) // task_1
	m.AddTask(2, "b", 1) // task_2
	m.AddUser(3, "u1", 5)
	m.AssignTask(10, "task_2", "u1", 50)
	m.AssignTask(11, "task_1", "u1", 50)
	eqList(t, m.GetOverdueAssignments(60, "u1"), []string{"task_1", "task_2"})
}
