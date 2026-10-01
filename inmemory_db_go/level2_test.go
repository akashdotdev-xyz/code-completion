package db

import "testing"

func setupLevel2() *InMemoryDB {
	d := NewInMemoryDB()
	d.Set(1, "A", "BC", "E")
	d.Set(2, "A", "C", "G")
	d.Set(3, "A", "BD", "F")
	d.Set(4, "B", "X", "1")
	return d
}

func TestLevel2Scan(t *testing.T) {
	d := setupLevel2()
	eqList(t, d.Scan(5, "A"), []string{"BC(E)", "BD(F)", "C(G)"})
	eqList(t, d.Scan(6, "B"), []string{"X(1)"})
}

func TestLevel2ScanMissingIsEmptyNotNil(t *testing.T) {
	d := setupLevel2()
	eqList(t, d.Scan(5, "nope"), []string{})
	d.Delete(6, "B", "X")
	eqList(t, d.Scan(7, "B"), []string{})
}

func TestLevel2ScanAfterDelete(t *testing.T) {
	d := setupLevel2()
	d.Delete(5, "A", "BD")
	eqList(t, d.Scan(6, "A"), []string{"BC(E)", "C(G)"})
}

func TestLevel2ScanByPrefix(t *testing.T) {
	d := setupLevel2()
	eqList(t, d.ScanByPrefix(5, "A", "B"), []string{"BC(E)", "BD(F)"})
	eqList(t, d.ScanByPrefix(6, "A", "BD"), []string{"BD(F)"})
	eqList(t, d.ScanByPrefix(7, "A", "Z"), []string{})
	eqList(t, d.ScanByPrefix(8, "nope", "B"), []string{})
	eqList(t, d.ScanByPrefix(9, "A", ""), []string{"BC(E)", "BD(F)", "C(G)"})
}

func TestLevel2StringOrder(t *testing.T) {
	d := NewInMemoryDB()
	d.Set(1, "K", "B2", "x")
	d.Set(2, "K", "B10", "y")
	d.Set(3, "K", "A", "z")
	eqList(t, d.Scan(4, "K"), []string{"A(z)", "B10(y)", "B2(x)"})
}
