package main

import "fmt"

func basename(s string) string {
	// The strategy is:
	// - starting from the end, discard everything until the last '/'
	//   and copy it into s
	// - starting from the end, discard everything until the first '.'

	// Trim everything until the last '/' if any
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Trim everything after the last '.' if any
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	fmt.Println("Run the tests!")
}