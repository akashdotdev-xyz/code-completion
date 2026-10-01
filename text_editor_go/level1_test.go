package editor

import "testing"

func eqText(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestLevel1Append(t *testing.T) {
	e := NewTextEditor()
	eqText(t, e.Append(1, "Hey"), "Hey")
	eqText(t, e.Append(2, " you"), "Hey you")
}

func TestLevel1MoveAndInsert(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "Hey you")
	eqText(t, e.Move(2, 3), "Hey you")
	eqText(t, e.Append(3, ","), "Hey, you")
	eqText(t, e.Append(4, "!"), "Hey,! you") // cursor moved past ","
}

func TestLevel1Backspace(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "Hey you")
	e.Move(2, 3)
	eqText(t, e.Backspace(3), "He you")
	eqText(t, e.Backspace(4), "H you")
	eqText(t, e.Append(5, "i"), "Hi you")
}

func TestLevel1Clamping(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "abc")
	e.Move(2, 100)
	eqText(t, e.Append(3, "!"), "abc!")
	e.Move(4, -5)
	eqText(t, e.Backspace(5), "abc!") // nothing before position 0
	eqText(t, e.Append(6, ">"), ">abc!")
}

func TestLevel1EmptyDocument(t *testing.T) {
	e := NewTextEditor()
	eqText(t, e.Backspace(1), "")
	eqText(t, e.Move(2, 3), "")
	eqText(t, e.Append(3, "x"), "x")
}
