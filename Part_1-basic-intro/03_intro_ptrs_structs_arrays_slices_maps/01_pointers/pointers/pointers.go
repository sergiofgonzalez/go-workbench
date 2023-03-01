package main

import "fmt"

func main() {
	var i, j = 42, 2701

	// 1
	p := &i
	fmt.Printf("*p=%v; i=%v\n", *p, i)

	// 2
	*p = 21
	fmt.Printf("i=%v\n", i)

	// 3
	p = &j
	*p /= 37
	fmt.Printf("j=%v\n", j)
}