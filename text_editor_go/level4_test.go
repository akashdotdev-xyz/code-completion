package editor

import "testing"

func TestLevel4CreateAndSwitch(t *testing.T) {
	e := NewTextEditor()
	eqBool(t, e.Create(1, "notes"), true)
	eqBool(t, e.Create(2, "main"), false)
	eqBool(t, e.Create(3, "notes"), false)
	e.Append(4, "hello") // still on main
	eqStr(t, e.Switch(5, "notes"), "")
	eqText(t, e.Append(6, "x"), "x")
	eqStr(t, e.Switch(7, "main"), "hello")
	eqText(t, e.Append(8, "!"), "hello!") // cursor kept at the end
	nilStr(t, e.Switch(9, "nope"))
	eqText(t, e.Append(10, "?"), "hello!?") // still on main
}

func TestLevel4SeparateHistory(t *testing.T) {
	e := NewTextEditor()
	e.Create(1, "notes")
	e.Append(2, "a")
	e.Switch(3, "notes")
	e.Append(4, "b")
	e.Switch(5, "main")
	eqText(t, e.Undo(6), "")
	eqStr(t, e.Switch(7, "notes"), "b")
	eqText(t, e.Undo(8), "")
	eqText(t, e.Redo(9), "b")
}

func TestLevel4SharedClipboard(t *testing.T) {
	e := NewTextEditor()
	e.Create(1, "notes")
	e.Append(2, "copyme")
	e.Select(3, 0, 6)
	e.Copy(4)
	e.Switch(5, "notes")
	eqText(t, e.Paste(6), "copyme")
}

func TestLevel4SelectionPerDocument(t *testing.T) {
	e := NewTextEditor()
	e.Create(1, "notes")
	e.Append(2, "abc")
	e.Select(3, 0, 1)
	e.Switch(4, "notes")
	e.Switch(5, "main")
	eqText(t, e.Append(6, "Z"), "Zbc")
}

func TestLevel4ListDocuments(t *testing.T) {
	e := NewTextEditor()
	eqList(t, e.ListDocuments(1), []string{"main(0)"})
	e.Create(2, "b")
	e.Create(3, "a")
	e.Append(4, "xyz") // main
	e.Switch(5, "b")
	e.Append(6, "xyz")
	eqList(t, e.ListDocuments(7), []string{"b(3)", "main(3)", "a(0)"})
}
