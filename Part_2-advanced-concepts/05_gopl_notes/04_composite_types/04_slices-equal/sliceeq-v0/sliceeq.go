package main

import "fmt"

func equal[T comparable](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	x := []int{1, 2, 3}
	y := []int{1, 2}
	z := []int{1, 2, 3}

	fmt.Println(equal(x, y))
	fmt.Println(equal(x, z))
	fmt.Println(equal(y, z))
}