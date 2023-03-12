package pocketlog_test

import "example.com/pocketlog/pocketlog"

/*
	Logf tests
*/

func Example_logf_debugThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Logf(pocketlog.LevelDebug, "Hello!")
	// Output:
	// [DEBUG] Hello!
}

func Example_logf_debugThreshold_debugArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Logf(pocketlog.LevelDebug, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [DEBUG] Hello to Jason "Isaacs"!
}

func Example_logf_debugThreshold_infoNoArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Logf(pocketlog.LevelInfo, "Hello!")
	// Output:
	// [INFO]  Hello!
}

func Example_logf_debugThreshold_infoArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Logf(pocketlog.LevelInfo, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [INFO]  Hello to Jason "Isaacs"!
}

func Example_logf_debugThreshold_errorNoArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Logf(pocketlog.LevelError, "Hello!")
	// Output:
	// [ERROR] Hello!
}

func Example_logf_debugThreshold_errorArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Logf(pocketlog.LevelError, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [ERROR] Hello to Jason "Isaacs"!
}

func Example_logf_infoThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Logf(pocketlog.LevelDebug, "Hello!")
	// Output:
	//
}

func Example_logf_infoThreshold_debugArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Logf(pocketlog.LevelDebug, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	//
}

func Example_logf_infoThreshold_infoNoArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Logf(pocketlog.LevelInfo, "Hello!")
	// Output:
	// [INFO]  Hello!
}

func Example_logf_infoThreshold_infoArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Logf(pocketlog.LevelInfo, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [INFO]  Hello to Jason "Isaacs"!
}

func Example_logf_infoThreshold_errorNoArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Logf(pocketlog.LevelError, "Hello!")
	// Output:
	// [ERROR] Hello!
}

func Example_logf_infoThreshold_errorArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Logf(pocketlog.LevelError, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [ERROR] Hello to Jason "Isaacs"!
}

func Example_logf_errorThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Logf(pocketlog.LevelDebug, "Hello!")
	// Output:
	//
}

func Example_logf_errorThreshold_debugArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Logf(pocketlog.LevelDebug, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	//
}

func Example_logf_errorThreshold_infoNoArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Logf(pocketlog.LevelInfo, "Hello!")
	// Output:
	//
}

func Example_logf_errorThreshold_infoArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Logf(pocketlog.LevelInfo, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	//
}

func Example_logf_errorThreshold_errorNoArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Logf(pocketlog.LevelError, "Hello!")
	// Output:
	// [ERROR] Hello!
}

func Example_logf_errorThreshold_errorArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Logf(pocketlog.LevelError, "Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [ERROR] Hello to Jason "Isaacs"!
}

/*
	Debugf tests
*/

func Example_debugf_debugThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Debugf("Hello!")
	// Output:
	// [DEBUG] Hello!
}

func Example_debugf_debugThreshold_debugArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Debugf("Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [DEBUG] Hello to Jason "Isaacs"!
}

func Example_debugf_infoThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Debugf("Hello!")
	// Output:
	//
}

func Example_debugf_errorThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Debugf("Hello!")
	// Output:
	//
}

/*
	Infof tests
*/

func Example_infof_debugThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Infof("Hello!")
	// Output:
	// [INFO]  Hello!
}

func Example_infof_debugThreshold_debugArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Infof("Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [INFO]  Hello to Jason "Isaacs"!
}

func Example_infof_infoThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Infof("Hello!")
	// Output:
	// [INFO]  Hello!
}

func Example_infof_errorThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Infof("Hello!")
	// Output:
	//
}

/*
	Errorf tests
*/

func Example_errorf_debugThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Errorf("Hello!")
	// Output:
	// [ERROR] Hello!
}

func Example_errorf_debugThreshold_debugArgs() {
	l := pocketlog.New(pocketlog.LevelDebug)

	l.Errorf("Hello to %s %q!", "Jason", "Isaacs")
	// Output:
	// [ERROR] Hello to Jason "Isaacs"!
}

func Example_errorf_infoThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelInfo)

	l.Errorf("Hello!")
	// Output:
	// [ERROR] Hello!
}

func Example_errorf_errorThreshold_debugNoArgs() {
	l := pocketlog.New(pocketlog.LevelError)

	l.Errorf("Hello!")
	// Output:
	// [ERROR] Hello!
}

/*
	Default initialization tests
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
