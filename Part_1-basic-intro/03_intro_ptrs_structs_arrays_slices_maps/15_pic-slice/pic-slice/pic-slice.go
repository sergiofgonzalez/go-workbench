package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)


// Pic1 returns a slice of length dx * dy, where each element
// is an 8-bit unsigned int
func Pic1(dx, dy int) [][]uint8 {
	var matrix = make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		var rowCells = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			rowCells[x] = uint8(x * y)
		}
		matrix[y] = rowCells
	}
	return matrix
}

// Pic2 returns a slice of length dx * dy, where each element
// is an 8-bit unsigned int (alternative implementation)
func Pic2(dx, dy int) [][]uint8 {
	var matrix [][]uint8

	for y := 0; y < dy; y++ {
		rowCells := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			rowCells[x] = uint8(x * y)
		}
		matrix = append(matrix, rowCells)
	}
	return matrix
}
func main() {
	fmt.Println(Pic1(10, 5))
	pic.Show(Pic1)

	fmt.Println(Pic2(10, 5))
	pic.Show(Pic2)
}