package bank

import (
	"fmt"
	"sort"
)

// BankingSystem: see README.md for the full spec.
// A nil pointer return means "invalid operation" (None in the Python version).

const (
	DAY = 86400000
)

type Payment struct {
	ID      string
	owner   string
	amount  int // cashback amount
	dueAt   int
	applied bool
}

type Account struct {
	ID     string
	Amount int
	Spend  int
}

func NewAccount(id string) *Account {
	return &Account{
		ID: id,
	}
}

type BankingSystem struct {
	accounts map[string]*Account
	payments map[string]*Payment
	pending  []*Payment
	counter  int
}

func NewBankingSystem() *BankingSystem {
	return &BankingSystem{
		accounts: make(map[string]*Account),
		payments: make(map[string]*Payment),
		pending:  make([]*Payment, 0),
		counter:  0,
	}
}

// ---------- Level 1 ----------

func (b *BankingSystem) applyCashback(timestamp int) {

	i := 0
	for i < len(b.pending) && b.pending[i].dueAt <= timestamp {
		payment := b.pending[i]

		b.pending = b.pending[1:]
		payment.applied = true

		b.accounts[payment.owner].Amount += payment.amount
		b.payments[payment.owner] = payment
	}
}

func (b *BankingSystem) CreateAccount(timestamp int, accountID string) bool {

	if _, ok := b.accounts[accountID]; ok {
		return false
	}

	b.accounts[accountID] = NewAccount(accountID)
	return true
}

func (b *BankingSystem) Deposit(timestamp int, accountID string, amount int) *int {
	b.applyCashback(timestamp)
	account, ok := b.accounts[accountID]
	if !ok {
		return nil
	}

	account.Amount += amount
	b.accounts[accountID] = account
	return &account.Amount
}

func (b *BankingSystem) Transfer(timestamp int, sourceID, targetID string, amount int) *int {
	b.applyCashback(timestamp)
	if sourceID == targetID {
		return nil
	}

	sourceAccount, ok := b.accounts[sourceID]
	if !ok {
		return nil
	}

	targetAccount, ok := b.accounts[targetID]
	if !ok {
		return nil
	}

	if sourceAccount.Amount < amount {
		return nil
	}

	sourceAccount.Amount -= amount
	targetAccount.Amount += amount
	b.accounts[sourceID] = sourceAccount
	b.accounts[targetID] = targetAccount

	// track spends
	sourceAccount.Spend += amount

	return &sourceAccount.Amount
}

// ---------- Level 2 ----------

func (b *BankingSystem) TopSpenders(timestamp int, n int) []string {

	var currAccounts []Account
	for _, value := range b.accounts {
		currAccounts = append(currAccounts, *value)
	}

	sort.Slice(currAccounts, func(i, j int) bool {
		if currAccounts[i].Spend == currAccounts[j].Spend {
			return currAccounts[i].ID < currAccounts[j].ID
		}
		return currAccounts[i].Spend > currAccounts[j].Spend
	})

	var result []string
	for i := 0; i < min(n, len(currAccounts)); i++ {
		account := currAccounts[i]
		result = append(result, account.ID+fmt.Sprintf("(%d)", account.Spend))
	}
	return result
}

// ---------- Level 3 ----------

func (b *BankingSystem) Pay(timestamp int, accountID string, amount int) *string {

	account, ok := b.accounts[accountID]
	if !ok {
		return nil
	}

	if account.Amount < amount {
		return nil
	}

	b.counter += 1
	account.Amount -= amount
	account.Spend += amount
	paymentID := fmt.Sprintf("payment%d", b.counter)

	payment := &Payment{
		ID:      paymentID,
		owner:   accountID,
		dueAt:   timestamp + DAY,
		applied: false,
		amount:  (amount * 2) / 100,
	}
	b.pending = append(b.pending, payment)
	b.payments[paymentID] = payment

	return &paymentID
}

func (b *BankingSystem) GetPaymentStatus(timestamp int, accountID, payment string) *string {

	paymentDetails, ok := b.payments[payment]
	if !ok {
		return nil
	}

	if paymentDetails.owner != accountID {
		return nil
	}

	inProgress := "IN_PROGRESS"
	success := "CASHBACK_RECEIVED"
	if paymentDetails.dueAt > timestamp {
		return &inProgress
	}
	return &success
}

// ---------- Level 4 ----------

func (b *BankingSystem) MergeAccounts(timestamp int, accountID1, accountID2 string) bool {
	return false
}

func (b *BankingSystem) GetBalance(timestamp int, accountID string, timeAt int) *int {
	return nil
}

/*
b := NewBankingSystem()
	b.CreateAccount(1, "acc1")
	b.CreateAccount(2, "acc2")
	b.Deposit(3, "acc1", 2000)
	b.Deposit(4, "acc2", 1000)

	pid := *b.Pay(10, "acc1", 500)
	eqStr(t, b.GetPaymentStatus(11, "acc1", pid), "IN_PROGRESS")
	nilStr(t, b.GetPaymentStatus(12, "acc2", pid))        // wrong account
	nilStr(t, b.GetPaymentStatus(13, "acc1", "payment9")) // no such payment
	nilStr(t, b.GetPaymentStatus(14, "nope", pid))
	eqStr(t, b.GetPaymentStatus(10+day-1, "acc1", pid), "IN_PROGRESS")
	eqStr(t, b.GetPaymentStatus(10+day, "acc1", pid), "CASHBACK_RECEIVED")
*/
