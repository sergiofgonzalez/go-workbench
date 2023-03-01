package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is the error that will be returned when
// Sqrt is invoked with a negative value such as -2
type ErrNegativeSqrt float64

// We make ErrNegativeSqrt an `error` by implementing
// the `error` interface which requires implementing the
// Error() string function
func (e ErrNegativeSqrt) Error() string {
	// e must be converted into a float first to prevent
	// an infinite recursive error
	// %v with an Error argument will invoke Error(), etc. etc.
	f := float64(e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", f)
}


// Sqrt returns an estimation of the square root of the given value
// using loops. When invoked with a negative value it returns an
// error as the second argument
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0 // option 1: use 1.0; option 2: use x; option 3: use x / 2
	prevZ := 0.0
	for i := 0; i < 10 && math.Abs((prevZ - z)) > 1.0e-25; i++ {
		prevZ = z
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}


func main() {
	// Positive path
	result, err := Sqrt(2)
	if err == nil {
		fmt.Printf("Sqrt(2)=%v\n", result)
	} else {
		fmt.Printf("An error occurred invoking Sqrt(2): %v\n", err)
	}

	// Error path
	result, err = Sqrt(-2)
	if err == nil {
		fmt.Printf("Sqrt(-2)=%v\n", result)
	} else {
		fmt.Printf("An error occurred invoking Sqrt(-2): %v\n", err)
	}
}