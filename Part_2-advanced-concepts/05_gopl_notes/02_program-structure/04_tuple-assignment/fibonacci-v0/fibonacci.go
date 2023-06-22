package main

import "fmt"


func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x + y
	}
	return x
}

func main() {
	fmt.Println(fib(0))	// 0
	fmt.Println(fib(1))	// 1
	fmt.Println(fib(2))	// 1
	fmt.Println(fib(3))	// 2
	fmt.Println(fib(4))	// 3
	fmt.Println(fib(5))	// 5
	fmt.Println(fib(6))	// 8
	fmt.Println(fib(7))	// 13
}