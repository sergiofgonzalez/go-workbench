package pocketlog

import "io"

// Option is ligtweight extensibility mechanism for the Logger consisting
// of functions that can be used to tailor the Logger capabilities if required
type Option func(*Logger)

// WithOutput lets you configure a custom output of type io.Writer for your
// Logger
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}

// WithMaxMsgLen lets you configure the maximum message length of log messages
func WithMaxMsgLen(maxMsgLen int) Option {
	return func(l *Logger) {
		l.maxMsgLen = maxMsgLen
	}
}

// WithTimestamp lets you timestamp the emitted log line messages
func WithTimestamp() Option {
	return func(l *Logger) {
		l.timestamp = true
	}
}

func WithJson() Option {
	return func(l *Logger) {
		l.useJson = true
	}
}