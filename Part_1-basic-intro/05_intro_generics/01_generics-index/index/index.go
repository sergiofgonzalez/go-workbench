package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))
	fmt.Println(Index(si, 8))

	ss := []string{"Hello", "to", "Jason", "Isaacs"}
	fmt.Println(Index(ss, "Jason"))
	fmt.Println(Index(ss, "Margot"))
}