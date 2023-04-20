package main

import "fmt"

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amt int)  {
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amt
}

func (w Wallet) Balance() int {
	return w.balance
}

func main() {
	fmt.Println("Run go test .")
}