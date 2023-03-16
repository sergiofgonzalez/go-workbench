package main

import "example.com/pocketlog"


func main() {
	log := pocketlog.New(pocketlog.LevelInfo)

	log.Debugf("Hi! this is a debug message that will be suppressed!")
	log.Infof("Hi, this is an informational message!")
	log.Errorf("Hi, this is an error message!")
}