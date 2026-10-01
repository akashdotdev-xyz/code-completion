package library

// Library: see README.md for the full spec.
// A nil pointer return means "invalid".
type Library struct {
	// TODO: your state here
}

func NewLibrary() *Library {
	return &Library{}
}

// ---------- Level 1 ----------

func (l *Library) AddBook(timestamp int, bookID, title string) bool {
	return false
}

func (l *Library) AddMember(timestamp int, memberID string) bool {
	return false
}

func (l *Library) Borrow(timestamp int, memberID, bookID string) bool {
	return false
}

func (l *Library) Return(timestamp int, memberID, bookID string) bool {
	return false
}

// ---------- Level 2 ----------

func (l *Library) SearchByTitle(timestamp int, keyword string) []string {
	return nil
}

func (l *Library) MostBorrowed(timestamp int, n int) []string {
	return nil
}

// ---------- Level 3 ----------

func (l *Library) GetFine(timestamp int, memberID string) *int {
	return nil
}

func (l *Library) PayFine(timestamp int, memberID string, amount int) *int {
	return nil
}

// ---------- Level 4 ----------

func (l *Library) Reserve(timestamp int, memberID, bookID string) bool {
	return false
}

func (l *Library) GetWaitlist(timestamp int, bookID string) []string {
	return nil
}
