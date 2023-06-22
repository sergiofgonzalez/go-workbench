package main

import (
	"fmt"

	"example.com/tempconv/tempconv"
)

func main() {
	fmt.Printf("Absolute zero in °F = %g\n", tempconv.CtoF(tempconv.AbsoluteZeroC))
	fmt.Printf("Water boiling point in °F = %g\n", tempconv.CtoF(tempconv.BoilingC))
	fmt.Printf("Water freezing point in °F = %g\n", tempconv.CtoF(tempconv.FreezingC))
	fmt.Printf("100 °F in °C = %g\n", tempconv.FtoC(100))

	boilingF := tempconv.CtoF(tempconv.BoilingC)
	fmt.Printf("Boiling - Freezing point in °F = %g\n", boilingF - tempconv.CtoF(tempconv.FreezingC))

	// because we defined the named types the following fails
	// fmt.Println(boilingF - tempconv.FreezingC) // ERR: mismatched types

	// Similarly with comparisons
	// fmt.Println(boilingF == tempconv.FreezingC) // ERR: mismatched types

	// But you can compare with the underlying types
	fmt.Println(tempconv.BoilingC == 100)
	fmt.Println(boilingF > 100)

	fmt.Println(boilingF)
}