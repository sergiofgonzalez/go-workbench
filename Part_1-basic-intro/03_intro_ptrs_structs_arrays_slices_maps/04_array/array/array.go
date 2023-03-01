package main

import "fmt"

func main() {
	var words[3]string

	words[0] = "Hello"
	words[1] = "to"
	words[2] = "Jason Isaacs"

	fmt.Println(words[0], words[2])

	primes := [7]int{1, 2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}