package pocketlog

import (
	"fmt"
)

// Logger allows an application to log info with different levels
type Logger struct {
}

// Logf emits a log line with the corresponding level with the
// given format string and arguments
func (l *Logger) Logf(lvl Level, format string, args ...any) {
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
	l.Logf(LevelDebug, format, args...)
}

// Infof emits an info level log line with the given format string and arguments
func (l *Logger) Infof(format string, args ...any) {
	l.Logf(LevelInfo, format, args...)
}

// Errorf emits an error level log line with the given format string and arguments
func (l *Logger) Errorf(format string, args ...any) {
	l.Logf(LevelError, format, args...)
}
