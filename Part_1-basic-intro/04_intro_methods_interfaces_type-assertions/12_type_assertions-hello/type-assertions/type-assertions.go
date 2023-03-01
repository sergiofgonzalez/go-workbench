package main

import "fmt"

func main() {
	var i interface{} = "Hello to Jason Isaacs"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Printf("s=%v, ok=%v\n", s, ok)

	// It's OK, it's just a test type assertion
	f, ok := i.(float64)
	fmt.Printf("s=%v, ok=%v\n", f, ok)

	// Panic!
	f = i.(float64)
	fmt.Printf("s=%v, ok=%v\n", f, ok)
}