# Practice: Library System (CodeSignal ICA style)

**Format:** 90 minutes, 4 levels, one codebase. Each level unlocks only after the
previous level passes. Level 3 needs to know when each loan started, and
level 4 needs to lend a book to someone automatically when it's returned, so
put the "can this member borrow right now?" check in one helper.

Implement `Library` in `library.go`. Run tests one level at a time:

```bash
cd library_go
go test -run Level1 -v
go test -run Level2 -v
go test -run Level3 -v
go test -run Level4 -v
```

**General rules**
- Every method takes a `timestamp` (milliseconds). Timestamps across calls are
  strictly increasing.
- `DAY = 86400000` milliseconds.
- A `nil` return means "invalid". Never panic.

---

## Level 1: Books and loans (target: ~10 min)

- `AddBook(timestamp int, bookID, title string) bool`
  Returns `false` if the book ID already exists.
- `AddMember(timestamp int, memberID string) bool`
  Returns `false` if the member already exists.
- `Borrow(timestamp int, memberID, bookID string) bool`
  Lends the book to the member. Returns `false` if the member or book doesn't
  exist, or the book is already lent out. A member may hold several books.
- `Return(timestamp int, memberID, bookID string) bool`
  Returns `false` if this member doesn't currently have this book.

## Level 2: Search and stats (target: ~15 min)

- `SearchByTitle(timestamp int, keyword string) []string`
  IDs of books whose title contains `keyword`, ignoring case, sorted by book ID
  (plain string order, so `"b10"` comes before `"b3"`).
- `MostBorrowed(timestamp int, n int) []string`
  Up to `n` books as `"bookID(count)"`, where count is how many times the book
  has been lent out in total. Sort by count descending, ties by book ID
  ascending. Books never borrowed are included with `(0)`.
- Both return an empty slice (not `nil`) when there's nothing to return.

## Level 3: Due dates and fines (target: ~25 min)

- Every loan is due `14 * DAY` after it starts.
- Returning a book **after** its due time adds a fine of `10` for every day or
  part of a day it's late: `10 * ceil((returnTime - dueTime) / DAY)`.
  Returning exactly at the due time is not late.
- `GetFine(timestamp int, memberID string) *int`
  The member's unpaid fines, or `nil` if the member doesn't exist.
- `PayFine(timestamp int, memberID string, amount int) *int`
  Pays off part of the fine and returns what's left. Returns `nil` if the member
  doesn't exist or `amount` is more than they owe.
- A member **can't borrow** while they have unpaid fines, or while they hold any
  book whose due time is before `timestamp`. `Borrow` returns `false` in that case.

## Level 4: Waitlists (target: ~30 min)

- `Reserve(timestamp int, memberID, bookID string) bool`
  Adds the member to the end of the book's waitlist. Returns `false` if the
  member or book doesn't exist, the book isn't currently lent out, the member is
  the one holding it, or the member is already on its waitlist.
- When a book with a waitlist is returned, it is immediately lent to the first
  member on the waitlist who is **allowed to borrow at that moment** (level 3
  rules). Members before them who aren't allowed are removed from the waitlist.
  The new loan starts at the return's timestamp and counts in `MostBorrowed`.
  If nobody on the waitlist is allowed, the waitlist ends up empty and the book
  is available.
- `GetWaitlist(timestamp int, bookID string) []string`
  Member IDs on the book's waitlist, in order. Empty slice if there are none or
  the book doesn't exist.
