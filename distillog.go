package distillog

const (
	timeFormatStr = "Mon Jan 2 15:04:05"
)

// Logger defines a distilled interface for logging messages from your program.
// Note: All functions append a trailing newline if one doesn't exist.
type Logger interface {
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})

	Infof(format string, v ...interface{})
	Infoln(v ...interface{})

	Warningf(format string, v ...interface{})
	Warningln(v ...interface{})

	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})

	Close() error
}
