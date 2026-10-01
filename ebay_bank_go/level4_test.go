package bank

import "testing"

func setupLevel4() *BankingSystem {
	b := NewBankingSystem()
	b.CreateAccount(1, "acc1")
	b.CreateAccount(2, "acc2")
	b.Deposit(3, "acc1", 1000)
	b.Deposit(4, "acc2", 500)
	return b
}

func TestLevel4MergeInvalid(t *testing.T) {
	b := setupLevel4()
	eqBool(t, b.MergeAccounts(10, "acc1", "acc1"), false)
	eqBool(t, b.MergeAccounts(11, "acc1", "nope"), false)
	eqBool(t, b.MergeAccounts(12, "nope", "acc1"), false)
}

func TestLevel4MergeBalancesAndDeletes(t *testing.T) {
	b := setupLevel4()
	eqBool(t, b.MergeAccounts(10, "acc1", "acc2"), true)
	eqInt(t, b.Deposit(11, "acc1", 1), 1501)
	nilInt(t, b.Deposit(12, "acc2", 1))
	eqBool(t, b.CreateAccount(13, "acc2"), true) // id reusable
	eqInt(t, b.Deposit(14, "acc2", 7), 7)
}

func TestLevel4MergeOutgoing(t *testing.T) {
	b := setupLevel4()
	b.CreateAccount(5, "acc3")
	b.Transfer(6, "acc1", "acc3", 100)
	b.Pay(7, "acc2", 300)
	b.MergeAccounts(10, "acc1", "acc2")
	eqList(t, b.TopSpenders(11, 5), []string{"acc1(400)", "acc3(0)"})
}

func TestLevel4MergeMovesPendingCashbackAndPayments(t *testing.T) {
	b := setupLevel4()
	pid := *b.Pay(10, "acc2", 500) // acc2 -> 0, cashback 10 at 10 + day
	b.MergeAccounts(20, "acc1", "acc2")
	eqStr(t, b.GetPaymentStatus(21, "acc1", pid), "IN_PROGRESS")
	nilStr(t, b.GetPaymentStatus(22, "acc2", pid))
	eqInt(t, b.Deposit(10+day, "acc1", 1), 1011)
	eqStr(t, b.GetPaymentStatus(10+day+1, "acc1", pid), "CASHBACK_RECEIVED")
}

func TestLevel4GetBalanceHistory(t *testing.T) {
	b := setupLevel4()
	b.Transfer(10, "acc1", "acc2", 200)    // acc1=800, acc2=700
	b.Deposit(20, "acc1", 50)              // acc1=850
	nilInt(t, b.GetBalance(30, "acc1", 0)) // before creation
	eqInt(t, b.GetBalance(31, "acc1", 1), 0)
	eqInt(t, b.GetBalance(32, "acc1", 3), 1000)
	eqInt(t, b.GetBalance(33, "acc1", 9), 1000)
	eqInt(t, b.GetBalance(34, "acc1", 10), 800)
	eqInt(t, b.GetBalance(35, "acc1", 25), 850)
	eqInt(t, b.GetBalance(36, "acc2", 15), 700)
	nilInt(t, b.GetBalance(37, "nope", 15))
}

func TestLevel4GetBalanceWithCashback(t *testing.T) {
	b := setupLevel4()
	b.Pay(10, "acc1", 1000) // acc1=0, cashback 20 due at 10 + day
	due := 10 + day
	b.Deposit(due+5, "acc2", 1) // any op after due time triggers processing
	eqInt(t, b.GetBalance(due+6, "acc1", due-1), 0)
	eqInt(t, b.GetBalance(due+7, "acc1", due), 20)
}

func TestLevel4GetBalanceAcrossMerge(t *testing.T) {
	b := setupLevel4()
	b.MergeAccounts(10, "acc1", "acc2")
	eqInt(t, b.GetBalance(11, "acc2", 9), 500) // history kept before merge
	nilInt(t, b.GetBalance(12, "acc2", 10))    // gone at merge time
	eqInt(t, b.GetBalance(13, "acc1", 9), 1000)
	eqInt(t, b.GetBalance(14, "acc1", 10), 1500)
	b.CreateAccount(20, "acc2")
	nilInt(t, b.GetBalance(21, "acc2", 15))
	eqInt(t, b.GetBalance(22, "acc2", 20), 0)
}
