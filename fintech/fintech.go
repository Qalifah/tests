package fintech

import (
	"errors"
	"fmt"
)

type ( 
	// Wallet stores info about a user savings
	Wallet struct {
		balance Bitcoin
	}

	// Bitcoin represents the user's money
	Bitcoin int
)

// ErrInsufficientFunds represent an insufficient funds error
var ErrInsufficientFunds = errors.New("can't withdraw, insufficient funds")

// Deposit adds amount to a user current balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the amount a user has
func(w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw subtracts amount from the user's current balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}