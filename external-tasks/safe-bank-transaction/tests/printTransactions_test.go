package tests

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintTransaction(t *testing.T) {
	bank := setup()

	bank.Transfer(User1, User2, 30)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	bank.PrintTransactions()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	got := buf.String()

	expected := "Transactions List:\n1. user1 transferred £30.00 to user2\n"

	if got != expected {
		t.Errorf("unexpected output:\nGot:\n%s\nWant:\n%s", got, expected)
	}

}
