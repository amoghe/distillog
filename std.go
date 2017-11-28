package distillog

import (
	"io"
)

//
// This file exposes package level logging functions so that this package
// can be used directly for logging (to stderr) without needing the caller
// to instantiate a logger (a la stdlib)
//

var (
	std = NewStderrLogger("")
)

// Debugf logs a message to stderr at 'debug' level
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

// Debugln logs a message to stderr at 'debug' level
func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

// Infof logs a message to stderr at 'info' level
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

// Infoln logs a message to stderr at 'info' level
func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

// Warningf logs a message to stderr at 'warn' level
func Warningf(format string, v ...interface{}) {
	std.Warningf(format, v...)
}

// Warningln logs a message to stderr at 'warn' level
func Warningln(v ...interface{}) {
	std.Warningln(v...)
}

// Errorf logs a message to stderr at 'error' level
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// Errorln logs a message to stderr at 'error' level
func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

// Close closes the stream to which the default logger logs to
func Close() {
	std.Close()
}

// SetStream allows you to configure package level logger to emit to the
// specified stream. NOTE: this is not safe when called concurrently from
// multiple routines. This should typically be called during program startup
func SetStream(s io.WriteCloser) {
	std = NewStreamLogger("", s)
}
