package main

import (
	"testing"
)

func Example_main() {
	main()
	// Output:
	// Hola, mundo!
}

func TestGreet(t *testing.T) {
	type TestCase struct {
		lang             language
		expectedGreeting string
	}

	tests := map[string]TestCase{
		"Spanish": {
			"es",
			"Hola, mundo!",
		},
		"English": {
			"en",
			"Hello, world!",
		},
		"French": {
			"fr",
			"Bonjour, le monde!",
		},
		"German": {
			"de",
			"Hallo, Welt!",
		},
		"Empty": {
			"",
			`unsupported language: ""`,
		},
		"Unsupported": {
			"xyz",
			`unsupported language: "xyz"`,
		},
	}

	for subtestLabel, subtestCase := range tests {
		t.Run(subtestLabel, func(t *testing.T) {
			expected := subtestCase.expectedGreeting
			actual := greet(subtestCase.lang)

			if actual != expected {
				t.Errorf("expected %q for %v, but found %q", expected, subtestCase.lang, actual)
			}
		})
	}
}
