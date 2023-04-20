package main

import "testing"


func TestWallet(t *testing.T) {

	assertBalance := func (t testing.TB, w *Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s, but wanted %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("Withdraw (no overdraft)", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(10)
		assertBalance(t, &wallet, Bitcoin(10))
		if err != nil {
			t.Errorf("expected no errors, but got %v", err)
		}
	})

	t.Run("Withdraw (overdraft)", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(31)
		assertBalance(t, &wallet, Bitcoin(20))
		wantErrString := "Wallet.Withdraw: insufficient funds"
		if err == nil {
			t.Errorf("expected an error, but didn't get one")
		}
		if err.Error() != wantErrString {
			t.Errorf("got err message %q, but wanted %q", err.Error(), wantErrString)
		}
	})
}