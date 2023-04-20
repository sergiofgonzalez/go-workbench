package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("Withdraw (no overdraft)", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(10)
		assertBalance(t, &wallet, Bitcoin(10))
		assertError(t, err, nil)
	})

	t.Run("Withdraw (overdraft)", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(31)
		assertBalance(t, &wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, w *Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("unexpected balance: got %s, but wanted %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("unexpected error: got %v, but wanted %v", got, want)
	}
}