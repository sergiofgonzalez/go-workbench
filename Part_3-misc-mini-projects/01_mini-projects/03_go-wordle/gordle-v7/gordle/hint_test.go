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

func TestEqual(t *testing.T) {
	tests := map[string]struct {
		fb1  feedback
		fb2  feedback
		want bool
	}{
		"zero values": {
			want: true,
		},
		"zero value and non-zero value": {
			fb2: feedback{absentCharacter},
			want: false,
		},
		"non-zero value and zero value": {
			fb1: feedback{absentCharacter},
			want: false,
		},
		"different len feedback": {
			fb1: feedback{absentCharacter, wrongPosition, correctPosition},
			fb2: feedback{absentCharacter, wrongPosition},
			want: false,
		},
		"same len different feedback": {
			fb1: feedback{absentCharacter, wrongPosition},
			fb2: feedback{absentCharacter, correctPosition},
			want: false,
		},
		"len one same fb": {
			fb1: feedback{absentCharacter},
			fb2: feedback{absentCharacter},
			want: true,
		},
		"len greater than one same fb": {
			fb1: feedback{absentCharacter, correctPosition, correctPosition, wrongPosition, wrongPosition},
			fb2: feedback{absentCharacter, correctPosition, correctPosition, wrongPosition, wrongPosition},
			want: true,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			got := tc.fb1.Equal(tc.fb2)
			if got != tc.want {
				t.Errorf("%v: wanted %v, but got %v", testName, tc.want, got)
			}
		})
	}
}
