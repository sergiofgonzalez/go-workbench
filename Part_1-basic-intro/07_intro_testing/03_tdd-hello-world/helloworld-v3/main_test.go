package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello to Jason Isaacs!
}

func TestHello(t *testing.T) {
	tests := map[string]struct {
		name string
		want string
	}{
		"no name given":             {name: "", want: "Hello, World!"},
		"'Jason Isaacs' name given": {name: "Jason Isaacs", want: "Hello to Jason Isaacs!"},
		"'Florence' name given":     {name: "Florence", want: "Hello to Florence!"},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := Hello(testCase.name)
			assertCorrectMessage(t, got, testCase.want)
		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
