package tracker

import "testing"

func eqResult(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestLevel3PromoteInvalid(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	eqResult(t, tt.Promote(2, "zed", "lead", 30, 100), "invalid_request")
	eqResult(t, tt.Promote(3, "ann", "lead", 30, 100), "success")
	eqResult(t, tt.Promote(4, "ann", "cto", 99, 200), "invalid_request") // one already pending
}

func TestLevel3PromotionTiming(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.AddWorker(2, "bob", "dev", 20)
	tt.Promote(3, "ann", "lead", 30, 50)
	tt.Register(10, "ann")
	tt.Register(20, "ann") // dev, 10 long
	tt.Register(40, "ann") // enters before 50: still dev
	tt.Register(60, "ann") // dev, 20 long
	tt.Register(70, "ann") // first entry at/after 50: now lead
	tt.Register(80, "ann") // lead, 10 long
	eqInt(t, tt.GetTime(81, "ann"), 40)
	eqList(t, tt.TopNWorkers(82, 5, "dev"), []string{"bob(0)"})
	eqList(t, tt.TopNWorkers(83, 5, "lead"), []string{"ann(10)"})
	eqResult(t, tt.Promote(84, "ann", "cto", 50, 90), "success") // earlier one took effect
}

func TestLevel3CalcSalary(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.Promote(2, "ann", "lead", 30, 50)
	tt.Register(10, "ann")
	tt.Register(20, "ann") // 10 @ 10
	tt.Register(40, "ann")
	tt.Register(60, "ann") // 20 @ 10
	tt.Register(70, "ann")
	tt.Register(80, "ann") // 10 @ 30
	eqInt(t, tt.CalcSalary(81, "ann", 0, 100), 100+200+300)
	// partial overlap: [15,20) @10 + [40,60) @10 + [70,75) @30
	eqInt(t, tt.CalcSalary(82, "ann", 15, 75), 50+200+150)
	eqInt(t, tt.CalcSalary(83, "ann", 20, 40), 0) // gap between sessions
	nilInt(t, tt.CalcSalary(84, "zed", 0, 100))
}

func TestLevel3PromotedWhileInside(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.Register(10, "ann")
	tt.Promote(11, "ann", "lead", 30, 12)
	tt.Register(20, "ann") // whole session stays dev @ 10
	tt.Register(30, "ann") // now lead
	tt.Register(40, "ann") // 10 @ 30
	eqInt(t, tt.CalcSalary(41, "ann", 0, 50), 100+300)
	eqList(t, tt.TopNWorkers(42, 5, "lead"), []string{"ann(10)"})
}

func TestLevel3OpenSessionNotPaid(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.Register(10, "ann")
	eqInt(t, tt.CalcSalary(50, "ann", 0, 100), 0)
}
