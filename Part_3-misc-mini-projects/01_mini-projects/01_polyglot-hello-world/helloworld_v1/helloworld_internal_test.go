package main

import (
	"testing"
)

func Example_main() {
	main()
	// Output:
	// Hola, mundo!
}

func TestGreetSupportedLanguages(t *testing.T) {

	// Arrange
	expected := map[language]string{
		"es": "Hola, mundo!",
		"en": "Hello, world!",
		"fr": "Bonjour, le monde!",
		"de": "Hallo, Welt!",
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

func TestGreetUnsupportedLanguages(t *testing.T) {

	// Arrange
	unsupportedLang := "zz"
	expected := `unsupported language: "zz"`

	// Act
	actual := greet(language(unsupportedLang))

	// Assert
	if actual != expected {
		t.Errorf("expected %q for %v, but got %q", expected, unsupportedLang, actual)
	}
}