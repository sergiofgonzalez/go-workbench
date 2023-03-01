package main

import "fmt"

func main() {
	result := 0
	for i := 0; i < 10; i++ {
		result += i
	}

	fmt.Printf("Result=%v\n", result)
}