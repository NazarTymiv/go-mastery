package main

import (
	"fmt"
)

type BalanceHolder interface {
	GetBalance() string
}

type Account struct {
	Owner   string
	Balance float64
}

type Pocket struct {
	Balance float64
}

func (a *Account) Deposit(p *Pocket, amount float64) error {
	if p.Balance < amount {
		return fmt.Errorf("cannot transfer £%.2f — pocket only has £%.2f", amount, p.Balance)
	}
	a.Balance += amount
	p.Balance -= amount
	return nil
}

func (a *Account) Withdraw(p *Pocket, amount float64) error {
	if a.Balance < amount {
		return fmt.Errorf("cannot transfer £%.2f - account only has £%.2f", amount, a.Balance)
	}
	a.Balance -= amount
	p.Balance += amount
	return nil
}

func (a *Account) GetBalance() string {
	return fmt.Sprintf("Account (%s): £%.2f", a.Owner, a.Balance)
}

func (p *Pocket) GetBalance() string {
	return fmt.Sprintf("Pocket: £%.2f", p.Balance)
}

func PrintBalances(items ...BalanceHolder) {
	for _, item := range items {
		fmt.Println(item.GetBalance())
	}
	fmt.Println("=============")
}

func PrintError(err error) {
	fmt.Println(err)
	fmt.Println("=============")
}

func main() {
	acc := &Account{Owner: "Nazar", Balance: 100}
	pocket := &Pocket{Balance: 100}

	if err := acc.Deposit(pocket, 50); err != nil {
		PrintError(err)
	} else {
		PrintBalances(acc, pocket)
	}

	if err := acc.Withdraw(pocket, 150); err != nil {
		PrintError(err)

	} else {
		PrintBalances(acc, pocket)
	}

	if err := acc.Withdraw(pocket, 50); err != nil {
		PrintError(err)

	} else {
		PrintBalances(acc, pocket)
	}

	if err := acc.Deposit(pocket, 200); err != nil {
		PrintError(err)

	} else {
		PrintBalances(acc, pocket)
	}
}
