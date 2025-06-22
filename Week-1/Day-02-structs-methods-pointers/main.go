package main

import (
	"fmt"
)

type Pocket struct {
	Balance float64
}

type Account struct {
	Owner   string
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
		return fmt.Errorf("cannot transfer £%.2f — account only has £%.2f", amount, a.Balance)
	}
	a.Balance -= amount
	p.Balance += amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func main() {
	acc := &Account{Owner: "Nazar", Balance: 100}
	pocket := &Pocket{Balance: 100}

	if err := acc.Deposit(pocket, 50); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Account Balance: £%.2f\n", acc.GetBalance())
	fmt.Printf("Pocket Balance: £%.2f\n", pocket.Balance)

	if err := acc.Withdraw(pocket, 30); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Account Balance: £%.2f\n", acc.GetBalance())
	fmt.Printf("Pocket Balance: £%.2f\n", pocket.Balance)
}
