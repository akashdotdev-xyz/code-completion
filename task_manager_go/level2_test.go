package tasks

import (
	"fmt"
	"testing"
)

func setupLevel2() *TaskManager {
	m := NewTaskManager()
	m.AddTask(1, "fix bug", 5)    // task_1
	m.AddTask(2, "write docs", 3) // task_2
	m.AddTask(3, "fix typo", 5)   // task_3
	m.AddTask(4, "deploy", 8)     // task_4
	return m
}

func TestLevel2ListTasks(t *testing.T) {
	m := setupLevel2()
	eqList(t, m.ListTasks(5, 10), []string{"task_4", "task_1", "task_3", "task_2"})
	eqList(t, m.ListTasks(6, 2), []string{"task_4", "task_1"})
	eqList(t, m.ListTasks(7, 0), []string{})
}

func TestLevel2Search(t *testing.T) {
	m := setupLevel2()
	eqList(t, m.SearchTasks(5, "fix", 5), []string{"task_1", "task_3"})
	eqList(t, m.SearchTasks(6, "fix", 1), []string{"task_1"})
	eqList(t, m.SearchTasks(7, "o", 10), []string{"task_4", "task_3", "task_2"})
	eqList(t, m.SearchTasks(8, "Fix", 5), []string{}) // case-sensitive
	eqList(t, m.SearchTasks(9, "zzz", 5), []string{})
}

func TestLevel2UpdateChangesOrder(t *testing.T) {
	m := setupLevel2()
	m.UpdateTask(5, "task_2", "fix docs", 9)
	eqList(t, m.SearchTasks(6, "fix", 5), []string{"task_2", "task_1", "task_3"})
}

func TestLevel2NumericCreationOrder(t *testing.T) {
	m := NewTaskManager()
	want := []string{}
	for i := 1; i <= 12; i++ {
		m.AddTask(i, "same", 1)
		want = append(want, fmt.Sprintf("task_%d", i))
	}
	eqList(t, m.ListTasks(13, 12), want) // task_2 before task_10
}
