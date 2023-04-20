package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello to Jason Isaacs!
}

func TestHello(t *testing.T) {
	t.Run("no name given", func(t *testing.T) {
		want := "Hello, World!"
		got := Hello("")
		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
		assertCorrectMessage(t, got, want)
	})

	t.Run("'Jason Isaacs' name given", func(t *testing.T) {
		want := "Hello to Jason Isaacs!"
		got := Hello("Jason Isaacs")
		assertCorrectMessage(t, got, want)
	})

	t.Run("'Florence Pugh' name given", func(t *testing.T) {
		want := "Hello to Florence Pugh!"
		got := Hello("Florence Pugh")
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

}