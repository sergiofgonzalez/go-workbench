package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown manages to write 3, 2, 1, Go! with 1-second pause
func Countdown(w io.Writer) {
	fmt.Fprintln(w, "3")
	fmt.Fprintln(w, "2")
	fmt.Fprintln(w, "1")
	fmt.Fprintln(w, "Go!")
}

func main() {
	Countdown(os.Stdout)
}