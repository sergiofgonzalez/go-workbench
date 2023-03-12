package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello to Jason Isaacs!
}

func TestGreet(t *testing.T) {
	// Arrange
	expected := "Hello to Jason Isaacs!"


	// Act
	actual := greet()

	// Assert
	if actual != expected {
		t.Errorf("expected %q, but found %q", expected, actual)
	}

	// Teardown: nothing to do
}