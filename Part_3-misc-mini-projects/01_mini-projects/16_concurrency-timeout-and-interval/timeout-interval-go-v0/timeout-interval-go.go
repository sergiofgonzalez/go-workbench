package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func main() {
	setTimeout(printHello, 50)
	fmt.Println("World!")


	setInterval(func() {
		fmt.Println("Hello to Jason Isaacs!")
	}, 1000)

	time.Sleep(12 * time.Second) // program is allowed to exit if the go routine has not completed
}

func setTimeout(fn func(), delay int) {
	go func() {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fn()
	}()
}

func setInterval(fn func(), delay int) {
	go func() {
		for {
			time.Sleep(time.Duration(delay) * time.Millisecond)
			fn()
		}
	}()
}