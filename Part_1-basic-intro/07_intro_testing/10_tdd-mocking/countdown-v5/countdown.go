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

// ConfigurableTimeSleeper is a construct that will be provided to Countdown
// it implements the Sleeper interface
type ConfigurableTimeSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep is the Sleeper interface implementation for the ConfigurableTimeSleeper
func (c *ConfigurableTimeSleeper) Sleep() {
	c.sleep(c.duration)
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
	sleeper := ConfigurableTimeSleeper{
		duration: 1 * time.Second,
		sleep: time.Sleep,
	}
	Countdown(os.Stdout, &sleeper) // won't compile for now
}
