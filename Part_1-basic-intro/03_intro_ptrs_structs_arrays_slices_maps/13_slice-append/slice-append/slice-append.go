package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// You can append elements to nil slices
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4, 5)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, isNil=%v, %v\n", len(s), cap(s), s == nil, s)
}