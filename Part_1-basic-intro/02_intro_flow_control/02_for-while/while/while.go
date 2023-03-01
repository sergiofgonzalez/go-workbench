package main

import "fmt"

func main()  {
	f0 := 0
	f1 := 1
	cur := f0 + f1
	for cur <= 21 {
		fmt.Println(cur)
		f0 = f1
		f1 = cur
		cur = f0 + f1
	}
}


// Fib
// f(0) = 0
// f(1) = 1
// f(2) = f(1) + f(0) = 1
// f(3) = f(2) + f(1) = 2
// f(4) = f(3) + f(2) = 3
// f(5) = f(4) + f(3) = 5