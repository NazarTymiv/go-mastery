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

type All struct {
	A *Account `json:"account"`
	P *Pocket  `json:"pocket"`
}

type Account struct {
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

type Pocket struct {
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

func (a *Account) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (a *Account) Deposit(p *Pocket, amount float64) error {
	if p.Balance < amount {
		return fmt.Errorf("cannot transfer £%.2f - pocket has only £%.2f", amount, p.Balance)
	}

	fmt.Printf("Depositing £%.2f to %s\n", amount, a.Owner)

	a.Balance += amount
	p.Balance -= amount

	if err := a.SaveToFile(AccountFile); err != nil {
		fmt.Println("Save error:", err)
	}

	return nil
}

func (a *Account) Withdraw(p *Pocket, amount float64) error {
	if a.Balance < amount {
		return fmt.Errorf("cannot transfer £%.2f - account has only £%.2f", amount, a.Balance)
	}

	fmt.Printf("Withdrawing £%.2f to %s\n", amount, p.Owner)

	a.Balance -= amount
	p.Balance += amount

	if err := a.SaveToFile(AccountFile); err != nil {
		fmt.Println("Save error:", err)
	}

	return nil
}

func (p *Pocket) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(p, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func LoadAccountFromFile(filename string) (*Account, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var acc Account
	if err := json.Unmarshal(data, &acc); err != nil {
		return nil, err
	}

	return &acc, nil
}

func LoadPocketFromFile(filename string) (*Pocket, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var pok Pocket
	if err := json.Unmarshal(data, &pok); err != nil {
		return nil, err
	}

	return &pok, nil
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

func SaveAllToFile(filename string, a *Account, p *Pocket) error {
	all := All{A: a, P: p}
	data, err := json.MarshalIndent(all, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func PrintLoadedAll() {
	loadedAll, err := LoadAllFromFile(AllFile)
	if err != nil {
		println("Load error:", err)
		return
	}

	fmt.Printf("Loaded account: Owner=%s, Balance=£%.2f\nLoaded Pocket: Owner=%s, Balance=£%.2f\n", loadedAll.A.Owner, loadedAll.A.Balance, loadedAll.P.Owner, loadedAll.P.Balance)
	fmt.Println("==================")
}

func main() {
	// Initial account
	acc := &Account{Owner: "Nazar", Balance: 150.0}
	pok := &Pocket{Owner: "Nazar", Balance: 100}

	if err := acc.Deposit(pok, 50); err != nil {
		fmt.Println(err)
	}
	SaveAllToFile(AllFile, acc, pok)
	PrintLoadedAll()

	if err := acc.Deposit(pok, 50); err != nil {
		fmt.Println(err)
	}
	SaveAllToFile(AllFile, acc, pok)
	PrintLoadedAll()

	if err := acc.Deposit(pok, 50); err != nil {
		fmt.Println(err)
	}
	SaveAllToFile(AllFile, acc, pok)
	PrintLoadedAll()

	if err := acc.Withdraw(pok, 250); err != nil {
		fmt.Println(err)
	}
	SaveAllToFile(AllFile, acc, pok)
	PrintLoadedAll()

	if err := acc.Withdraw(pok, 1); err != nil {
		fmt.Println(err)
	}
	SaveAllToFile(AllFile, acc, pok)
	PrintLoadedAll()
}
