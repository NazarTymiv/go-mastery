package tests

import "testing"

func TestGetBalance(t *testing.T) {
	bank := setup()

	balanceUser1 := bank.GetBalance(User1)
	expectedBalanceUser1 := 100.00

	balanceUser2 := bank.GetBalance(User2)
	expectedBalanceUser2 := 50.00

	if balanceUser1 != float64(expectedBalanceUser1) {
		t.Errorf("Expected balance for %s %.2f, got %.2f", User1, expectedBalanceUser1, balanceUser1)
	}

	if balanceUser2 != float64(expectedBalanceUser2) {
		t.Errorf("Expected balance for %s %.2f, got %.2f", User2, expectedBalanceUser2, balanceUser2)
	}
}
