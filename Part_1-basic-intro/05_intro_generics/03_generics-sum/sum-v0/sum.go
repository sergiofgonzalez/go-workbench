package main

import "fmt"

// SumInts returns the sum of the numbers given
func SumInts(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

// SumInt64s returns the sum of the numbers given
func SumInt64s(nums []int64) int64 {
	var sum int64
	for _, n := range nums {
		sum += n
	}
	return sum
}


// SumFloat64s returns the sum of the numbers given
func SumFloat64s(nums []float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum
}


func main() {
	fmt.Println("sum ints   : ", SumInts([]int{1, 2, 3, 4, 5}))
	fmt.Println("sum int64  : ", SumInt64s([]int64{1, 2, 3, 4, 5}))
	fmt.Println("sum float64: ", SumFloat64s([]float64{1.1, 2.2, 3.3, 4.4, 5.5}))
}
