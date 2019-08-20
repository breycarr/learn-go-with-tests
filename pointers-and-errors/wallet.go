package main

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is an error message for when there is insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin represents a Bitcoin
type Bitcoin int

// Stringer allows types that accept the String method
type Stringer interface {
	String() string
}

// String formats the Bitcoins into a string with BTC currency code
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet contains a balance
type Wallet struct {
	balance Bitcoin
}

// Deposit adds an amount to the balance in the Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw deducts an amount from the balance in the Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance shows the current balance of the Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {

}
