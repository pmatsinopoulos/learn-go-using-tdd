package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run(" ", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("Expected %s, actual %s", want, got)
		}
	})
}
