package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("expected and error, but didn't get one")
		}
	}

	t.Run("#Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertCorrectMessage(t, wallet, Bitcoin(10))

	})

	t.Run("#Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertCorrectMessage(t, wallet, Bitcoin(10))
	})

	t.Run("# withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertCorrectMessage(t, wallet, startingBalance)
		assertError(t, err)
	})
}
