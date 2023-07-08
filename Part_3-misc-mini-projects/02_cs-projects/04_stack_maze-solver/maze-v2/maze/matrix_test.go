package maze_test

import (
	"testing"

	"example.com/maze/maze"
	"golang.org/x/exp/slices"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		maze     []string
		expected maze.Matrix
	}{
		"empty2x2": {
			maze: []string{
				"  ",
				"  "},
			expected: [][]string{
				{" ", " "},
				{" ", " "},
			},
		},
		"nonEmpty3x3": {
			maze: []string{
				"A X",
				"X B"},
			expected: [][]string{
				{"A", " ", "X"},
				{"X", " ", "B"},
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := maze.New(testCase.maze)
			assertEqual(*got, testCase.expected, t)
		})
	}
}

func assertEqual(mGot, mExpected maze.Matrix, t testing.TB) {
	t.Helper()
	if len(mGot) != len(mExpected) {
		t.Errorf("different lengths: m1 is %d and m2 is %d", len(mGot), len(mExpected))
	}

	for i := range mGot {
		if !slices.Equal(mGot[i], mExpected[i]) {
			t.Errorf("row %d: expected %q, but got %q", i, mExpected[i], mGot[i])
		}
	}
}