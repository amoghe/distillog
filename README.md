[![Build Status](https://travis-ci.org/amoghe/distillog.svg)](https://travis-ci.org/amoghe/distillog)
[![Documentation](https://godoc.org/github.com/amoghe/distillog?status.svg)](http://godoc.org/github.com/amoghe/distillog)
[![Go Report Card](https://goreportcard.com/badge/github.com/amoghe/distillog)](https://goreportcard.com/report/github.com/amoghe/distillog)

# What is `distillog`?

`distillog` aims to offer a minimalistic logging interface that also supports
log levels. It takes the `stdlib` API and only slightly enhances it. Hence, you
could think of it as levelled logging, _distilled_.

# Yet _another_ logging library for go(lang)?

> Logging libraries are like opinions, everyone seems to have one -- Anon(?)

Most other logging libraries do either __too little__ ([stdlib][0])
or __too much__ ([glog][1]).

As with most other libraries, this one is opinionated. In terms of functionality
it exposes, it attempts to sit somewhere between the stdlib and the majority of
other logging libraries available (but leans mostly towards the spartan side
of stdlib).

## The stdlib does _too little_, you say?

Just a smidge.

Presenting varying levels of verbosity (or severity) are an important part of
what makes a program more usable or debuggable. For example, `debug` or `info`
level messages may be useful to the developers during the development cycle.
These messages may be dropped or suppressed in production since they are not
useful to everyone. Similarly `warning` messages may be emitted when a error has
been gracefully handled but the program would like to notify its human overlords
of some impending doom.

In most cases, some downstream entity "knows" how to filter the messages and
keep those that are relevant to the environment. As evidence of this, most
other languages have log libraries that support levels. Similarly some programs
offer varying verbosity levels (e.g. `-v`, `-vv` etc). The golang stdlib takes
a much more spartan approach (exposing only `Println` and friends) so using it
in programs to emit messages of varying interest/levels can get tedious (manual
prefixes, anyone?). This is where `distillog` steps in. It aims to slightly
improve on this minimalstic logging interface. __Slightly__.

## Other libraries do _too much_, you say?

Ever used `log.Panicf` or `log.Fatalf`? Exiting your program is *not* something
your log library should be doing! Similarly, other libraries offer options for
maintaining old log files and rotating them. Your logging library shouldn't need
to care about this. Whatever facility (other libraries call this a "backend")
messages are sent to should determine how old messages are handled. `distillog`
prefers that you use `lumberjack` (or an equivalent WriteCloser) depending on
where you choose to persist the messages.

> But log file rotation is absolutely necessary!

Agreed, and someone's gotta do it, but it need not be your logging library!

You can use `distillog` along with a [lumberjack][2] "backend". It provides an
`io.WriteCloser` which performs all the magic you need. Initialize a logger
using `distillog.NewStream`, pass it an instance of the `io.WriteCloser`
that lumberjack returns, _et voila_, you have a logger that does what you need.

## And how is `distillog` different?

`distillog` aims to offer a only slightly richer interface than the stdlib.

To this end, it restricts itself to:
- presenting a minimal interface so that you can emit levelled log messages
- providing logger implementations for logging to the most common backends
	- streams - e.g. stderr/stdout 
	- files - anything via `io.WriteCloser` (via `lumberjack`)
	- syslog
- avoid taking on any non-essential responsibilities (colors, _ahem_)
- expose a logger interface, instead of an implementation

## Expose an interface? Why?

By exposing an interface you can write programs that use levelled log messages,
but switch between logging to various facilities by simply instantiating the
appropriate logger as determined by the caller (Your program can offer a
command-line switch like so - `--log-to=[syslog,stderr,<file>]` and the simply
instantiate the appropriate logger).

# Usage/examples:

As seen in the [godoc](https://godoc.org/github.com/amoghe/distillog#Logger),
the interface is limited to:

```golang
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
```

Log to stdout, or stderr using a logger instantiated like so:

```golang
outLogger := distillog.NewStdoutLogger("test")

errLogger := distillog.NewStderrLogger("test")

sysLogger := distillog.NewSyslogLogger("test")
```

Alternatively, you can use the package for your logging needs:

```golang
import log "github.com/amoghe/distillog"

// ... later ...

log.Infoln("Starting program")
log.Debugln("initializing the frobnicator")
log.Warningln("frobnicator failure detected, proceeding anyways...")
log.Infoln("Exiting")
```

If you have a file you wish to log to, you should open the file and instantiate
a logger using the file handle, like so:

```golang
if fileHandle, err := ioutil.Tempfile("/tmp", "distillog-test"); err == nil {
        fileLogger := distillog.NewStreamLogger("test", fileHandle)
}
```

If you need a logger that manages the rotation of its files, use `lumberjack`,
like so:

```golang
lumberjackHandle := &lumberjack.Logger{
        Filename:   "/var/log/myapp/foo.log",
        MaxSize:    500,                       // megabytes
        MaxBackups: 3,
        MaxAge:     28,                        // days
}

logger := distillog.NewStreamLogger("tag", lumberjackHandle)

// Alternatively, configure the pkg level logger to emit here

distillog.SetOutput(lumberjackHandle)
```

Once instantiated, you can log messages, like so:

```golang
var := "World!"
myLogger.Infof("Hello, %s", var)
myLogger.Warningln("Goodbye, cruel world!")

```

# Contributing

1. Create an issue, describe the bugfix/feature you wish to implement.
2. Fork the repository
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request

# License

See [LICENSE.txt](LICENSE.txt)

[0]: https://golang.org/pkg/log/
[1]: https://github.com/golang/glog
[2]: https://github.com/natefinch/lumberjack
