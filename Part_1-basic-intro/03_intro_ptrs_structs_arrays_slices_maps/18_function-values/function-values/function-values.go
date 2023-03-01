package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64, num1, num2 float64) float64 {
	return fn(num1, num2)
}

func main() {
	hypotenuse := func(num1, num2 float64) float64 {
		return math.Sqrt(num1 * num1 + num2 * num2)
	}
	fmt.Println(compute(hypotenuse, 3, 4))
	fmt.Println(compute(math.Pow, 3, 4))

	// Note that we can use the function value as a regular function too
	fmt.Println(hypotenuse(5, 12))
}