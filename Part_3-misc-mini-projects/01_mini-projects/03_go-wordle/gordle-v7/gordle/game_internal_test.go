package gordle

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameAskHappyPathSimple(t *testing.T) {
	input := "Hello"
	var expected []rune = []rune("HELLO")

	g := New(strings.NewReader(input), "Hello", 3)
	actual := g.ask()

	if !slices.Equal(actual, expected) {
		t.Errorf("ask() returned %v, expected %v", actual, expected)
	}
}

func TestGameAskHappyPath(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in English": {
			input: "Hello",
			want:  []rune("HELLO"),
		},
		"5 characters in Spanish": {
			input: "cañón",
			want:  []rune("CAÑÓN"),
		},
		"5 characters in Chinese": {
			input: "你好世界!",
			want:  []rune("你好世界!"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), string(tc.input), 3)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("%v: expected %q (%v), but got %q (%v)", name, string(tc.want), tc.want, string(got), got)
			}
		})
	}
}

func TestValidateGuess(t *testing.T) {
	tests := map[string]struct {
		input    []rune
		expected error
	}{
		"Correct length": {
			input:    []rune("Hello"),
			expected: nil,
		},
		"Shorter length": {
			input:    []rune("Tree"),
			expected: errInvalidLengthWord,
		},
		"Longer length": {
			input:    []rune("Dracula"),
			expected: errInvalidLengthWord,
		},
		"Empty word": {
			input:    []rune(""),
			expected: errInvalidLengthWord,
		},
		"Nil slice": {
			input:    nil,
			expected: errInvalidLengthWord,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			g := New(nil, "Hello", 3) // We need g only to call validateGuess

			got := g.validateGuess([]rune(test.input))
			if !errors.Is(got, test.expected) {
				t.Errorf("%v: got error %v, but expected %v", name, got, test.expected)
			}
		})
	}
}

func TestSplitToUppercaseCharacters(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []rune
	}{
		"all lowercase": {
			input: "hello!",
			want:  []rune("HELLO!"),
		},
		"all uppercase": {
			input: "HELLO!",
			want:  []rune("HELLO!"),
		},
		"mixed case": {
			input: "Hello!",
			want:  []rune("HELLO!"),
		},
		"mixed Spanish": {
			input: "Cañón",
			want:  []rune("CAÑÓN"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			splitToUppercaseCharacters(tc.input)

		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tests := map[string]struct {
		solution []rune
		guess    []rune
		want     feedback
	}{
		"zero value solution and guess": {
		},
		"zero value solution non zero guess": {
			guess: []rune("Hello"),
			want: feedback{absentCharacter, absentCharacter, absentCharacter, absentCharacter, absentCharacter},
		},
		"zero value guess and non zero solution": {
			solution: []rune("Hello"),
		},
		"different len": {
			solution: []rune("Hello"),
			guess: []rune("Hell"),
			want: feedback{absentCharacter, absentCharacter, absentCharacter, absentCharacter},
		},
		"same len no hits": {
			solution: []rune("Hello"),
			guess: []rune("Whizz"),
			want: feedback{absentCharacter, absentCharacter, absentCharacter, absentCharacter, absentCharacter},
		},
		"same len a few hits": {
			solution: []rune("Hello"),
			guess: []rune("Plays"),
			want: feedback{absentCharacter, wrongPosition, absentCharacter, absentCharacter, absentCharacter},
		},
		"same len several hits": {
			solution: []rune("Hello"),
			guess: []rune("Fella"),
			want: feedback{absentCharacter, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"same word": {
			solution: []rune("Hello"),
			guess: []rune("Hello"),
			want: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := computeFeedback(testCase.guess, testCase.solution)
			if !got.Equal(testCase.want) {
				t.Errorf("%v: got %v, but wanted %v", testName, got, testCase.want)
			}
		})
	}
}
