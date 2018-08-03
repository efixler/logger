// +build !appengine

package logger

import (
	"context"
	"net/http"
)


type ContextLogger struct {

}


func (l *ContextLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.output(ctx, Debug, 1, format, args...)
}

func (l *ContextLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.output(ctx, Info, 1, format, args...)
}

func (l *ContextLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	l.output(ctx, Warning, 1, format, args...)
}

func (l *ContextLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.output(ctx, Error, 1, format, args...)
}

func (l *ContextLogger) output(ctx context.Context, level LogLevel, stackAdjust int, format string, args ...interface{}) {
	Std.output(level, stackAdjust + 1, format, args...)
}

func init() {
	Context = &ContextLogger{}
}

type RequestLogger struct {

}

func newRequestLogger() *RequestLogger {
	return &RequestLogger{}
}

func (l *RequestLogger) Debugf(r *http.Request, format string, args ...interface{}) {
	Context.output(r.Context(), Debug, 1, format, args...)
}

func (l *RequestLogger) Infof(r *http.Request, format string, args ...interface{}) {
	Context.output(r.Context(), Info, 1, format, args...)
}

func (l *RequestLogger) Errorf(r *http.Request, format string, args ...interface{}) {
	Context.output(r.Context(), Error, 1, format, args...)
}

func (l *RequestLogger) Warningf(r *http.Request, format string, args ...interface{}) {
	Context.output(r.Context(), Warning, 1, format, args...)
}

func init() {
	Request = newRequestLogger()
}
