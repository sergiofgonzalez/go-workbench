package main

import (
	"fmt"
	"math"
)

// Sqrt returns an estimation of the square root of the given value
// using loops
func Sqrt(x float64) float64 {
	z := 1.0 // option 1: use 1.0; option 2: use x; option 3: use x / 2
	prevZ := 0.0
	for i := 0; i < 10 && math.Abs((prevZ - z)) > 1.0e-25; i++ {
		fmt.Printf("\titeration: %v, x: %g; z: %g\n", i, x, z)
		candidateZ2 := z * z
		fmt.Printf("\t\tabs(z^2 - x=%g)\n", math.Abs(candidateZ2 - x))
		prevZ = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}


func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sqrt(%v):\n", i)
		fmt.Printf("\n\t ==>: sqrt(%v)=%g (math.Sqrt(%v)=%g)\n\n", i, Sqrt(float64(i)), i, math.Sqrt(float64(i)))
	}
}

