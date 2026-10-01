package editor

import "testing"

func TestLevel2CopyPaste(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "Hello world")
	eqText(t, e.Select(2, 0, 5), "Hello world")
	eqText(t, e.Copy(3), "Hello world")
	e.Move(4, 11)
	eqText(t, e.Paste(5), "Hello worldHello")
	eqText(t, e.Paste(6), "Hello worldHelloHello")
}

func TestLevel2AppendReplacesSelection(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "Hello world")
	e.Select(2, 5, 11)
	eqText(t, e.Append(3, "!"), "Hello!")
	eqText(t, e.Append(4, "?"), "Hello!?") // selection cleared, cursor after "!"
}

func TestLevel2BackspaceDeletesSelection(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "abcdef")
	e.Select(2, 1, 4)
	eqText(t, e.Backspace(3), "aef")
	eqText(t, e.Append(4, "X"), "aXef")
}

func TestLevel2SelectClampAndEmpty(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "abc")
	e.Select(2, 2, 100)
	eqText(t, e.Backspace(3), "ab") // selected just "c"
	e.Select(4, 1, 1)               // empty: cursor at 1
	eqText(t, e.Append(5, "X"), "aXb")
}

func TestLevel2MoveClearsSelection(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "abc")
	e.Select(2, 0, 2)
	e.Move(3, 3)
	eqText(t, e.Append(4, "Z"), "abcZ")
}

func TestLevel2CopyKeepsSelection(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "abc")
	e.Select(2, 0, 1)
	e.Copy(3)
	eqText(t, e.Append(4, "Q"), "Qbc")
}

func TestLevel2ClipboardRules(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "abc")
	eqText(t, e.Paste(2), "abc") // empty clipboard
	e.Select(3, 0, 1)
	e.Copy(4)
	e.Move(5, 3)
	e.Copy(6) // no selection: clipboard still "a"
	eqText(t, e.Paste(7), "abca")
}

func TestLevel2PasteOverSelection(t *testing.T) {
	e := NewTextEditor()
	e.Append(1, "one two")
	e.Select(2, 0, 3)
	e.Copy(3)
	e.Select(4, 4, 7)
	eqText(t, e.Paste(5), "one one")
}
