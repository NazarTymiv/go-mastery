package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Account struct {
	name    string
	balance float64
}

type Transaction struct {
	from   string
	to     string
	amount float64
}

type Bank struct {
	Accounts     map[string]*Account
	Transactions []Transaction
	mutex        sync.Mutex
}

func NewBank() Bank {
	return Bank{
		Accounts:     map[string]*Account{},
		Transactions: []Transaction{},
	}
}

func (b *Bank) CreateAccount(name string, balance float64) {
	if name == "" {
		log.Println("Error handling the name")
		return
	}

	b.Accounts[name] = &Account{
		name:    name,
		balance: balance,
	}

	log.Printf("Account %s created successfully with balance £%.2f", name, balance)
}

// Make transaction
func (b *Bank) Transfer(from string, to string, amount float64) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	sender := b.Accounts[from]
	receiver := b.Accounts[to]
	if sender == nil || receiver == nil {
		return errors.New("Account not found")
	}

	if sender.balance < amount {
		return errors.New("Sender does not have enough money to send")
	}

	sender.balance -= amount
	receiver.balance += amount

	b.Transactions = append(b.Transactions, Transaction{
		to:     to,
		from:   from,
		amount: amount,
	})

	return nil
}

// Get balance by username
func (b *Bank) GetBalance(name string) float64 {
	account := b.Accounts[name]
	if account == nil {
		log.Printf("Account %s not found", name)
		return 0
	}
	return account.balance
}

// PrintTransactions
func (b *Bank) PrintTransactions() {
	fmt.Println("Transactions List:")

	for i, t := range b.Transactions {
		fmt.Printf("%d. %s transferred £%.2f to %s\n", i+1, t.from, t.amount, t.to)
	}
}

func main() {
	bank := NewBank()

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
