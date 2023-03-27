package main

import "fmt"

func main() {
	tests := map[string]struct {
		input string
		want  string
	}{
		"test1": {
			input: "input for test1",
			want: "expected result for test1",
		},
		"test2": {
			input: "input for test2",
			want: "expected result for test2",
		},
		"test3": {
			input: "input for test3",
			want: "expected result for test3",
		},
	}

	for testName, testCase := range tests {
		fmt.Printf("%v:\n\tinput: %q\n\twant: %q\n====\n", testName, testCase.input, testCase.want)
	}
}
