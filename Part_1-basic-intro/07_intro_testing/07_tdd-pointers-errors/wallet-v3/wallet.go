package main

import "fmt"

// Bitcoin represents an integer amount of Bitcoins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("₿%d", b)
}

// Wallet represents an electronic wallet that keeps track of the amount of money you have
type Wallet struct {
	balance Bitcoin
}

// Deposit lets you increase the balance of your wallet by the given amount
func (w *Wallet) Deposit(amt Bitcoin)  {
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amt
}

// Balance returns the current balance of your wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	fmt.Println("Run go test .")
}