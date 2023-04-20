package shapes

import "math"

// Rectangle is a shape defined by its width and height dimensions
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is a shape defined by its radius
type Circle struct {
	Radius float64
}

// Triangle is a shape defined by the dimensions of his base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Shape is an interface that describes the available methods on Shapes
type Shape interface {
	Perimeter() float64
	Area() float64
}

// Perimeter returns the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the perimeter of the rectangle
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of the circle
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area returns the perimeter of the circle
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of the circle
func (t *Triangle) Perimeter() float64 {
	panic("Can't compute")
}

// Area returns the perimeter of the circle
func (t *Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
