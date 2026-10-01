package db

import "testing"

func TestLevel3ExpiryBoundary(t *testing.T) {
	d := NewInMemoryDB()
	d.SetWithTTL(1, "A", "B", "C", 10) // alive during [1, 11)
	eqStr(t, d.Get(10, "A", "B"), "C")
	nilStr(t, d.Get(11, "A", "B"))
}

func TestLevel3DeleteExpired(t *testing.T) {
	d := NewInMemoryDB()
	d.SetWithTTL(1, "A", "B", "C", 5)
	eqBool(t, d.Delete(6, "A", "B"), false)
}

func TestLevel3DeleteBeforeExpiry(t *testing.T) {
	d := NewInMemoryDB()
	d.SetWithTTL(1, "A", "B", "C", 5)
	eqBool(t, d.Delete(3, "A", "B"), true)
	nilStr(t, d.Get(4, "A", "B"))
}

func TestLevel3ScanSkipsExpired(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "A", "X", "1")
	d.SetWithTTL(2, "A", "Y", "2", 5)  // gone at 7
	d.SetWithTTL(3, "A", "Z", "3", 20) // gone at 23
	eqList(t, d.Scan(6, "A"), []string{"X(1)", "Y(2)", "Z(3)"})
	eqList(t, d.Scan(7, "A"), []string{"X(1)", "Z(3)"})
	eqList(t, d.ScanByPrefix(8, "A", "Y"), []string{})
	eqList(t, d.Scan(23, "A"), []string{"X(1)"})
}

func TestLevel3PlainSetClearsTTL(t *testing.T) {
	d := NewInMemoryDB()
	d.SetWithTTL(1, "A", "B", "C", 10)
	d.Set(5, "A", "B", "D")
	eqStr(t, d.Get(100, "A", "B"), "D")
}

func TestLevel3SetWithTTLReplacesTTL(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "A", "X", "1")
	d.SetWithTTL(2, "A", "X", "2", 5) // now expires at 7
	eqStr(t, d.Get(6, "A", "X"), "2")
	nilStr(t, d.Get(7, "A", "X"))

	d.SetWithTTL(10, "A", "Y", "a", 5)  // expires at 15
	d.SetWithTTL(12, "A", "Y", "b", 20) // now expires at 32
	eqStr(t, d.Get(20, "A", "Y"), "b")
}

func TestLevel3SetAfterExpiry(t *testing.T) {
	d := NewInMemoryDB()
	d.SetWithTTL(1, "A", "B", "old", 2)
	d.Set(5, "A", "B", "new")
	eqStr(t, d.Get(6, "A", "B"), "new")
}
