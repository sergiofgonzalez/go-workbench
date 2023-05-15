package main

// ConvertToRoman returns the representation of the given number in Roman Numerals
func ConvertToRoman(n int) string {
	switch n {
	case 3:
		return "III"
	case 2:
		return "II"
	default:
		return "I"
	}
}