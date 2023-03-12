package main

import (
	"os"

	"example.com/pocketlog/pocketlog"
)

func main() {

	// Getting a logger with a certain threshold
	l := pocketlog.New(pocketlog.LevelInfo, os.Stdout)

	// Logf lets you choose the level
	// log lines below the threshold will be suppressed
	l.Logf(pocketlog.LevelDebug, "This will be suppressed")
	l.Logf(pocketlog.LevelInfo, "This will not be suppressed")
	l.Logf(pocketlog.LevelError, "This will not be suppressed either")

	// Alternatively you can use the convenience methods
	l.Debugf("This will be suppressed")
	l.Infof("This will not be suppressed")
	l.Errorf("This will not be suppressed either")
}
