package rectangle

import (
	"fmt"

	"example.com/shapes/shapes"
)

// Rectangle is a struct that models a Rectangle given its width and height
type Rectangle struct {
	Width float64
	Height float64
}

// New returns a pointer to an initialized rectangle given its width and height.
func New(width float64, height float64) *Rectangle {
	return &Rectangle{width, height}
}

// Area returns the surface area of a rectangle
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Equal returns true if the given rectangles have the same dimensions
func (r *Rectangle) Equal(other shapes.Shape) bool {
	rOther, ok := other.(*Rectangle)
	if !ok {
		return false
	}

	return r.Width == rOther.Width && r.Height == rOther.Height
}


// Scale returns a new rectangle with its dimensions scaled by the given factor
func (r *Rectangle) Scale(factor float64) shapes.Shape {
	return New(r.Width * factor, r.Height * factor)
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle (%.2f by %.2f)", r.Width, r.Height)
}

// Square function returns a rectangle whose sides are equal to the given value
func Square(side float64) *Rectangle {
	return New(side, side)
}