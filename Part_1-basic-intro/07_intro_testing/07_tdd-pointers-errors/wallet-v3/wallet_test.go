package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	fmt.Printf("address of balance in TestWaller: %v\n", &wallet.balance)

	got := wallet.Balance()
	want := Bitcoin(10)
	if got != want {
		t.Errorf("got %s, but wanted %s", got, want)
	}
}