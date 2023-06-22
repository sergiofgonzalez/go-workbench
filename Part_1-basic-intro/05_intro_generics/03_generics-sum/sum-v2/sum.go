package main

import "fmt"

// Number declares the different types a Number can hold
type Number interface {
	int | int64 | float64
}

// Sum returns the sum of the numbers given
func Sum[T Number](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println("sum ints   : ", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println("sum int64  : ", Sum([]int64{1, 2, 3, 4, 5}))
	fmt.Println("sum float64: ", Sum([]float64{1.1, 2.2, 3.3, 4.4, 5.5}))
}