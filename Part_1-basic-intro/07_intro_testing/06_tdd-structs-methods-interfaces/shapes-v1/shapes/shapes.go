package shapes

// Rectangle is shape define by its width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of the rectangle given its width and height
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the perimeter of the rectangle given its width and height
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
