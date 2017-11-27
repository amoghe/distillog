package distillog

//
// This file exposes package level logging functions so that this package
// can be used directly for logging (to stderr) without needing the caller
// to instantiate a logger (a la stdlib)
//

var (
	std = NewStderrLogger("")
)

func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v)
}

func Debugln(v ...interface{}) {
	std.Debugln(v)
}

func Infof(format string, v ...interface{}) {
	std.Infof(format, v)
}

func Infoln(v ...interface{}) {
	std.Infoln(v)
}

func Warningf(format string, v ...interface{}) {
	std.Warningf(format, v)
}

func Warningln(v ...interface{}) {
	std.Warningln(v)
}

func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v)
}

func Errorln(v ...interface{}) {
	std.Errorln(v)
}

