package pocketlog_test

import (
	"testing"

	"example.com/pocketlog/pocketlog"
)

/*
	Using default output
*/

func Example_logf_debugThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Logf(pocketlog.LevelDebug, "Hello!")
	// Output:
	// [DEBUG] Hello!
}

/*
	Initialization tests (when not using New())
*/
func Example_logf_structInit_defaultThresIsDebug() {
	l := pocketlog.Logger{}

	l.Logf(pocketlog.LevelDebug, "Hello!")
	// Output:
	// [DEBUG] Hello!
}

func Example_debugf_structInit_defaultThresIsDebug() {
	l := pocketlog.Logger{}

	l.Debugf("Hello!")
	// Output:
	// [DEBUG] Hello!
}

func Example_logf_zeroValue_defaultThresIsDebug() {
	var l pocketlog.Logger

	l.Logf(pocketlog.LevelDebug, "Hello!")
	// Output:
	// [DEBUG] Hello!
}

func Example_debugf_zeroValue_defaultThresIsDebug() {
	var l pocketlog.Logger

	l.Debugf("Hello!")
	// Output:
	// [DEBUG] Hello!
}



/*
	Custom output tests
*/

// a testWriter is a mock implementation of an io.Writer for the tests
type testWriter struct {
	contents string
}


// Write implementatiton for a testWriter receiver
func (tw *testWriter) Write(p []byte) (int, error) {
	tw.contents += string(p)
	return len(p), nil
}

// testMessage represents the content to log which is used in the logger methods
type testMessage struct {
	format 	string
	args		[]any
}

// testCase represents the data to be used in each subtest and the expected
// outcome
type testCase struct {
	logLevel 		pocketlog.Level
	messages		[]testMessage
	expectedOut	string
}

// tests represents all the test cases to run indexed by a string title for the
// test
type tests map[string]testCase

// messagesToLog is the collection of testMessages to use. The first element
// will be used as a debug message, the second as an info message, and the third
// one as an error message
var messagesToLog = []testMessage{
	// debug message
	{
		format: "Hello to %s!",
		args: []any{"Idris"},
	},
	// info message
	{
		format: "%s. How are you?",
		args: []any{"Jason"},
	},
	// error message
	{
		format: "My fav numbers are %d and %v",
		args: []any{5, 14},
	},
}

// an enum to simplify access to the messagesToLog array
const (
	debugTestMessage int = iota
	infoTestMessage
	errorTestMessage
)

var testsToRun = tests{
	"log level=debug": {
		logLevel: pocketlog.LevelDebug,
		messages: messagesToLog,
		expectedOut: "[DEBUG] Hello to Idris!\n" +
								 "[INFO]  Jason. How are you?\n" +
								 "[ERROR] My fav numbers are 5 and 14\n",
	},
	"log level=info": {
		logLevel: pocketlog.LevelInfo,
		messages: messagesToLog,
		expectedOut: "[INFO]  Jason. How are you?\n" +
								 "[ERROR] My fav numbers are 5 and 14\n",
	},
	"log level=error": {
		logLevel: pocketlog.LevelError,
		messages: messagesToLog,
		expectedOut: "[ERROR] My fav numbers are 5 and 14\n",
	},
}

func TestLogFWithCustomOutput(t *testing.T) {
	// Arrange

	for subtestLabel, subtestCase := range testsToRun {
		t.Run(subtestLabel, func(t *testing.T) {
			tw := &testWriter{}
			l := pocketlog.New(subtestCase.logLevel, pocketlog.WithOutput(tw))

			l.Logf(
				pocketlog.LevelDebug,
				subtestCase.messages[debugTestMessage].format,
				subtestCase.messages[debugTestMessage].args...)

			l.Logf(
				pocketlog.LevelInfo,
				subtestCase.messages[infoTestMessage].format,
				subtestCase.messages[infoTestMessage].args...)

			l.Logf(
				pocketlog.LevelError,
				subtestCase.messages[errorTestMessage].format,
				subtestCase.messages[errorTestMessage].args...)

			actual := tw.contents
			if actual != subtestCase.expectedOut {
				t.Errorf("Logf tests: %v: expected %q, but found %q", subtestLabel, subtestCase.expectedOut, actual)
			}
		})
	}
}

func TestDebugfInfofErrorfWithCustomOutput(t *testing.T) {
	// Arrange

	for subtestLabel, subtestCase := range testsToRun {
		t.Run(subtestLabel, func(t *testing.T) {
			tw := &testWriter{}
			l := pocketlog.New(subtestCase.logLevel, pocketlog.WithOutput(tw))

			l.Debugf(subtestCase.messages[debugTestMessage].format, subtestCase.messages[debugTestMessage].args...)
			l.Infof(subtestCase.messages[infoTestMessage].format, subtestCase.messages[infoTestMessage].args...)
			l.Errorf(subtestCase.messages[errorTestMessage].format, subtestCase.messages[errorTestMessage].args...)

			actual := tw.contents
			if actual != subtestCase.expectedOut {
				t.Errorf("Logger method tests: %v: expected %q, but found %q", subtestLabel, subtestCase.expectedOut, actual)
			}
		})
	}
}
