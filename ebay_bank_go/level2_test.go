package bank

import "testing"

func setupLevel2() *BankingSystem {
	b := NewBankingSystem()
	for i, acc := range []string{"acc1", "acc2", "acc3"} {
		b.CreateAccount(i+1, acc)
		b.Deposit(11+i, acc, 1000)
	}
	return b
}

func TestLevel2AllZeroSortedByID(t *testing.T) {
	b := setupLevel2()
	eqList(t, b.TopSpenders(20, 3), []string{"acc1(0)", "acc2(0)", "acc3(0)"})
}

func TestLevel2Ranking(t *testing.T) {
	b := setupLevel2()
	b.Transfer(20, "acc1", "acc2", 300)
	b.Transfer(21, "acc3", "acc1", 500)
	b.Transfer(22, "acc1", "acc3", 100)
	// acc1: 400 out, acc3: 500 out, acc2: 0
	eqList(t, b.TopSpenders(23, 2), []string{"acc3(500)", "acc1(400)"})
	eqList(t, b.TopSpenders(24, 3), []string{"acc3(500)", "acc1(400)", "acc2(0)"})
}

func TestLevel2TiesBrokenByID(t *testing.T) {
	b := setupLevel2()
	b.Transfer(20, "acc3", "acc1", 200)
	b.Transfer(21, "acc2", "acc1", 200)
	eqList(t, b.TopSpenders(22, 3), []string{"acc2(200)", "acc3(200)", "acc1(0)"})
}

func TestLevel2NLargerThanAccounts(t *testing.T) {
	b := setupLevel2()
	b.Transfer(20, "acc2", "acc1", 50)
	eqList(t, b.TopSpenders(21, 10), []string{"acc2(50)", "acc1(0)", "acc3(0)"})
}

func TestLevel2FailedTransferNotCounted(t *testing.T) {
	b := setupLevel2()
	b.Transfer(20, "acc1", "acc2", 5000) // insufficient
	eqList(t, b.TopSpenders(21, 1), []string{"acc1(0)"})
}

func TestLevel2IncomingNotCounted(t *testing.T) {
	b := setupLevel2()
	b.Transfer(20, "acc1", "acc2", 700)
	eqList(t, b.TopSpenders(21, 2), []string{"acc1(700)", "acc2(0)"})
}
