package db

import "testing"

func TestLevel1SetGet(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "A", "B", "E")
	d.Set(2, "A", "C", "F")
	eqStr(t, d.Get(3, "A", "B"), "E")
	eqStr(t, d.Get(4, "A", "C"), "F")
	nilStr(t, d.Get(5, "A", "D")) // missing field
	nilStr(t, d.Get(6, "X", "B")) // missing record
}

func TestLevel1Overwrite(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "A", "B", "E")
	d.Set(2, "A", "B", "G")
	eqStr(t, d.Get(3, "A", "B"), "G")
}

func TestLevel1Delete(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "A", "B", "E")
	d.Set(2, "A", "C", "F")
	eqBool(t, d.Delete(3, "A", "B"), true)
	eqBool(t, d.Delete(4, "A", "B"), false)
	nilStr(t, d.Get(5, "A", "B"))
	eqStr(t, d.Get(6, "A", "C"), "F") // other fields untouched
	eqBool(t, d.Delete(7, "X", "Y"), false)
}

func TestLevel1SameFieldDifferentKeys(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "A", "F", "1")
	d.Set(2, "B", "F", "2")
	eqStr(t, d.Get(3, "A", "F"), "1")
	eqStr(t, d.Get(4, "B", "F"), "2")
}
