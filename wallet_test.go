package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if wallet.Balance() != want {
			t.Errorf("Expected %s, got %s", want, got)
		}
	}

	assertError := func(t testing.TB, err error, want string) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
		if err.Error() != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(30))

		assertError(t, err, "not enough balance")
		assertBalance(t, wallet, startingBalance)
	})
}
