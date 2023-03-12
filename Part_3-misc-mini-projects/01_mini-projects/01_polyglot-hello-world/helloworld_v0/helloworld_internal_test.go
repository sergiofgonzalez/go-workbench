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

	// Arrange
	expected := map[language]string{
		"es": "Hola, mundo!",
		"en": "Hello, world!",
		"fr": "Bonjour, le monde!",
		"de": "Hallo, Welt!",
		"zh": "",
	}

	// Act
	for k, v := range expected {
		actual := greet(k)

		// Assert
		if actual != v {
			t.Errorf("expect %q for %v, but got %q", v, k, actual)
		}
	}

}