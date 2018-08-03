package logger 

import (
	//"fmt"
	"log"
	"os"
	"testing"
)

func Example() {
	Std.SetOutput(os.Stdout)
	Std.Debugf("This is a %s log message", "simple")
	Std.Debugf("Timestamps are suppressed for the test example, but will appear in runtime output.")
	Std.Infof("The interfaces for logger.Request and logger.Context are similar, except they take %s and %s as their first arguments, respectively.",
			"*http.Request", "context.Context")
	Std.Warningf("Warning and Error level messages also provide file and line number")
	Std.Infof("When used inside appengine, output will be a little bit different.")
	Level = Warning
	Std.Infof("Log level can be set on a package-wide basis. In appengine, logs are transmitted regardless of level.")
	
	// Output:
	// DEBUG: This is a simple log message
	// DEBUG: Timestamps are suppressed for the test example, but will appear in runtime output.
	// INFO: The interfaces for logger.Request and logger.Context are similar, except they take *http.Request and context.Context as their first arguments, respectively.
	// generic_example_test.go:16: WARNING: Warning and Error level messages also provide file and line number
	// INFO: When used inside appengine, output will be a little bit different.
	
}

func TestMain(m *testing.M) {
	Std.dlog.SetFlags(0)
	Std.elog.SetFlags(log.Lshortfile)
	ecode := m.Run()
	os.Exit(ecode)
}
