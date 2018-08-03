// Package logger provides a single logging facility/interface that can be used with
// Appengine-contained code (e.g. GCP Standard Enviromnment for Go) or inside 'normal' Go runtime
// environments (GCP Flexible and most others).
//
// This package was written specifically to enable same-logging-package usage in both types of environments,
// and to facilitate porting code from the Standard Environment to the Flexible Environment.
//
// Appengine logging differs in two key ways from plain-jane go log:
//
// 1) It uses level-tied methods for its logging, e.g.:
//  - Debugf
//  - Infof 
//  - Warningf 
//  - Errorf 
// 
// 2) Appengine logging always requires a context as the first argument. The GCP logging infrastructure uses
// the context to tie a log line to a particular request, and also does not accepts any logs that are not
// tied to a particular request via a context. (You can see comtext-less logs in development, but not in GCP
// infrastructure)
// 
// This logger uses the appengine-style interface, but let's you use it with appengine-hosted code or not.
// It uses the 'appengine' build tag to conditionally build the right logger for the environment. 
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"github.com/efixler/config"
)

const LogPrefixConfigKey = "LOG_PREFIX"
const LogLevelConfKey = "LOG_LEVEL"

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

var (
	Request *RequestLogger
	Context *ContextLogger
	Std		*StdLogger
	// NB: Package-wide log-level setting is ignored when using appengine logging providers
	Level LogLevel
	headers = [4]string{"DEBUG: ","INFO: ","WARNING: ","ERROR: "}
)

type StdLogger struct {
	dlog *log.Logger
	elog *log.Logger
}

func newStdLogger() *StdLogger {
	prefix := config.Default().GetOrDefault(LogPrefixConfigKey, "")
	if len(prefix) > 0 {
		prefix = prefix + " "
	}
	d := log.New(os.Stderr, prefix, log.Ldate|log.Ltime|log.LUTC)
	e := log.New(os.Stderr, prefix, log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	sl := &StdLogger{ dlog: d, elog: e }
	return sl
}

func (l *StdLogger) Debugf(format string, args ...interface{}) {
	l.output(Debug, 1, format, args...)
}

func (l *StdLogger) Infof(format string, args ...interface{}) {
	l.output(Info, 1, format, args...)
}

func (l *StdLogger) Warningf(format string, args ...interface{}) {
	l.output(Warning, 1, format, args...)
}

func (l *StdLogger) Errorf(format string, args ...interface{}) {
	l.output(Error, 1, format, args...)
}

func (l *StdLogger) SetOutput(w io.Writer) {
	l.dlog.SetOutput(w)
	l.elog.SetOutput(w)
}

func (l *StdLogger) setFlags(flag int) {
	l.dlog.SetFlags(flag)
	l.elog.SetFlags(flag)
}


// This function is provided so in-package callers can adjust the stack backcount
func (l *StdLogger) output(level LogLevel, stackAdjust int, format string, args ...interface{}) {
	if level < Level {
		return
	} 
	ll := l.dlog
	switch level {
		case Error: fallthrough
		case Warning: ll = l.elog
	}
	if int(level) < len(headers) && int(level) >= 0 {
		format = headers[int(level)] + format
	}
	ll.Output(stackAdjust + 2, fmt.Sprintf(format, args...))
}

func levelStringToLogLevel(ls string) LogLevel {
	ls = strings.ToUpper(strings.TrimSpace(ls))
	switch ls {
		case "ERROR": return Error
		case "WARNING": return Warning
		case "INFO": return Info
		case "DEBUG": fallthrough
		default: 	return Debug
	}
}


func init() {
	Level = levelStringToLogLevel(config.Default().GetOrDefault(LogLevelConfKey, "DEBUG"))
	Std = newStdLogger()
}
