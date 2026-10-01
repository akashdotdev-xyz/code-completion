package library

import "testing"

func setupLevel2() *Library {
	l := NewLibrary()
	l.AddBook(1, "b1", "Go Programming")
	l.AddBook(2, "b2", "Learning Python")
	l.AddBook(3, "b3", "The Go Way")
	l.AddBook(4, "b10", "Algorithms") // "Al-go-rithms" contains "go"
	l.AddMember(5, "m1")
	return l
}

func TestLevel2Search(t *testing.T) {
	l := setupLevel2()
	eqList(t, l.SearchByTitle(6, "go"), []string{"b1", "b10", "b3"})
	eqList(t, l.SearchByTitle(7, "PYTHON"), []string{"b2"})
	eqList(t, l.SearchByTitle(8, "rust"), []string{})
}

func TestLevel2MostBorrowed(t *testing.T) {
	l := setupLevel2()
	eqList(t, l.MostBorrowed(6, 2), []string{"b1(0)", "b10(0)"})
	l.Borrow(10, "m1", "b2")
	l.Return(11, "m1", "b2")
	l.Borrow(12, "m1", "b2")
	l.Return(13, "m1", "b2")
	l.Borrow(14, "m1", "b1")
	eqList(t, l.MostBorrowed(15, 2), []string{"b2(2)", "b1(1)"})
	eqList(t, l.MostBorrowed(16, 10), []string{"b2(2)", "b1(1)", "b10(0)", "b3(0)"})
	eqList(t, l.MostBorrowed(17, 0), []string{})
}

func TestLevel2FailedBorrowNotCounted(t *testing.T) {
	l := setupLevel2()
	l.AddMember(6, "m2")
	l.Borrow(7, "m1", "b1")
	l.Borrow(8, "m2", "b1") // fails
	eqList(t, l.MostBorrowed(9, 1), []string{"b1(1)"})
}
