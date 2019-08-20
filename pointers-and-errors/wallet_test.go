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

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("expected an error, but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("unexpect error!")
		}
	}

	t.Run("#Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertCorrectMessage(t, wallet, Bitcoin(10))
	})

	t.Run("#Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertCorrectMessage(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("# withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertCorrectMessage(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
