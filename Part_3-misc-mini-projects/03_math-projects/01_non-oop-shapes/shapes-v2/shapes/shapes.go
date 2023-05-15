package shapes

// Shape is an interface that describes the common behavior of a shape
type Shape interface {
	Area() float64
	Scale(factor float64) Shape
	Equal(other Shape) bool
}

// Area returns the area of the given shape
func Area(s Shape) float64 {
	return s.Area()
}

// Equal returns true if the shapes are of the same type and have the same dimensions
func Equal(s1 Shape, s2 Shape) bool {
	return s1.Equal(s2)
}

// Scale returns a shape of the same type with its dimensions scaled by the given factor
func Scale(s Shape, factor float64) Shape {
	return s.Scale(factor)
}