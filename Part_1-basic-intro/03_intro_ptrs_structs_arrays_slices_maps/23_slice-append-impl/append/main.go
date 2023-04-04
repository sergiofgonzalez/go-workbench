package main

import "fmt"

func main() {
	s := make([]byte, 5, 5)
	fmt.Printf("0: Initial   : %v; len=%v, cap=%v\n", s, len(s), cap(s))

	// setting some values
	for i := 0; i < len(s); i++ {
		s[i] = byte(i)
	}
	fmt.Printf("1: Values Set: %v; len=%v, cap=%v\n", s, len(s), cap(s))

	s = Append(s, []byte{1})
	fmt.Printf("1: After     : %v; len=%v, cap=%v\n", s, len(s), cap(s))
}

// Append is a custom implementation of `append` used for learning purposes
func Append(s []byte, data []byte) []byte {
	l := len(s) // Initial length

	if l+len(data) > cap(s) {
		aux := make([]byte, 2*(l+len(data)))
		copy(aux, s)
		s = aux
	}
	s = s[0 : l+len(data)] // Enlarge the slice
	copy(s[l:], data)      // Copy the data into the slice

	return s
}

// test with zero values
// test with one value beyond cap
// test with multiple values beyond cap
