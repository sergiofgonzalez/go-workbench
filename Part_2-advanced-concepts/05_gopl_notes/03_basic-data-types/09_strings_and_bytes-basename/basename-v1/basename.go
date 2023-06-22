package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	// Trim everything until the last '/' if any
	s = s[strings.LastIndex(s, "/")+1:]

	// Trim everything after the last '.' if any
	if posDot := strings.LastIndex(s, "."); posDot >= 0 {
		s = s[:posDot]
	}

	return s
}

func main() {
	fmt.Println("Run the tests!")
}
