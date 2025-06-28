package main

import (
	"os"
	"testing"
)

const TestFile = "all_test.json"

func setup(t *testing.T) (*Account, *Pocket, *[]Transaction) {
	t.Cleanup(func() {
		loadedAll.Save(TestFile)
		if err := os.Remove(TestFile); err != nil {
			t.Logf("cleanup warning: failed to remove file %s", TestFile)
		}
	})

	loadedAll = &All{
		A:            &Account{Owner: "Nazar", Balance: 100},
		P:            &Pocket{Owner: "Nazar", Balance: 50},
		Transactions: []Transaction{},
	}

	return loadedAll.A, loadedAll.P, &loadedAll.Transactions
}

func TestDepositSuccess(t *testing.T) {
	acc, pok, tx := setup(t)

	err := acc.Deposit(pok, 50, tx)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if acc.Balance != 150 {
		t.Errorf("expected account balance to be 150, got %.2f", acc.Balance)
		return
	}

	if pok.Balance != 0 {
		t.Errorf("expected pocket balance to be 0, got %.2f", pok.Balance)
		return
	}
}

func TestDepositFail(t *testing.T) {
	acc, pok, tx := setup(t)

	err := acc.Deposit(pok, 100, tx)

	if err == nil || err.Error() != "\ncannot transfer £100.00 - pocket has only £50.00" {
		t.Errorf("expected deposit error 'cannot transfer £100.00 - pocket has only £50.00', got %v", err)
		return
	}
}

func TestWithdrawSuccess(t *testing.T) {
	acc, pok, tx := setup(t)

	err := acc.Withdraw(pok, 100, tx)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if acc.Balance != 0 {
		t.Errorf("expected account balance to be 0, got %.2f", acc.Balance)
		return
	}

	if pok.Balance != 150 {
		t.Errorf("expected pocket balance to be 150, got %.2f", pok.Balance)
		return
	}
}

func TestWithdrawFail(t *testing.T) {
	acc, pok, tx := setup(t)

	err := acc.Withdraw(pok, 150, tx)

	if err == nil || err.Error() != "\ncannot transfer £150.00 - account has only £100.00" {
		t.Errorf("expected withdraw error 'cannot transfer £150.00 - account has only £100.00', got %v", err)
		return
	}
}

func TestTransactionsSavings(t *testing.T) {
	acc, pok, tx := setup(t)

	if err := acc.Deposit(pok, 50, tx); err != nil {
		t.Fatalf("unexpected error in [TestTransactionsSavings]: %v", err)
	}

	if err := acc.Withdraw(pok, 150, tx); err != nil {
		t.Fatalf("unexpected error in [TestTransactionsSavings]: %v", err)
	}

	saveAll(t, TestFile)

	if !FileExists(TestFile) {
		t.Errorf("expected file %v to be existed", TestFile)
		return
	}

	if len(*tx) != 2 {
		t.Errorf("expected 2 records of transactions, got %v", len(*tx))
		t.Logf("transactions: %+v", *tx)
		return
	}

	loadedFromDisk, err := LoadAllFromFile(TestFile)

	if err != nil {
		t.Errorf("failed to reload file: %v", err)
	}

	if len(loadedFromDisk.Transactions) != 2 {
		t.Errorf("expected 2 transactions in saved file, got %v", len(loadedFromDisk.Transactions))
	}
}
