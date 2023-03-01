package main

import "fmt"

// add returns the integer resulting from adding the given values
func add(x int, y int) int {
	return x + y
}

func main() {
	result := add(2, 3)

	fmt.Println("The result of adding 2 and 3 is", result)
}