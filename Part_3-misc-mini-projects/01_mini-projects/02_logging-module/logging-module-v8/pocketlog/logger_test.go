package pocketlog_test

import (
	"strings"
	"testing"
	"time"

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

/*
	Max message length
*/
func TestDefaultMaxMsgLenIs1000NoArgs(t *testing.T) {
	var longMsg string
	for i := 0; i < 1001; i++ {
		longMsg += "a"
	}
	tw := &testWriter{}
	l := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(tw))

	l.Logf(pocketlog.LevelDebug, longMsg)

	if len(tw.contents) != 1000 {
		t.Errorf("expected log message of len 1000, but was %v", len(tw.contents))
	}

	if tw.contents[995:1000] != "a...\n" {
		t.Errorf("expected end string to be %q but was %q", "a...\n", tw.contents[995:1000])
	}
}

func TestDefaultMaxMsgLenIs1000WithArgs(t *testing.T) {
	var longMsg string
	for i := 0; i < 1000 - len("[DEBUG] ") - 5; i++ {
		longMsg += "a"
	}
	longMsg += "%s"

	tw := &testWriter{}
	l := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(tw))

	l.Logf(pocketlog.LevelDebug, longMsg, "1234567890")

	if len(tw.contents) != 1000 {
		t.Errorf("expected log message of len 1000, but was %v", len(tw.contents))
	}

	if tw.contents[995:1000] != "1...\n" {
		t.Errorf("expected end string to be %q but was %q", "1...\n", tw.contents[995:1000])
	}
}

func TestCustomMaxMsgLenNoArgs(t *testing.T) {
	logMsg := "1234567890ABCDEF"
	maxMsgLen := 15
	tw := &testWriter{}
	l := pocketlog.New(
		pocketlog.LevelDebug,
		pocketlog.WithOutput(tw),
		pocketlog.WithMaxMsgLen(maxMsgLen))
	expectedLogMsg := "[DEBUG] 123...\n"

	l.Logf(pocketlog.LevelDebug, logMsg)

	if (len(tw.contents) != maxMsgLen) {
		t.Errorf("expected log message to be %v long, but was %v", maxMsgLen, len(tw.contents))
	}

	if tw.contents != expectedLogMsg {
		t.Errorf("expected log message to be %q, but was %q", expectedLogMsg, tw.contents)
	}
}

func TestCustomMaxMsgLenWithArgs(t *testing.T) {
	logMsg := "%s1234567890ABCDEF%s"
	maxMsgLen := 20
	tw := &testWriter{}
	l := pocketlog.New(
		pocketlog.LevelDebug,
		pocketlog.WithOutput(tw),
		pocketlog.WithMaxMsgLen(maxMsgLen))
	expectedLogMsg := "[DEBUG] abc12345...\n"

	l.Logf(pocketlog.LevelDebug, logMsg, "abc", "XYZ")

	if (len(tw.contents) != maxMsgLen) {
		t.Errorf("expected log message to be %v long, but was %v", maxMsgLen, len(tw.contents))
	}

	if tw.contents != expectedLogMsg {
		t.Errorf("expected log message to be %q, but was %q", expectedLogMsg, tw.contents)
	}
}

/*
	Timestamp tests
*/
func TestWithTime(t *testing.T) {

	logMsg := "Hello, world!"
	tw := &testWriter{}
	l := pocketlog.New(
		pocketlog.LevelDebug,
		pocketlog.WithOutput(tw),
		pocketlog.WithTimestamp())

	l.Logf(pocketlog.LevelDebug, logMsg)

	loglinePrefix := tw.contents[0:strings.Index(tw.contents, " [DEBUG]")]
	_, err := time.Parse(time.RFC3339, loglinePrefix)

	if err != nil {
		t.Errorf("expected log message to start with a timestamp but found %v", err)
	}

}
