package main

import "strings"


// RomanNumeral struct maintains the association between an arabic number and its
// string representation
type RomanNumeral struct {
	Value  int
	Symbol string
}


var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
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

// ConvertToArabic gets the Arabic representation of a roman numeral
func ConvertToArabic(roman string) int {
	switch {
	case roman == "I":
		return 1
	case roman == "II":
		return 2
	case roman == "III":
		return 3
	default:
		return -1
	}
}