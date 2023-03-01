package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1	// non-blocking
	ch <- 2	// blocking from this point blocking because buffer is full
	// ch <- 3	// This will trigger an error because buffer is full
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	// fmt.Println(<- ch) // This will trigger an error because buffer is empty
}