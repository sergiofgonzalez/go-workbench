package main

import "math"
import "fmt"

func main() {

	// %g uses either %f or %e depending on the exponent
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
}
