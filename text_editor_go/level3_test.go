package editor

import "testing"

func TestLevel3UndoRedo(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "a")
	e.Append(2, "b")
	e.Append(3, "c")
	eqText(t, e.Undo(4), "ab")
	eqText(t, e.Undo(5), "a")
	eqText(t, e.Redo(6), "ab")
	eqText(t, e.Append(7, "X"), "abX")
	eqText(t, e.Redo(8), "abX") // new change cleared redo
	eqText(t, e.Undo(9), "ab")
	eqText(t, e.Undo(10), "a")
	eqText(t, e.Undo(11), "")
	eqText(t, e.Undo(12), "") // nothing left
}

func TestLevel3UndoRestoresCursor(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "hello")
	e.Move(2, 0)
	e.Append(3, ">")
	eqText(t, e.Undo(4), "hello")
	eqText(t, e.Append(5, "<"), "<hello") // cursor back at 0
}

func TestLevel3UndoRestoresSelection(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "abcd")
	e.Select(2, 1, 3)
	eqText(t, e.Backspace(3), "ad")
	eqText(t, e.Undo(4), "abcd")
	eqText(t, e.Append(5, "X"), "aXd") // "bc" was selected again
}

func TestLevel3RedoRestoresCursor(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "abc")
	e.Undo(2)
	eqText(t, e.Redo(3), "abc")
	eqText(t, e.Append(4, "d"), "abcd")
}

func TestLevel3NonChangesNotRecorded(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "ab")
	e.Move(2, 0)
	e.Select(3, 0, 1)
	e.Copy(4)
	e.Move(5, 0)
	e.Backspace(6) // at position 0: no change
	eqText(t, e.Undo(7), "") // undoes the Append
}

func TestLevel3PasteIsUndoable(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "ab")
	e.Select(2, 0, 2)
	e.Copy(3)
	e.Move(4, 2)
	eqText(t, e.Paste(5), "abab")
	eqText(t, e.Undo(6), "ab")
	eqText(t, e.Redo(7), "abab")
}
