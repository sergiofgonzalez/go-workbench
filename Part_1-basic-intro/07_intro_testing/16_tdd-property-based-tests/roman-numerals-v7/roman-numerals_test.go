package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	arabic int
	roman  string
}

var tests = []TestCase{
	{
		arabic: 1,
		roman:  "I",
	},
	{
		arabic: 2,
		roman:  "II",
	},
	{
		arabic: 3,
		roman:  "III",
	},
	{
		arabic: 4,
		roman:  "IV",
	},
	{
		arabic: 5,
		roman:  "V",
	},
	{
		arabic: 6,
		roman:  "VI",
	},
	{
		arabic: 7,
		roman:  "VII",
	},
	{
		arabic: 8,
		roman:  "VIII",
	},
	{
		arabic: 9,
		roman:  "IX",
	},
	{
		arabic: 10,
		roman:  "X",
	},
	{
		arabic: 14,
		roman:  "XIV",
	},
	{
		arabic: 18,
		roman:  "XVIII",
	},
	{
		arabic: 20,
		roman:  "XX",
	},
	{
		arabic: 30,
		roman:  "XXX",
	},
	{
		arabic: 39,
		roman:  "XXXIX",
	},
	{
		arabic: 40,
		roman:  "XL",
	},
	{
		arabic: 47,
		roman:  "XLVII",
	},
	{
		arabic: 49,
		roman:  "XLIX",
	},
	{
		arabic: 50,
		roman:  "L",
	},
	{
		arabic: 1984,
		roman:  "MCMLXXXIV",
	},
	{
		arabic: 3999,
		roman:  "MMMCMXCIX",
	},
	{
		arabic: 2014,
		roman:  "MMXIV",
	},
	{
		arabic: 1974,
		roman:  "MCMLXXIV",
	},
	{
		arabic: 1006,
		roman:  "MVI",
	},
	{
		arabic: 798,
		roman:  "DCCXCVIII",
	},
}

func TestConvertToRoman(t *testing.T) {

	for _, testCase := range tests {
		testName := fmt.Sprintf("Converting %d to %s", testCase.arabic, testCase.roman)
		t.Run(testName, func(t *testing.T) {
			got := ConvertToRoman(testCase.arabic)
			if got != testCase.roman {
				t.Errorf("%s: got %q, but wanted %q", testName, got, testCase.roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {

	for _, testCase := range tests {
		testName := fmt.Sprintf("Converting %d to %s", testCase.arabic, testCase.roman)
		t.Run(testName, func(t *testing.T) {
			got := ConvertToArabic(testCase.roman)
			if got != testCase.arabic {
				t.Errorf("%s: got %d, but wanted %d", testName, got, testCase.arabic)
			}
		})
	}
}
