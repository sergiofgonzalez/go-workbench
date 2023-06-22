package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "世界"
	fmt.Println(s)                         // prints Chinese characters
	fmt.Println(len(s))                    // 6 bytes
	fmt.Println(utf8.RuneCountInString(s)) // 2 chars

	// To process the characters individually we need to decode them
	fmt.Println()
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// We can also use a simple range s to get the runes
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%c\n", r)
	}

	// Another range loop decoding a string with Western and Eastern runes
	fmt.Println()
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%[2]d\t0x%[2]x\n", i, r)
	}

	// Counting the runes in a string
	n := 0
	s = "Hello, 世界"
	for range s {
		n++
	}
	fmt.Printf("runes: %d, bytes: %d\n", n, len(s))
}
