package pocketlog

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Logger allows an application to log info with different levels.
type Logger struct {
	threshold Level
	maxMsgLen int
	timestamp bool
	useJson   bool
	output		io.Writer
}

// New returns a Logger variable initialized with the given threshold for the
// logging level and a sequence of zer or more options that configure the Logger
// behavior.
func New(threshold Level, options ...Option) *Logger {

	l := Logger{
		threshold: threshold,
		output: os.Stdout,			// default output is stdout
		maxMsgLen: 1000,				// default is 1000 runes
		useJson: false,					// default is non-JSON
	}

	for _, cfgFunc := range options {
		cfgFunc(&l)
	}

	return &l
}

// Logf emits a log line with the corresponding level with the
// given format string and arguments
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if lvl < l.threshold {
		return
	}

	if l.output == nil {
		l.output = os.Stdout
	}

	if l.maxMsgLen == 0 {
		l.maxMsgLen = 1000
	}

	if l.useJson {
		logMsgFmt := fmt.Sprintf(format, args...)
		if len(logMsgFmt) > l.maxMsgLen {
			logMsgFmt = logMsgFmt[0:l.maxMsgLen - 3] + "..."
		}

		logLineFmt := `{ "level": %v, "time": %v, "msg": %q }`
		_, _ = fmt.Fprintf(l.output, logLineFmt, lvl, time.Now().UnixMilli(), logMsgFmt)

	} else {
		var strLogLevel string
		switch lvl {
		case LevelDebug:
			strLogLevel = "[DEBUG] "
		case LevelInfo:
			strLogLevel = "[INFO]  "
		case LevelError:
			strLogLevel = "[ERROR] "
		}

		if l.timestamp {
			strLogLevel = time.Now().Format(time.RFC3339) + " " + strLogLevel
		}

		loglineFormat := fmt.Sprintf(strLogLevel + format + "\n", args...)
		if len(loglineFormat) > l.maxMsgLen {
			loglineFormat = loglineFormat[0:l.maxMsgLen - 4] + "..." + "\n"
		}
		_, _ = fmt.Fprint(l.output, loglineFormat)
	}
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
