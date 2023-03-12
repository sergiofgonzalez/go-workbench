package main

import (
	"fmt"

	"example.com/pocketlog/pocketlog"
)

func main() {

	// Getting a logger with a certain threshold
	l := pocketlog.New(pocketlog.LevelInfo)

	// Logf lets you choose the level
	// log lines below the threshold will be suppressed
	l.Logf(pocketlog.LevelDebug, "This will be suppressed")
	l.Logf(pocketlog.LevelInfo, "This will not be suppressed")
	l.Logf(pocketlog.LevelError, "This will not be suppressed either")

	// Alternatively you can use the convenience methods
	l.Debugf("This will be suppressed")
	l.Infof("This will not be suppressed")
	l.Errorf("This will not be suppressed either")

	// You can also configure it to show the timestamp
	l = pocketlog.New(pocketlog.LevelInfo, pocketlog.WithTimestamp())
	l.Infof("Now with time!")

	// Or to log JSON messages
	l = pocketlog.New(pocketlog.LevelInfo, pocketlog.WithJson())
	l.Infof("Hello, %s", "Jason")


	helloVariadic("one", "Jason", "Isaacs")
}


func helloVariadic(s string, other ...any) {
	for i, v := range other {
		fmt.Printf("%v: %q\n", i, v)
	}
}