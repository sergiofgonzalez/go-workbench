package main

import (
	"fmt"
	"io"
	"os"
)

// Greet writes a personalized message in the given Writer
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello to %s!", name)
}

func main() {
	Greet(os.Stdout, "Jason Isaacs")
}
