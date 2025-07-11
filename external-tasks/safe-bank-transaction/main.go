package main

import (
	"fmt"
	"log"

	"github.com/nazartymiv/go-mastery/external-tasks/safe-bank-transaction/logic"
)

func main() {
	bank := logic.NewBank()

	bank.CreateAccount("user1", 100)
	bank.CreateAccount("user2", 50)

	err := bank.Transfer("user1", "user2", 30)
	if err != nil {
		log.Println(err)
	}
	err = bank.Transfer("user1", "user2", 30)
	if err != nil {
		log.Println(err)
	}
	err = bank.Transfer("user1", "user2", 30)
	if err != nil {
		log.Println(err)
	}
	err = bank.Transfer("user1", "user2", 30)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(bank.GetBalance("user1")) // 70
	fmt.Println(bank.GetBalance("user2")) // 80

	bank.PrintTransactions()
}
