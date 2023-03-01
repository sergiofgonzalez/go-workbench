package main

import (
	"fmt"
)

func main() {
	powers := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range powers {
		fmt.Printf("2**%d = %v\n", i, v)
	}

	powers2 := make([]int, 10)
	for i := range powers2 {
		powers2[i] = 1 << uint(i)
	}

	for _, value := range powers2 {
		fmt.Printf("%d ", value)
	}

}