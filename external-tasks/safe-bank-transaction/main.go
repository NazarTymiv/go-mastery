package main

import (
	"fmt"
	"log"
	"sync"
)

var mutex sync.Mutex

type Account struct {
	name    string
	balance float64
}

type Transaction struct {
	from   string
	to     string
	amound float64
}

type Bank struct {
	Accounts     []Account
	Transactions []Transaction
}

func NewBank() Bank

func (b *Bank) CreateAccount(name string, balance float64) {
	if name == "" {
		log.Println("Error handling the name")
		return
	}

	b.Accounts = append(b.Accounts, Account{name: name, balance: balance})

	log.Printf("Account %s created successfully with balance £%.2f", name, balance)
}

func (b *Bank) Transfer(from string, to string, amount float64) error {
}

func main() {
	bank := NewBank()

	bank.CreateAccount("user1", 100)
	bank.CreateAccount("user2", 50)

	err := bank.Transfer("user1", "user2", 30)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bank.GetBalance("user1")) // 70
	fmt.Println(bank.GetBalance("user2")) // 80

	bank.PrintTransactions()

}
