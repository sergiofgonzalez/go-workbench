package main

import "fmt"

// reverse reverses the elements of a slice of ints in place without creating a copy
func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}


func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}  // array
	reverse(a[:])	// reverse the whole array
	fmt.Println(a)

	// Rotate a slice left by 2 positions
	s := []int{0, 1, 2, 3, 4, 5} // I want [2, 3, 4, 5, 0, 1]
	reverse(s[:2])	// [1, 0, 2, 3, 4, 5]
	reverse(s[2:])	// [1, 0, 5, 4, 3, 2]
	reverse(s)			// [2, 3, 4, 5, 0, 1]
	fmt.Println(s)

	// Rotate a slice right by 2 positions
	s = []int{0, 1, 2, 3, 4, 5} // I want [4, 5, 0, 1, 2, 3]
	reverse(s)			// [5, 4, 3, 2, 1, 0]
	reverse(s[:2])	// [4, 5, 3, 2, 1, 0]
	reverse(s[2:])	// [4, 5, 0, 1, 2, 3]
	fmt.Println(s)



}