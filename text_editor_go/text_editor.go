package editor

// TextEditor: see README.md for the full spec.
type TextEditor struct {
	// TODO: your state here
}

func NewTextEditor() *TextEditor {
	return &TextEditor{}
}

// ---------- Level 1 ----------

func (e *TextEditor) Append(timestamp int, text string) string {
	return ""
}

func (e *TextEditor) Move(timestamp int, position int) string {
	return ""
}

func (e *TextEditor) Backspace(timestamp int) string {
	return ""
}

// ---------- Level 2 ----------

func (e *TextEditor) Select(timestamp int, left, right int) string {
	return ""
}

func (e *TextEditor) Copy(timestamp int) string {
	return ""
}

func (e *TextEditor) Paste(timestamp int) string {
	return ""
}

// ---------- Level 3 ----------

func (e *TextEditor) Undo(timestamp int) string {
	return ""
}

func (e *TextEditor) Redo(timestamp int) string {
	return ""
}

// ---------- Level 4 ----------

func (e *TextEditor) Create(timestamp int, name string) bool {
	return false
}

func (e *TextEditor) Switch(timestamp int, name string) *string {
	return nil
}

func (e *TextEditor) ListDocuments(timestamp int) []string {
	return nil
}
