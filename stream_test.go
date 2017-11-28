package distillog

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type dummyStream struct {
	bytes.Buffer
}

func (d *dummyStream) Close() error { return nil }

func TestOutput(t *testing.T) {

	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}
	tlog.output("12345", "LEVEL", "message")

	got := strm.String()
	exp := "12345 TAG [LEVEL] message\n"
	if got != exp {
		t.Error("unexpected output", got, "vs", exp)
	}
}

func TestLevelDebugf(t *testing.T) {
	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}

	msg := "informational message"
	tlog.Debugf("%s", msg)

	got := strm.String()
	exp := fmt.Sprintf("[DEBUG] %s", msg)

	if !strings.Contains(got, exp) {
		t.Error("output does not contain message", got)
	}
}

func TestLevelDebugln(t *testing.T) {
	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}

	msg := "informational message"
	tlog.Debugln(msg)

	got := strm.String()
	exp := fmt.Sprintf("[DEBUG] %s", msg)

	if !strings.Contains(got, exp) {
		t.Error("output does not contain message", got)
	}
}

func TestLevelInfoln(t *testing.T) {
	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}

	msg := "informational message"
	tlog.Infoln(msg)

	got := strm.String()
	exp := fmt.Sprintf("[INFO ] %s", msg)

	if !strings.Contains(got, exp) {
		t.Error("output does not contain message", got)
	}
}

func TestLevelInfof(t *testing.T) {
	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}

	msg := "informational message"
	tlog.Infof("%s", msg)

	got := strm.String()
	exp := fmt.Sprintf("[INFO ] %s", msg)

	if !strings.Contains(got, exp) {
		t.Error("output does not contain message", got)
	}
}

func TestLevelWarningln(t *testing.T) {
	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}

	msg := "informational message"
	tlog.Warningln(msg)

	got := strm.String()
	exp := fmt.Sprintf("[WARN ] %s", msg)

	if !strings.Contains(got, exp) {
		t.Error("output does not contain message", got)
	}
}

func TestLevelWarningf(t *testing.T) {
	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}

	msg := "informational message"
	tlog.Warningf("%s", msg)

	got := strm.String()
	exp := fmt.Sprintf("[WARN ] %s", msg)

	if !strings.Contains(got, exp) {
		t.Error("output does not contain message", got)
	}
}

func TestLevelErrorln(t *testing.T) {
	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}

	msg := "informational message"
	tlog.Errorln(msg)

	got := strm.String()
	exp := fmt.Sprintf("[ERROR] %s", msg)

	if !strings.Contains(got, exp) {
		t.Error("output does not contain message", got)
	}
}

func TestLevelErrorf(t *testing.T) {
	strm := dummyStream{}
	tlog := &streamLogger{
		tag:     "TAG",
		linebuf: []byte{},
		stream:  &strm,
	}

	msg := "informational message"
	tlog.Errorf("%s", msg)

	got := strm.String()
	exp := fmt.Sprintf("[ERROR] %s", msg)

	if !strings.Contains(got, exp) {
		t.Error("output does not contain message", got)
	}
}

func BenchmarkThroughput(b *testing.B) {
	tlog := NewNullLogger("")

	runParallelBody := func(pb *testing.PB) {
		for pb.Next() {
			tlog.Infoln("one", "iteration")
		}
	}

	b.RunParallel(runParallelBody)
}

// func BenchmarkStdlibThroughput(b *testing.B) {
// 	log.SetOutput(ioutil.Discard)

// 	runParallelBody := func(pb *testing.PB) {
// 		for pb.Next() {
// 			log.Println("one", "iteration")
// 		}
// 	}

// 	b.RunParallel(runParallelBody)
// }
