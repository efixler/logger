// +build appengine

package logger

import (
	"context"
	"net/http"
	aelog "google.golang.org/appengine/log"
)

type ContextLogger struct {

}


func (l *ContextLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	aelog.Debugf(ctx, format, args...)
}

func (l *ContextLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	aelog.Infof(ctx, format, args...)
}

func (l *ContextLogger) Warningf(ctx context.Context, format string, args ...interface{}) {
	aelog.Warningf(ctx, format, args...)
}

func (l *ContextLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	aelog.Errorf(ctx, format, args...)
}

type RequestLogger struct {

}

func newRequestLogger() *RequestLogger {
	return &RequestLogger{}
}

func (l *RequestLogger) Debugf(r *http.Request, format string, args ...interface{}) {
	Context.Debugf(r.Context(), format, args...)
}

func (l *RequestLogger) Infof(r *http.Request, format string, args ...interface{}) {
	Context.Infof(r.Context(), format, args...)
}

func (l *RequestLogger) Errorf(r *http.Request, format string, args ...interface{}) {
	Context.Errorf(r.Context(), format, args...)
}

func (l *RequestLogger) Warningf(r *http.Request, format string, args ...interface{}) {
	Context.Warningf(r.Context(), format, args...)
}


func init() {
	Request = newRequestLogger()
}
