package iteration

import "strings"

// Repeat returns the result of concatenating the given string 5 times
func Repeat(s string) string {
	result := ""
	for i := 0; i < 5; i++ {
		result += s
	}
	return result
}

// RepeatWithSb is an alternative implementation of Repeat using strings.Builder
func RepeatWithSb(s string) string {
	sb := strings.Builder{}
	for i := 0; i < 5; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

// RepeatNative is an alternative implementation of Repeat using strings.Repeat
func RepeatNative(s string) string {
	return strings.Repeat(s, 5)
}
