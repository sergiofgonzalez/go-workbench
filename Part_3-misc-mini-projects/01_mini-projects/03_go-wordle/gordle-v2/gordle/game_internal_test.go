package gordle

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameAskHappyPath(t *testing.T) {
	input := "Hello"
	var expected []rune = []rune("Hello")

	g := New(strings.NewReader(input))
	actual := g.ask()

	if !slices.Equal(actual, expected) {
		t.Errorf("ask() returned %v, expected %v", actual, expected)
	}
}