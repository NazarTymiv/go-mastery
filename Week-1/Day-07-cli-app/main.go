package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const (
	AllFile = "all.json"

	TypeDeposit  = "Deposit"
	TypeWithdraw = "Withdraw"
)

var loadedAll *All

type All struct {
	Account      *Account      `json:"account"`
	Pocket       *Pocket       `json:"pocket"`
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
func (acc *Account) Deposit(pok *Pocket, amount float64, transactions *[]Transaction) error {
	if amount > pok.Balance {
		return fmt.Errorf("can't deposit £%.2f to your account. you have only £%.2f in your pocket", amount, pok.Balance)
	}

	acc.Balance += amount
	pok.Balance -= amount

	*transactions = append(*transactions, Transaction{
		Type:   TypeDeposit,
		Amount: amount,
		From:   pok.Owner,
		To:     acc.Owner,
	})

	return nil
}

func (acc *Account) Withdraw(pok *Pocket, amount float64, transactions *[]Transaction) error {
	if amount > acc.Balance {
		return fmt.Errorf("can't withdraw £%.2f from your account. your account has only £%.2f", amount, acc.Balance)
	}

	acc.Balance -= amount
	pok.Balance += amount

	*transactions = append(*transactions, Transaction{
		Type:   TypeWithdraw,
		Amount: amount,
		From:   acc.Owner,
		To:     pok.Owner,
	})

	return nil
}

// All
func (a *All) Save(filename string) error {
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Helpers
func PrintTransactionsHistory(transactions *[]Transaction) {
	fmt.Println("\nTransactions History: ")

	for i, v := range *transactions {
		fmt.Printf("%d. %s £%.2f from %s to %s\n", i+1, v.Type, v.Amount, v.From, v.To)
	}
}

func PrintBalance(acc *Account, pok *Pocket) {
	fmt.Printf("\nBalance:\n - %s (account): £%.2f\n - %s (pocket): £%.2f\n",
		acc.Owner, acc.Balance, pok.Owner, pok.Balance)
}

func GetAmountInput(scanner *bufio.Scanner) (float64, error) {
	scanner.Scan()
	amountStr := scanner.Text()
	return strconv.ParseFloat(amountStr, 64)
}

func RunCLI() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1) Deposit")
		fmt.Println("2) Withdraw")
		fmt.Println("3) History")
		fmt.Println("4) Balance")
		fmt.Println("5) Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter amount to deposit: ")

			amount, err := GetAmountInput(scanner)

			if err != nil {
				fmt.Println("Invalid amount")
				break
			}

			err = loadedAll.Account.Deposit(loadedAll.Pocket, amount, &loadedAll.Transactions)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful")
			}

		case "2":
			fmt.Print("Enter amount to withdraw: ")

			amount, err := GetAmountInput(scanner)

			if err != nil {
				fmt.Println("Invalid amount")
				break
			}

			err = loadedAll.Account.Withdraw(loadedAll.Pocket, amount, &loadedAll.Transactions)

			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Withdraw successful")
			}

		case "3":
			PrintTransactionsHistory(&loadedAll.Transactions)

		case "4":
			PrintBalance(loadedAll.Account, loadedAll.Pocket)
		case "5":
			fmt.Println("Goodbye!")
			loadedAll.Save(AllFile)
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

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

func init() {
	var err error

	loadedAll, err = LoadAllFromFile(AllFile)

	if err != nil {
		fmt.Println("Error in loading all.json", err)
		loadedAll = &All{
			Account:      &Account{Owner: "Nazar", Balance: 150.0},
			Pocket:       &Pocket{Owner: "Nazar", Balance: 100.0},
			Transactions: []Transaction{},
		}
	}
}

func main() {
	defer func() {
		if err := loadedAll.Save(AllFile); err != nil {
			fmt.Println("Failed to save data: ", err)
		}
	}()

	RunCLI()
}
