package gordle

import "testing"

func TestFeedbackString(t *testing.T) {
	tests := map[string]struct {
		fb   feedback
		want string
	}{
		"unknown position": {
			fb:   feedback{69},
			want: "💔",
		},
		"one of each in order": {
			fb:   feedback{absentCharacter, wrongPosition, correctPosition},
			want: "⬜️🟡💚",
		},
		"one of each random": {
			fb:   feedback{wrongPosition, correctPosition, absentCharacter},
			want: "🟡💚⬜️",
		},
		"all correct pos": {
			fb:   feedback{correctPosition, correctPosition, correctPosition, correctPosition},
			want: "💚💚💚💚",
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			got := tc.fb.String()
			if got != tc.want {
				t.Errorf("%v: wanted %q, but got %q", testName, tc.want, got)
			}
		})
	}
}
