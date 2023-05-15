package main

import "strings"

// ConvertToRoman returns the representation of the given number in Roman Numerals
func ConvertToRoman(n int) string {
	sb := strings.Builder{}

	for n > 0 {
		switch {
		case n > 4:
			sb.WriteString("V")
			n -= 5
		case n > 3:
			sb.WriteString("IV")
			n -= 4
		default:
			sb.WriteString("I")
			n--
		}
	}

	return sb.String()
}