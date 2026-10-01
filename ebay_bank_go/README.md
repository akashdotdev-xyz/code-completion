# Practice: Banking System (eBay-style CodeSignal ICA)

**Format:** 90 minutes, 4 levels, one codebase. Each level unlocks only after the
previous level passes. Don't over-engineer level 1, but pick data structures
that won't need rewriting later (hint: level 4 needs balance *history*).

Implement `BankingSystem` in `banking_system.go`. Run tests one level at a time:

```bash
cd ebay_bank_go
go test -run Level1 -v
go test -run Level2 -v
go test -run Level3 -v
go test -run Level4 -v
```

**General rules**
- Every method takes a `timestamp` (milliseconds). Timestamps across calls are
  strictly increasing.
- Amounts are positive integers.
- Return `nil` / `false` for invalid operations, as specified. Never panic. `None` below means a `nil` pointer.

---

## Level 1: Basic operations (target: ~10–15 min)

- `create_account(timestamp, account_id) -> bool`
  Creates an account. Returns `False` if it already exists.
- `deposit(timestamp, account_id, amount) -> Optional[int]`
  Adds money. Returns the new balance, or `None` if the account doesn't exist.
- `transfer(timestamp, source_id, target_id, amount) -> Optional[int]`
  Moves money from source to target. Returns the **source's** new balance.
  Returns `None` if either account doesn't exist, `source_id == target_id`, or
  the source doesn't have enough funds.

## Level 2: Ranking (target: ~15 min)

- `top_spenders(timestamp, n) -> list[str]`
  Returns the top `n` accounts by **total outgoing** money, formatted
  `"account_id(total)"`, e.g. `["acc2(500)", "acc1(300)"]`.
  - Outgoing = successful transfers out + successful payments (level 3).
  - Sort by total descending, ties by `account_id` ascending.
  - Accounts with 0 outgoing are included. If `n` exceeds the number of accounts, return all.

## Level 3: Payments with cashback (target: ~25 min)

- `pay(timestamp, account_id, amount) -> Optional[str]`
  Withdraws `amount`. Returns a payment id: `"payment1"`, `"payment2"`, …
  (a **global** counter across all accounts, counting successful payments only).
  Returns `None` if the account doesn't exist or has insufficient funds.
  - 2% cashback (rounded down) is credited to the account at
    `timestamp + 86400000` (24h later).
  - Cashback due at time `t` must be applied **before** any other operation
    executed at time `t` or later.
- `get_payment_status(timestamp, account_id, payment) -> Optional[str]`
  Returns `"IN_PROGRESS"` or `"CASHBACK_RECEIVED"`.
  Returns `None` if the account doesn't exist, the payment doesn't exist, or the
  payment belongs to a different account.

## Level 4: Merging and history (target: ~30 min)

- `merge_accounts(timestamp, account_id_1, account_id_2) -> bool`
  Merges account 2 **into** account 1:
  - account 1's balance increases by account 2's balance
  - account 1's outgoing total increases by account 2's outgoing total
  - pending cashbacks for account 2 are paid to account 1 when due
  - account 2's payments now belong to account 1 (for `get_payment_status`)
  - account 2 is deleted (its id can be re-created later as a brand-new account)

  Returns `False` if the ids are equal or either account doesn't exist.
- `get_balance(timestamp, account_id, time_at) -> Optional[int]`
  Returns the balance of `account_id` as of `time_at` (after all operations at
  `time_at`, including cashback due at or before `time_at`). `time_at <= timestamp`.
  Returns `None` if the account did not exist at `time_at` (not yet created, or
  already merged away at or before `time_at`).

---

## Go signatures

```go
CreateAccount(timestamp int, accountID string) bool
Deposit(timestamp int, accountID string, amount int) *int
Transfer(timestamp int, sourceID, targetID string, amount int) *int
TopSpenders(timestamp int, n int) []string
Pay(timestamp int, accountID string, amount int) *string
GetPaymentStatus(timestamp int, accountID, payment string) *string
MergeAccounts(timestamp int, accountID1, accountID2 string) bool
GetBalance(timestamp int, accountID string, timeAt int) *int
```

Go stdlib pieces you'll likely want: `sort.Slice`, `container/heap` (or a sorted
slice), `sort.Search`, `fmt.Sprintf`, `strconv.Itoa`.
