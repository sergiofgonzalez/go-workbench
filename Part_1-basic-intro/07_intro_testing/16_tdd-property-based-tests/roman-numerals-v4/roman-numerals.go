package main

import "strings"


// RomanNumeral struct maintains the association between an arabic number and its
// string representation
type RomanNumeral struct {
	Value  int
	Symbol string
}


var allRomanNumerals = []RomanNumeral{
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ConvertToRoman returns the representation of the given number in Roman Numerals
func ConvertToRoman(n int) string {
	sb := strings.Builder{}

	for _, numeral := range allRomanNumerals {
		for n >= numeral.Value {
			sb.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return sb.String()
}
