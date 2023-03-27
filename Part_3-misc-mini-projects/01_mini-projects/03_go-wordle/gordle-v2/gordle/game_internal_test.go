package gordle

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameAskHappyPathSimple(t *testing.T) {
	input := "Hello"
	var expected []rune = []rune("Hello")

	g := New(strings.NewReader(input))
	actual := g.ask()

	if !slices.Equal(actual, expected) {
		t.Errorf("ask() returned %v, expected %v", actual, expected)
	}
}

func TestGameAskHappyPath(t *testing.T) {
	tests := map[string]struct {
		input string
		want	[]rune
	}{
		"5 characters in English": {
			input: "Hello",
			want: []rune("Hello"),
		},
		"5 characters in Spanish": {
			input: "cañón",
			want: []rune("cañón"),
		},
		"5 characters in Chinese": {
			input: "你好世界!",
			want: []rune("你好世界!"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input))

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("%v: expected %q (%v), but got %q (%v)", name, string(tc.want), tc.want, string(got), got)
			}
		})
	}
}
