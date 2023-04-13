package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func main() {
	c := make(chan int)

	setTimeout(printHello, 50, c)
	fmt.Println("World!")
	<- c // receivers block until there is data to receive

	// In setInterval, we can use the channel to cancel the timer
	setInterval(func() {
	 	fmt.Println("Hello to Jason Isaacs!")
	}, 1000, c)
	time.Sleep(10 * time.Second)
	c <- 1 // send a signal to cancel the timer
}

func setTimeout(fn func(), delay int, c chan int) {
	go func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fn()
		c <- 1
	}()
}

func setInterval(fn func(), delay int, c chan int) {
	go func() {
		for {
			time.Sleep(time.Duration(delay) * time.Millisecond)
			fn()
			select {
			case <-c:
				break
			default:
				continue
			}
		}
	}()
}