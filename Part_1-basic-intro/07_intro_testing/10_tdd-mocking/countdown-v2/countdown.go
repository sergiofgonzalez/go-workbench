package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3

// Countdown manages to write 3, 2, 1, Go! with 1-second pause
func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintln(w, "Go!")
}

func main() {
	Countdown(os.Stdout)
}