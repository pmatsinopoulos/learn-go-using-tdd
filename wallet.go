package main

type Wallet struct {
	balance int
}

func (wallet *Wallet) Deposit(amount int) {
	wallet.balance += amount
}

func (wallet Wallet) Balance() (balance int) {
	balance = wallet.balance

	return
}
