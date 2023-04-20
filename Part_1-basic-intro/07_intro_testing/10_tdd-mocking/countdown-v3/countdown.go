package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3

// Sleeper interface let us apply DI to the Countdown
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is a construct that will be provided to CountDown in main
type DefaultSleeper struct{}

// Sleep is the default implementation for the Sleeper interface
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown manages to write 3, 2, 1, Go! with 1-second pause
func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprintln(w, "Go!")
}

func main() {
	d := DefaultSleeper{}
	Countdown(os.Stdout, &d)  // won't compile for now
}
