# Practice: Text Editor (CodeSignal ICA style)

**Format:** 90 minutes, 4 levels, one codebase. Each level unlocks only after the
previous level passes. Level 3 undoes changes to text, cursor *and* selection,
and level 4 needs several independent documents, so keep all editing state in
one struct from the start.

Implement `TextEditor` in `text_editor.go`. Run tests one level at a time:

```bash
cd text_editor_go
go test -run Level1 -v
go test -run Level2 -v
go test -run Level3 -v
go test -run Level4 -v
```

**General rules**
- Every method takes a `timestamp`. Timestamps across calls are strictly
  increasing. (Nothing here depends on time; it's there to match the real test's
  format.)
- The editor has a document with a **cursor**: a position between characters,
  from `0` (before the first character) to `len(text)` (after the last).
- Unless stated otherwise, each method returns the full text of the document
  after the operation.
- Text is plain ASCII.

---

## Level 1: Typing (target: ~10 min)

- `Append(timestamp int, text string) string`
  Inserts `text` at the cursor. The cursor moves to just after the inserted text.
- `Move(timestamp int, position int) string`
  Moves the cursor to `position`, clamped to `[0, len(text)]`.
- `Backspace(timestamp int) string`
  Deletes the character just before the cursor. Does nothing at position `0`.

## Level 2: Selection and clipboard (target: ~20 min)

- `Select(timestamp int, left, right int) string`
  Selects the characters in `[left, right)`. Both are clamped to
  `[0, len(text)]`, and `left <= right` is guaranteed. The cursor moves to
  `right`. If `left == right`, nothing is selected and the cursor is at `left`.
- While there's a selection:
  - `Append` replaces the selected text with the new text.
  - `Backspace` deletes the selected text (and nothing else).
  - Either way, the cursor ends up where the selection started (plus the
    inserted text for `Append`), and the selection is cleared.
- `Move` clears the selection.
- `Copy(timestamp int) string`
  Copies the selected text to the clipboard. The selection stays. With no
  selection, the clipboard is left unchanged.
- `Paste(timestamp int) string`
  Same as `Append` with the clipboard's contents. With an empty clipboard, does
  nothing.

## Level 3: Undo and redo (target: ~25 min)

- Only operations that **change the text** can be undone: `Append`, `Backspace`,
  `Paste`. A `Backspace` at position 0 or a `Paste` with an empty clipboard
  changes nothing and isn't recorded. `Move`, `Select` and `Copy` are never
  recorded.
- `Undo(timestamp int) string`
  Reverts the most recent recorded change, restoring the text, cursor and
  selection to exactly what they were right before it. Does nothing if there's
  nothing to undo.
- `Redo(timestamp int) string`
  Re-applies the most recently undone change, restoring the text, cursor and
  selection to what they were right after it. Any new recorded change clears
  everything that could be redone.

## Level 4: Multiple documents (target: ~25 min)

- The editor starts with one empty document named `"main"`, which is active.
  All earlier methods act on the active document.
- `Create(timestamp int, name string) bool`
  Creates a new empty document. Doesn't switch to it. Returns `false` if the
  name is taken.
- `Switch(timestamp int, name string) *string`
  Makes the document active and returns its text, or `nil` if it doesn't exist
  (the active document stays the same).
- Each document has its own text, cursor, selection and undo/redo history.
  The clipboard is shared by all documents.
- `ListDocuments(timestamp int) []string`
  All documents as `"name(length)"`, sorted by text length descending, ties by
  name ascending.
