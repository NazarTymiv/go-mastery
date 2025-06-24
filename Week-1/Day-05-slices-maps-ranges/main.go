package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	AccountFile = "account.json"
	PocketFile  = "pocket.json"
	AllFile     = "all.json"
)

var loadedAll *All

type All struct {
	A            *Account      `json:"account"`
	P            *Pocket       `json:"pocket"`
	Transactions []Transaction `json:"transactions"`
}

type Account struct {
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

type Pocket struct {
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

// Account
func (a *Account) Deposit(p *Pocket, amount float64, history *[]Transaction) error {
	if p.Balance < amount {
		return fmt.Errorf("\ncannot transfer £%.2f - pocket has only £%.2f", amount, p.Balance)
	}

	fmt.Printf("\nDepositing £%.2f to %s\n", amount, a.Owner)

	a.Balance += amount
	p.Balance -= amount

	*history = append(*history, Transaction{
		Type:   "deposit",
		Amount: amount,
		From:   p.Owner,
		To:     a.Owner,
	})

	return nil
}

func (a *Account) Withdraw(p *Pocket, amount float64, history *[]Transaction) error {
	if a.Balance < amount {
		return fmt.Errorf("\ncannot transfer £%.2f - account has only £%.2f", amount, a.Balance)
	}

	fmt.Printf("\nWithdrawing £%.2f to %s\n", amount, p.Owner)

	a.Balance -= amount
	p.Balance += amount

	*history = append(*history, Transaction{
		Type:   "withdraw",
		Amount: amount,
		From:   p.Owner,
		To:     a.Owner,
	})

	return nil
}

// All
func (a *All) Save(filename string) error {
	data, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Global
func LoadAllFromFile(filename string) (*All, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var all All
	if err := json.Unmarshal(data, &all); err != nil {
		return nil, err
	}

	return &all, nil
}

func PrintHistory(history *[]Transaction) {
	fmt.Println("\nTransaction History:")

	for i, t := range *history {
		fmt.Printf("%d. %s £%.2f from %s to %s\n", i+1, t.Type, t.Amount, t.From, t.To)
	}
}

func PrintSummary(acc *Account, pok *Pocket) {
	balances := map[string]float64{
		acc.Owner + " (account)": acc.Balance,
		pok.Owner + " (pocket)":  pok.Balance,
	}

	fmt.Println("Summary:")
	for k, v := range balances {
		fmt.Printf(" - %s: £%.2f\n", k, v)
	}
}

func ExecuteAndPrint(opName string, err error, acc *Account, pok *Pocket) {
	if err != nil {
		fmt.Printf("\n[%s Failed] %v\n", opName, err)
	} else {
		PrintSummary(acc, pok)
	}
}

func init() {
	var err error

	loadedAll, err = LoadAllFromFile(AllFile)

	if err != nil {
		fmt.Println("Error in loading all.json", err)
		loadedAll = &All{
			A:            &Account{Owner: "Nazar", Balance: 150.0},
			P:            &Pocket{Owner: "Nazar", Balance: 100.0},
			Transactions: []Transaction{},
		}
	}
}

func main() {
	acc := loadedAll.A
	pok := loadedAll.P
	transactions := &loadedAll.Transactions

	err := acc.Deposit(pok, 50, transactions)
	ExecuteAndPrint("Deposit", err, acc, pok)

	err = acc.Deposit(pok, 50, transactions)
	ExecuteAndPrint("Deposit", err, acc, pok)

	err = acc.Deposit(pok, 50, transactions)
	ExecuteAndPrint("Deposit", err, acc, pok)

	err = acc.Withdraw(pok, 250, transactions)
	ExecuteAndPrint("Withdraw", err, acc, pok)

	err = acc.Withdraw(pok, 1, transactions)
	ExecuteAndPrint("Withdraw", err, acc, pok)

	defer loadedAll.Save(AllFile)
	PrintHistory(transactions)
}
