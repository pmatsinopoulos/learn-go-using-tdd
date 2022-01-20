package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return errors.New("not enough balance")
	}
	wallet.balance -= amount
	return nil
}

func (wallet Wallet) Balance() (balance Bitcoin) {
	balance = wallet.balance

	return
}
