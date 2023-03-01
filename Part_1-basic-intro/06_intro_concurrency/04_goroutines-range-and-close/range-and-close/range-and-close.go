package main

import "fmt"

// F0 = 0
// F1 = 1
// Fn = Fn-1 + Fn-2, for n > 2


// func fibonacci(n int) {
// 	x := 0
// 	y := 1
// 	for i := 0; i < n; i++ {
// 		if n == 1 {
// 			fmt.Println(0)
// 		} else if n == 2 {
// 			fmt.Println(1)
// 		} else {
// 			fmt.Println(y + x)
// 			x, y = y, y + x
// 		}
// 	}
// }

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, y + x

	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for v := range ch {
		fmt.Println(v)
	}
}