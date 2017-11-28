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

// appends a fixed width string 'str' into byte buffer 'b'. Appends spaces if 'str' is too short.
func fixedWidthStr(width int, str string, b []byte) []byte {
	// Write as many bytes as 'width', writing spaces if we run out of chars
	for i := 0; i < width; i++ {
		if i < len(str) {
			b = append(b, str[i])
		} else {
			b = append(b, ' ')
		}
	}
	return b
}
