package tracker

import "testing"

func TestLevel1AddWorker(t *testing.T) {
	tt := NewTimeTracker()
	eqBool(t, tt.AddWorker(1, "ann", "dev", 10), true)
	eqBool(t, tt.AddWorker(2, "ann", "qa", 5), false)
}

func TestLevel1Register(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	if got := tt.Register(2, "bob"); got != "invalid_request" {
		t.Fatalf("got %q, want invalid_request", got)
	}
	if got := tt.Register(3, "ann"); got != "registered" {
		t.Fatalf("got %q, want registered", got)
	}
	if got := tt.Register(4, "ann"); got != "registered" {
		t.Fatalf("got %q, want registered", got)
	}
}

func TestLevel1GetTime(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	eqInt(t, tt.GetTime(2, "ann"), 0)
	tt.Register(10, "ann")             // in
	eqInt(t, tt.GetTime(15, "ann"), 0) // still inside: not counted yet
	tt.Register(20, "ann")             // out: 10
	eqInt(t, tt.GetTime(21, "ann"), 10)
	tt.Register(30, "ann")
	tt.Register(45, "ann") // +15
	eqInt(t, tt.GetTime(50, "ann"), 25)
	nilInt(t, tt.GetTime(51, "bob"))
}

func TestLevel1WorkersIndependent(t *testing.T) {
	tt := NewTimeTracker()
	tt.AddWorker(1, "ann", "dev", 10)
	tt.AddWorker(2, "bob", "dev", 10)
	tt.Register(10, "ann") // ann in
	tt.Register(12, "bob") // bob in
	tt.Register(20, "ann") // ann out: 10
	eqInt(t, tt.GetTime(21, "ann"), 10)
	eqInt(t, tt.GetTime(22, "bob"), 0)
}
