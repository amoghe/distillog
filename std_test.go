package distillog

import (
	"strings"
	"testing"
)

func TestSetStream(t *testing.T) {
	strm := &dummyStream{}

	SetOutput(strm)
	Infoln("random message")

	if !strings.Contains(strm.String(), "random message") {
		t.Error("expected string not found in stream buffer")
	}
}
