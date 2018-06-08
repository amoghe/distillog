// +build linux darwin freebsd netbsd openbsd !windows

package distillog

import (
	"fmt"
	"log/syslog"
)

// wraps around a syslog.Writer to make it adhere to the `Logger` interface
// exposed by this package
type wrappedSyslogWriter struct {
	writer *syslog.Writer
}

// NewSyslogLogger returns a Logger that sends messages to the Syslog daemon.
// This will panic if it is unable to connect to the local syslog daemon.
func NewSyslogLogger(tag string) Logger {
	l, err := syslog.New(syslog.LOG_DAEMON, tag)
	if err != nil {
		panic(err)
	}
	return &wrappedSyslogWriter{l}
}

func (w *wrappedSyslogWriter) Debugf(f string, v ...interface{}) {
	w.writer.Debug(fmt.Sprintf(f, v...))
}
func (w *wrappedSyslogWriter) Debugln(v ...interface{}) {
	w.writer.Debug(fmt.Sprintln(v...))
}
func (w *wrappedSyslogWriter) Infof(f string, v ...interface{}) {
	w.writer.Info(fmt.Sprintf(f, v...))
}
func (w *wrappedSyslogWriter) Infoln(v ...interface{}) {
	w.writer.Info(fmt.Sprintln(v...))
}
func (w *wrappedSyslogWriter) Warningf(f string, v ...interface{}) {
	w.writer.Warning(fmt.Sprintf(f, v...))
}
func (w *wrappedSyslogWriter) Warningln(v ...interface{}) {
	w.writer.Warning(fmt.Sprintln(v...))
}
func (w *wrappedSyslogWriter) Errorf(f string, v ...interface{}) {
	w.writer.Err(fmt.Sprintf(f, v...))
}
func (w *wrappedSyslogWriter) Errorln(v ...interface{}) {
	w.writer.Err(fmt.Sprintln(v...))
}
func (w *wrappedSyslogWriter) Close() error { return nil }
