package distillog

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

// streamLogger implements the Logger interface and writes to the specified io.Writer ("stream").
// It mimics the stdlib logger including memory optimizations such as minimizing calls to fmt.Sprintf
// and using a shared buffer to format the message before writing it out.
type streamLogger struct {
	stream  io.WriteCloser
	tag     string
	linebuf []byte
	lock    sync.Mutex
}

// NewStderrLogger returns a Logger that outputs messages to Stderr.
func NewStderrLogger(tag string) Logger {
	return &streamLogger{
		tag:     tag,
		linebuf: []byte{},
		stream:  noopCloser{os.Stderr},
	}
}

// NewStdoutLogger returns a Logger that outputs messages to Stdout.
func NewStdoutLogger(tag string) Logger {
	return &streamLogger{
		tag:     tag,
		linebuf: []byte{},
		stream:  noopCloser{os.Stdout},
	}
}

// NewNullLogger returns a logger that drops messages (outputs to /dev/null).
func NewNullLogger(tag string) Logger {
	return &streamLogger{
		tag:     tag,
		linebuf: []byte{},
		stream:  noopCloser{ioutil.Discard},
	}
}

// NewStreamLogger returns a Logger that outputs messages to the specified stream.
func NewStreamLogger(tag string, stream io.WriteCloser) Logger {
	return &streamLogger{
		tag:     tag,
		linebuf: []byte{},
		stream:  stream,
	}
}

// writes a formatted message (w/ timestamp, level) to the output stream.
func (w *streamLogger) output(timeStr, level, msg string) {
	// We need to serialize access to the linebuffer that is used to assemble the message \
	// as well as the output stream we will print to.
	w.lock.Lock()

	// save memory, (re)use a buffer instead of relying on fmt.Sprintf to format the output string
	w.linebuf = w.linebuf[:0]

	w.linebuf = append(w.linebuf, timeStr...)
	w.linebuf = append(w.linebuf, ' ')
	w.linebuf = append(w.linebuf, w.tag...)
	w.linebuf = append(w.linebuf, ' ')

	w.linebuf = append(w.linebuf, '[')
	w.linebuf = fixedWidthStr(5, level, w.linebuf)
	w.linebuf = append(w.linebuf, ']')

	w.linebuf = append(w.linebuf, ' ')
	w.linebuf = append(w.linebuf, msg...)

	if len(msg) == 0 || msg[len(msg)-1] != '\n' {
		w.linebuf = append(w.linebuf, '\n')
	}

	w.stream.Write(w.linebuf)
	w.lock.Unlock()
}

func (w *streamLogger) Debugf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "DEBUG", msg)
}

func (w *streamLogger) Debugln(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "DEBUG", msg)
}

func (w *streamLogger) Infof(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "INFO", msg)
}

func (w *streamLogger) Infoln(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "INFO", msg)
}

func (w *streamLogger) Warningf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "WARN", msg)
}

func (w *streamLogger) Warningln(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "WARN", msg)
}

func (w *streamLogger) Errorf(f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "ERROR", msg)
}

func (w *streamLogger) Errorln(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	now := time.Now().Format(timeFormatStr)
	w.output(now, "ERROR", msg)
}

func (w *streamLogger) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	return w.stream.Close()
}
