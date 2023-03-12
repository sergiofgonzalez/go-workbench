package pocketlog

import (
	"fmt"
)

// Logger allows an application to log info with different levels
type Logger struct {
	threshold Level
}

// New returns a Logger variable initialized with the given threshold
// for the logging level.
func New(threshold Level) *Logger {
	return &Logger{
		threshold,
	}
}

// Logf emits a log line with the corresponding level with the
// given format string and arguments
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if lvl < l.threshold {
		return
	}

	var strLogLevel string
	switch lvl {
	case LevelDebug:
		strLogLevel = "[DEBUG] "
	case LevelInfo:
		strLogLevel = "[INFO]  "
	case LevelError:
		strLogLevel = "[ERROR] "
	}

	loglineFormat := strLogLevel + format + "\n"
	fmt.Printf(loglineFormat, args...)
}

// Debugf emits a debug level log line with the given format string and arguments
func (l *Logger) Debugf(format string, args ...any) {
	// We could delegate log suppression to Logf() but
	// this simple condition saves a func call
	if l.threshold > LevelDebug {
		return
	}
	l.Logf(LevelDebug, format, args...)
}

// Infof emits an info level log line with the given format string and arguments
func (l *Logger) Infof(format string, args ...any) {
	// We could delegate log suppression to Logf() but
	// this simple condition saves a func call
	if l.threshold > LevelInfo {
		return
	}
	l.Logf(LevelInfo, format, args...)
}

// Errorf emits an error level log line with the given format string and arguments
func (l *Logger) Errorf(format string, args ...any) {
	l.Logf(LevelError, format, args...)
}
