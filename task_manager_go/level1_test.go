package tasks

import "testing"

func eqPlain(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestLevel1AddAndGet(t *testing.T) {
	m := NewTaskManager()
	eqPlain(t, m.AddTask(1, "write", 5), "task_1")
	eqPlain(t, m.AddTask(2, "read", 3), "task_2")
	eqStr(t, m.GetTask(3, "task_1"), "write(5)")
	eqStr(t, m.GetTask(4, "task_2"), "read(3)")
	nilStr(t, m.GetTask(5, "task_9"))
}

func TestLevel1DuplicateNames(t *testing.T) {
	m := NewTaskManager()
	eqPlain(t, m.AddTask(1, "same", 1), "task_1")
	eqPlain(t, m.AddTask(2, "same", 1), "task_2")
}

func TestLevel1Update(t *testing.T) {
	m := NewTaskManager()
	m.AddTask(1, "read", 3)
	eqBool(t, m.UpdateTask(2, "task_1", "review", 7), true)
	eqStr(t, m.GetTask(3, "task_1"), "review(7)")
	eqBool(t, m.UpdateTask(4, "task_5", "x", 1), false)
	eqPlain(t, m.AddTask(5, "next", 1), "task_2") // counter unaffected by updates
}
