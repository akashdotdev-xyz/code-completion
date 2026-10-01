package bank

import "testing"

func TestLevel1CreateAccount(t *testing.T) {
	b := NewBankingSystem()
	eqBool(t, b.CreateAccount(1, "acc1"), true)
	eqBool(t, b.CreateAccount(2, "acc2"), true)
	eqBool(t, b.CreateAccount(3, "acc1"), false)
}

func TestLevel1Deposit(t *testing.T) {
	b := NewBankingSystem()
	b.CreateAccount(1, "acc1")
	eqInt(t, b.Deposit(2, "acc1", 100), 100)
	eqInt(t, b.Deposit(3, "acc1", 250), 350)
	nilInt(t, b.Deposit(4, "missing", 100))
}

func TestLevel1TransferSuccess(t *testing.T) {
	b := NewBankingSystem()
	b.CreateAccount(1, "acc1")
	b.CreateAccount(2, "acc2")
	b.Deposit(3, "acc1", 500)
	eqInt(t, b.Transfer(4, "acc1", "acc2", 200), 300)
	eqInt(t, b.Deposit(5, "acc2", 1), 201)
}

func TestLevel1TransferExactBalance(t *testing.T) {
	b := NewBankingSystem()
	b.CreateAccount(1, "acc1")
	b.CreateAccount(2, "acc2")
	b.Deposit(3, "acc1", 100)
	eqInt(t, b.Transfer(4, "acc1", "acc2", 100), 0)
}

func TestLevel1TransferFailures(t *testing.T) {
	b := NewBankingSystem()
	b.CreateAccount(1, "acc1")
	b.CreateAccount(2, "acc2")
	b.Deposit(3, "acc1", 100)
	nilInt(t, b.Transfer(4, "acc1", "acc2", 101)) // insufficient
	nilInt(t, b.Transfer(5, "acc1", "acc1", 10))  // same account
	nilInt(t, b.Transfer(6, "acc1", "nope", 10))  // missing target
	nilInt(t, b.Transfer(7, "nope", "acc1", 10))  // missing source
	// failed transfers must not change balances
	eqInt(t, b.Deposit(8, "acc1", 1), 101)
	eqInt(t, b.Deposit(9, "acc2", 1), 1)
}
