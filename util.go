package distillog

import (
	"io"
	"runtime"
)

// wrap around a io.Writer, and provide a dummy Close method.
type noopCloser struct {
	io.Writer
}

func (n noopCloser) Close() error { return nil }

// If we ever want to print callers file:line info in the message.
func callerFileLine() (string, int) {
	if _, file, line, ok := runtime.Caller(3); ok {
		return file, line
	}
	return "???", 0

}
