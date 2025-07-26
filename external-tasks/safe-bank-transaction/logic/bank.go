package logic

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From   string
	To     string
	Amount float64
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

func (b *Bank) CreateAccount(name string, balance float64) error {
	if name == "" {
		return errors.New("Error handling the name")
	}

	b.Accounts[name] = &Account{
		Name:    name,
		Balance: balance,
	}

	log.Printf("Account %s created successfully with balance £%.2f", name, balance)

	return nil
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

	if sender.Balance < amount {
		return errors.New("Sender does not have enough money to send")
	}

	sender.Balance -= amount
	receiver.Balance += amount

	b.Transactions = append(b.Transactions, Transaction{
		To:     to,
		From:   from,
		Amount: amount,
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
	return account.Balance
}

// PrintTransactions
func (b *Bank) PrintTransactions() {
	fmt.Println("Transactions List:")

	for i, t := range b.Transactions {
		fmt.Printf("%d. %s transferred £%.2f to %s\n", i+1, t.From, t.Amount, t.To)
	}
}
