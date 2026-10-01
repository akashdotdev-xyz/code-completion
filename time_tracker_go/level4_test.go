package tracker

import "testing"

func TestLevel4DoublePaidPastSession(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.Register(10, "ann")
	tt.Register(30, "ann") // 20 @ 10
	tt.SetDoublePaid(31, 15, 25)
	eqInt(t, tt.CalcSalary(32, "ann", 0, 100), 200+100) // 10 of the 20 are doubled
	eqInt(t, tt.CalcSalary(33, "ann", 20, 40), 100+50)  // [20,30) worked, [20,25) doubled
}

func TestLevel4DoublePaidFutureSession(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.SetDoublePaid(2, 50, 60)
	tt.Register(55, "ann")
	tt.Register(70, "ann")                             // 15 worked, [55,60) doubled
	eqInt(t, tt.CalcSalary(71, "ann", 0, 100), 150+50) // 15*10 + 5*10 extra
}

func TestLevel4MultiplePeriodsAndPromotion(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.Promote(2, "ann", "lead", 30, 25)
	tt.SetDoublePaid(3, 5, 15)
	tt.SetDoublePaid(4, 35, 38)
	tt.Register(10, "ann")
	tt.Register(20, "ann") // dev: 10 @ 10, [10,15) doubled
	tt.Register(30, "ann") // lead from here
	tt.Register(40, "ann") // lead: 10 @ 30, [35,38) doubled
	want := (100 + 5*10) + (300 + 3*30)
	eqInt(t, tt.CalcSalary(41, "ann", 0, 100), want)
}

func TestLevel4TimeNotAffected(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.SetDoublePaid(2, 0, 100)
	tt.Register(10, "ann")
	tt.Register(20, "ann")
	eqInt(t, tt.GetTime(21, "ann"), 10)
	eqList(t, tt.TopNWorkers(22, 1, "dev"), []string{"ann(10)"})
	eqInt(t, tt.CalcSalary(23, "ann", 0, 100), 200)
}
