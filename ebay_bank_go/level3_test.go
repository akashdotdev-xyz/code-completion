package bank

import "testing"

func setupLevel3() *BankingSystem {
	b := NewBankingSystem()
	b.CreateAccount(1, "acc1")
	b.CreateAccount(2, "acc2")
	b.Deposit(3, "acc1", 2000)
	b.Deposit(4, "acc2", 1000)
	return b
}

func TestLevel3PayReturnsGlobalIDs(t *testing.T) {
	b := setupLevel3()
	eqStr(t, b.Pay(10, "acc1", 300), "payment1")
	eqStr(t, b.Pay(11, "acc2", 100), "payment2")
	nilStr(t, b.Pay(12, "acc2", 99999)) // insufficient: no id consumed
	nilStr(t, b.Pay(13, "nope", 1))
	eqStr(t, b.Pay(14, "acc1", 100), "payment3")
}

func TestLevel3PayDeductsBalance(t *testing.T) {
	b := setupLevel3()
	b.Pay(10, "acc1", 300)
	eqInt(t, b.Deposit(11, "acc1", 1), 1701)
}

func TestLevel3CashbackTiming(t *testing.T) {
	b := setupLevel3()
	b.Pay(10, "acc1", 1000) // cashback 20 at 10 + day
	eqInt(t, b.Deposit(10+day-1, "acc1", 1), 1001)
	eqInt(t, b.Deposit(10+day, "acc1", 1), 1022) // cashback applied first
}

func TestLevel3CashbackRoundsDown(t *testing.T) {
	b := setupLevel3()
	b.Pay(10, "acc2", 149) // 2% = 2.98 -> 2
	eqInt(t, b.Deposit(10+day, "acc2", 1), 1000-149+2+1)
}

func TestLevel3PaymentStatus(t *testing.T) {
	b := setupLevel3()
	pid := *b.Pay(10, "acc1", 500)
	eqStr(t, b.GetPaymentStatus(11, "acc1", pid), "IN_PROGRESS")
	nilStr(t, b.GetPaymentStatus(12, "acc2", pid))        // wrong account
	nilStr(t, b.GetPaymentStatus(13, "acc1", "payment9")) // no such payment
	nilStr(t, b.GetPaymentStatus(14, "nope", pid))
	eqStr(t, b.GetPaymentStatus(10+day-1, "acc1", pid), "IN_PROGRESS")
	eqStr(t, b.GetPaymentStatus(10+day, "acc1", pid), "CASHBACK_RECEIVED")
}

func TestLevel3PaymentsCountAsOutgoing(t *testing.T) {
	b := setupLevel3()
	b.Pay(10, "acc2", 400)
	b.Transfer(11, "acc1", "acc2", 300)
	b.Pay(12, "acc1", 50)
	eqList(t, b.TopSpenders(13, 2), []string{"acc2(400)", "acc1(350)"})
}

func TestLevel3CashbackNotCountedAsOutgoing(t *testing.T) {
	b := setupLevel3()
	b.Pay(10, "acc1", 1000)
	eqList(t, b.TopSpenders(10+day, 1), []string{"acc1(1000)"})
}

func TestLevel3CashbackEnablesTransfer(t *testing.T) {
	b := setupLevel3()
	b.Pay(10, "acc2", 1000) // balance 0, cashback 20 later
	nilInt(t, b.Transfer(11, "acc2", "acc1", 20))
	eqInt(t, b.Transfer(10+day, "acc2", "acc1", 20), 0)
}
