package main

import (
	"errors"
	"fmt"
)

//ErrInsufficientFunds is insufficient funds error
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Bitcoin type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet a wallet
type Wallet struct {
	balance Bitcoin
}

//Deposit money into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance  get balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Withdraw from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
