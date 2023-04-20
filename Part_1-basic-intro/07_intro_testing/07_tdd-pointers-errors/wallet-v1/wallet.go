package main

import "fmt"

// Wallet represents an electronic wallet that keeps track of the amount of money you have
type Wallet struct {
	balance int
}

// Deposit lets you increase the balance of your wallet by the given amount
func (w *Wallet) Deposit(amt int)  {
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amt
}

// Balance returns the current balance of your wallet
func (w *Wallet) Balance() int {
	return w.balance
}

func main() {
	fmt.Println("Run go test .")
}