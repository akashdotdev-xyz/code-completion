package db

import "testing"

func eqCount(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func setupLevel4() *InMemoryDB {
	d := NewInMemoryDB()
	d.Set(1, "A", "B", "1")
	d.SetWithTTL(2, "A", "C", "2", 10) // expires at 12
	d.SetWithTTL(3, "X", "Y", "3", 5)  // expires at 8
	return d
}

func TestLevel4BackupCount(t *testing.T) {
	d := NewInMemoryDB()
	eqCount(t, d.Backup(1), 0)
	d = setupLevel4()
	eqCount(t, d.Backup(4), 2)
	eqCount(t, d.Backup(9), 1) // X's only field expired at 8
	d.Delete(10, "A", "B")
	eqCount(t, d.Backup(11), 1) // A still has C until 12
	eqCount(t, d.Backup(12), 0)
}

func TestLevel4RestoreRecomputesTTL(t *testing.T) {
	d := setupLevel4()
	d.Backup(4) // remaining: A.C = 8, X.Y = 4
	d.Set(5, "A", "B", "changed")
	d.Delete(6, "A", "C")
	d.Restore(20, 5) // latest backup at or before 5 is the one at 4
	// new expiries: A.C = 28, X.Y = 24
	eqStr(t, d.Get(21, "A", "B"), "1")
	eqStr(t, d.Get(23, "X", "Y"), "3")
	nilStr(t, d.Get(24, "X", "Y"))
	eqStr(t, d.Get(27, "A", "C"), "2")
	nilStr(t, d.Get(28, "A", "C"))
}

func TestLevel4RestorePicksLatestBackup(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "K", "F", "v1")
	d.Backup(2)
	d.Set(3, "K", "F", "v2")
	d.Backup(4)
	d.Set(5, "K", "F", "v3")
	d.Restore(6, 3) // backup at 2
	eqStr(t, d.Get(7, "K", "F"), "v1")
	d.Restore(8, 4) // backup at exactly 4
	eqStr(t, d.Get(9, "K", "F"), "v2")
	d.Restore(10, 100) // latest backup overall is still at 4
	eqStr(t, d.Get(11, "K", "F"), "v2")
}

func TestLevel4RestoreReplacesEverything(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "A", "B", "1")
	d.Backup(2)
	d.Set(3, "New", "F", "x")
	d.Set(4, "A", "Extra", "y")
	d.Restore(5, 2)
	nilStr(t, d.Get(6, "New", "F"))
	eqList(t, d.Scan(7, "A"), []string{"B(1)"})
}

func TestLevel4BackupIsIsolated(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "A", "B", "1")
	d.Backup(2)
	d.Restore(3, 2)
	d.Set(4, "A", "B", "changed") // must not leak into the saved backup
	d.Set(5, "A", "C", "added")
	d.Restore(6, 2)
	eqList(t, d.Scan(7, "A"), []string{"B(1)"})
}

func TestLevel4ExpiredFieldsNotRestored(t *testing.T) {
	d := NewInMemoryDB()
	d.SetWithTTL(1, "A", "B", "gone", 3) // expires at 4
	d.Set(2, "A", "C", "stays")
	d.Backup(5)
	d.Restore(6, 5)
	eqList(t, d.Scan(7, "A"), []string{"C(stays)"})
}
