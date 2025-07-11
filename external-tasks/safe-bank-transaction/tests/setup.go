package tests

import (
	"github.com/nazartymiv/go-mastery/external-tasks/safe-bank-transaction/logic"
)

const (
	User1 = "user1"
	User2 = "user2"
)

func setup() *logic.Bank {
	bank := logic.NewBank()

	bank.CreateAccount(User1, 100)
	bank.CreateAccount(User2, 50)

	return &bank
}
