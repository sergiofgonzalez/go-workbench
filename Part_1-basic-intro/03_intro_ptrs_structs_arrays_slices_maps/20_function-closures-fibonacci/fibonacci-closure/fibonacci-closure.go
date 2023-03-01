package main

import "fmt"

func getFibonacci() func() int {
	iter := 0
	f0 := 0
	f1 := 1
	return func() int {
		switch iter {
		case 0:
			iter++
			return f0
		case 1:
			iter++
			return f1
		default:
			curr := f0 + f1
			f0, f1 = f1, curr
			return curr
		}
	}
}


func main() {
	fib := getFibonacci()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %d\n", i, fib())
	}
}