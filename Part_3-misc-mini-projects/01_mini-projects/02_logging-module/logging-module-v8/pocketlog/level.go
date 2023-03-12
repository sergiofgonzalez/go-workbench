package pocketlog

// Level identifies the logging level as an integer number
type Level byte

const (
	// LevelDebug is the logging level for debugging log lines
	LevelDebug Level = iota

	// LevelInfo is the logging level for informational log lines
	LevelInfo

	// LevelError is the logging level for error information in the logs
	LevelError
)