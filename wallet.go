package main

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet Wallet) Balance() (balance Bitcoin) {
	balance = wallet.balance

	return
}
