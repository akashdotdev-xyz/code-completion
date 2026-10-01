package library

import "testing"

func TestLevel1Add(t *testing.T) {
	l := NewLibrary()
	eqBool(t, l.AddBook(1, "b1", "Go Programming"), true)
	eqBool(t, l.AddBook(2, "b1", "Other"), false)
	eqBool(t, l.AddMember(3, "m1"), true)
	eqBool(t, l.AddMember(4, "m1"), false)
}

func TestLevel1BorrowAndReturn(t *testing.T) {
	l := NewLibrary()
	l.AddBook(1, "b1", "Go Programming")
	l.AddMember(2, "m1")
	l.AddMember(3, "m2")
	eqBool(t, l.Borrow(4, "m1", "b1"), true)
	eqBool(t, l.Borrow(5, "m2", "b1"), false) // already lent out
	eqBool(t, l.Return(6, "m2", "b1"), false) // m2 doesn't have it
	eqBool(t, l.Return(7, "m1", "b1"), true)
	eqBool(t, l.Return(8, "m1", "b1"), false)
	eqBool(t, l.Borrow(9, "m2", "b1"), true)
}

func TestLevel1BorrowInvalid(t *testing.T) {
	l := NewLibrary()
	l.AddBook(1, "b1", "Go Programming")
	l.AddMember(2, "m1")
	eqBool(t, l.Borrow(3, "nobody", "b1"), false)
	eqBool(t, l.Borrow(4, "m1", "nothing"), false)
	eqBool(t, l.Return(5, "nobody", "b1"), false)
}

func TestLevel1SeveralBooks(t *testing.T) {
	l := NewLibrary()
	l.AddBook(1, "b1", "One")
	l.AddBook(2, "b2", "Two")
	l.AddMember(3, "m1")
	eqBool(t, l.Borrow(4, "m1", "b1"), true)
	eqBool(t, l.Borrow(5, "m1", "b2"), true)
	eqBool(t, l.Return(6, "m1", "b1"), true)
	eqBool(t, l.Return(7, "m1", "b2"), true)
}
