package tracker

import "testing"

func setupLevel2() *TimeTracker {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.AddWorker(2, "bob", "dev", 10)
	tt.AddWorker(3, "cat", "qa", 10)
	tt.AddWorker(4, "dan", "dev", 10)
	tt.Register(10, "ann")
	tt.Register(11, "bob")
	tt.Register(12, "dan")
	tt.Register(13, "cat")
	tt.Register(20, "ann")  // ann: 10
	tt.Register(22, "dan")  // dan: 10
	tt.Register(31, "bob")  // bob: 20
	tt.Register(100, "cat") // cat: 87
	return tt
}

func TestLevel2Ranking(t *testing.T) {
	tt := setupLevel2()
	eqList(t, tt.TopNWorkers(101, 2, "dev"), []string{"bob(20)", "ann(10)"})
	eqList(t, tt.TopNWorkers(102, 5, "dev"), []string{"bob(20)", "ann(10)", "dan(10)"})
	eqList(t, tt.TopNWorkers(103, 5, "qa"), []string{"cat(87)"})
}

func TestLevel2NoMatchIsEmptyNotNil(t *testing.T) {
	tt := setupLevel2()
	eqList(t, tt.TopNWorkers(101, 3, "pm"), []string{})
	eqList(t, tt.TopNWorkers(102, 0, "dev"), []string{})
}

func TestLevel2ZeroTimeAndOpenSession(t *testing.T) {
	tt := setupLevel2()
	tt.AddWorker(105, "eve", "dev", 10)
	tt.Register(106, "ann") // ann back inside; doesn't count yet
	eqList(t, tt.TopNWorkers(200, 5, "dev"), []string{"bob(20)", "ann(10)", "dan(10)", "eve(0)"})
}
