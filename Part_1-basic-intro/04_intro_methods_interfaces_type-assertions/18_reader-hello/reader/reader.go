package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello to Jason Isaacs!")

	// Define the slice of len 8
	b := make([]byte, 8)

	// while loop :p
	for {
		n, err := r.Read(b)
		fmt.Printf("\tn=%v, err=%v, b=%v\n", n, err, b)
		fmt.Printf("b[:n]: %q\n", b[:n])

		// Have we exhausted the stream?
		if err == io.EOF {
			break
		}
	}
}