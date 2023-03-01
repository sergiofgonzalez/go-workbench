package main

import "fmt"

func main() {
	primes := []int{ 2, 3, 5, 7, 11, 13 }

	// stage 0: use bounds
	s := primes[1:4] // -> 3, 5, 7
	fmt.Println(s)

	// stage 1: use lower-bound only
	s = s[:2] // -> 3, 5
	fmt.Println(s)

	// stage 2: use upper-bound only
	s = s[1:] // -> 5
	fmt.Println(s)
}