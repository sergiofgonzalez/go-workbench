package tempconv

import "fmt"

// Celsius temperature type declaration
type Celsius float64

// Fahrenheit temperature type declaration
type Fahrenheit float64

const (
	// AbsoluteZeroC is the absolute zero temperature in Celsius
	AbsoluteZeroC Celsius = -273.15

	// FreezingC is the absolute zero temperature in Celsius
	FreezingC Celsius = 0

	// BoilingC is the absolute zero temperature in Celsius
	BoilingC Celsius= 100
)

// CtoF converts a Celsius temperature into its Fahrenheit equivalent
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FtoC converts a Fahrenheit temperature into its Celsius equivalent
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}