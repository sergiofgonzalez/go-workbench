package main

import (
	"fmt"

	"example.com/pocketlog/pocketlog"
)

func main() {

	/*
		Informal tests 1: not using New()
	*/
	l := pocketlog.Logger{} // default threshold is Debug

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

	fmt.Println("=============\n")

	/*
		Informal tests 2: Using a zero value Logger
	*/
	var l2 pocketlog.Logger

	// Logf samples
	l2.Logf(pocketlog.LevelInfo, "This is a %s message", "info")
	l2.Logf(pocketlog.LevelDebug, "This is a %s message", "debug")
	l2.Logf(pocketlog.LevelError, "This is a %s message", "error")

	// Debugf sample
	l2.Debugf("This is another %s message", "debug")

	// Infof sample
	l2.Infof("This is another %s message", "info")

	// Infof sample
	l2.Errorf("This is another %s message", "error")

	fmt.Println("=============\n")

	/*
		Informal tests 3: Using New()
	*/
	l3 := pocketlog.New(pocketlog.LevelDebug)

	// Logf samples
	l3.Logf(pocketlog.LevelInfo, "This is a %s message", "info")
	l3.Logf(pocketlog.LevelDebug, "This is a %s message", "debug")
	l3.Logf(pocketlog.LevelError, "This is a %s message", "error")

	// Debugf sample
	l3.Debugf("This is another %s message", "debug")

	// Infof sample
	l3.Infof("This is another %s message", "info")

	// Infof sample
	l3.Errorf("This is another %s message", "error")

	fmt.Println()
	l3 = pocketlog.New(pocketlog.LevelInfo)
	l3.Logf(pocketlog.LevelDebug, "This will be suppressed")
	l3.Logf(pocketlog.LevelInfo, "This will not be suppressed")
	l3.Logf(pocketlog.LevelError, "This will not be suppressed either")

	l3.Debugf("This will be suppressed")
	l3.Infof("This will not be suppressed")
	l3.Errorf("This will not be suppressed either")

	fmt.Println()
	l3 = pocketlog.New(pocketlog.LevelError)
	l3.Logf(pocketlog.LevelDebug, "This will be suppressed")
	l3.Logf(pocketlog.LevelInfo, "This will be suppressed")
	l3.Logf(pocketlog.LevelError, "This will not be suppressed")

	l3.Debugf("This will be suppressed")
	l3.Infof("This will be suppressed")
	l3.Errorf("This will not be suppressed")
}
