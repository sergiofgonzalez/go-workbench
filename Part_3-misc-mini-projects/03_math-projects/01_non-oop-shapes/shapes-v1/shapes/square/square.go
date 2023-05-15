package square

import (
	"fmt"

	"example.com/shapes/shapes/rectangle"
)

// Square is a struct that models a Square given the size of its side
type Square struct {
	rect *rectangle.Rectangle
	Side float64
}

// New returns a pointer to an initialized rectangle given its width and height.
func New(side float64) *Square {
	square := Square{rectangle.New(side, side), side}
	return &square
}

// Area returns the surface area of a square
func (r *Square) Area() float64 {
	return r.rect.Area()
}

// Equal returns true if the given squares have the same dimensions
func (r *Square) Equal(other *Square) bool {
	return r.rect.Equal(other.rect)
}

// Scale returns a new square with its dimensions scaled by the given factor
func (r *Square) Scale(factor float64) *Square {
	return New(r.rect.Width * factor)
}

func (r *Square) String() string {
	return fmt.Sprintf("Square (%.2f)", r.Side)
}
