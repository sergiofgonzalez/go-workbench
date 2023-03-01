package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image will implement the Image interface
type Image struct {
	matrix [][]uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i.matrix[0]), len(i.matrix))
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(X, Y int) color.Color {
	return color.RGBA{i.matrix[Y][X], i.matrix[Y][X], 255, 255}
}

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
	m := Image{Pic1(10, 5)}
	pic.ShowImage(m)

	// fmt.Println(Pic1(10, 5))
	// pic.Show(Pic1)

	// fmt.Println(Pic2(10, 5))
	// pic.Show(Pic2)
}