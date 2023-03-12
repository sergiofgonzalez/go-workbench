package main

import (
	"example.com/pocketlog/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	// Logf samples
	l.Logf(pocketlog.LevelInfo, "This is a %s message", "info")
	l.Logf(pocketlog.LevelDebug, "This is a %s message", "debug")
	l.Logf(pocketlog.LevelError, "This is a %s message", "error")

	// Debugf sample
	l.Debugf("This is another %s message", "debug")

	// Infof sample
	l.Infof("This is another %s message", "info")

	// Infof sample
	l.Errorf("This is another %s message", "error")
}
