package main

import "fmt"

func fibonacci(ch chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
				x, y = y, y + x
		case <-quit:
				fmt.Println("quit signal received")
				return
		}
	}
}


func main() {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		// will wait for info available on ch for 10 times, then it'll signal it to quit
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)		// this will block until fibonacci has sent info to the ch channel
		}
		quit <- 0	// If you comment this line you'll get a deadlock panic
	}()
	fibonacci(ch, quit)
}