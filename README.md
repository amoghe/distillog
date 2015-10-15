# What is `distillog`?

`distillog` is an _opinionated_ golang package that aims to provide the most
minimal logging interface (and implementation) that allow a program to emit log
messages. It is logging, _distilled_.

# Yet _another_ logging library for go(lang)?

> Logging libraries are like opinions, everyone seems to have one -- Anon(?)

Most other logging libraries do either too little ([stdlib](https://golang.org/pkg/log/))
or too much ([glog](https://github.com/golang/glog)).

## Too _little_, you say?

Presenting varying levels of verbosity (or severity) are an important part of
what makes a program more usable or debuggable. The stdlib has an approach that is
too spartan (exposing only `Println` and friends) to be used directly in programs
that wish to offer better control over their log output. This is a sufficient
mechanism for small programs, but starts to feel insufficient very quickly.

## Too _much_, you say?

Ever used `log.Panicf` or `log.Fatalf`? Exiting your program is *not* something
your log library should be doing! Similarly, other libraries offer options for
maintaining old log files and rotating them. Your logging library should not need
to care about this. Whatever facility (other libraries call this a "backend")
messages are sent to should determine how old messages are handled.

## But log file rotation is absolutely necessary!

Agreed, and someone's gotta do it, but it need not be your logging library!

You can use `distillog` along with a [lumberjack](https://github.com/natefinch/lumberjack) "backend". It provides an `io.WriteCloser` which performs all the magic you need. Initialize a logger
using `distillog.NewStream`, pass it an instance of the `io.WriteCloser`
that lumberjack returns, _et voila_, you have a logger that does what you need.

## And how is `distillog` different?

`distillog` restricts itself to:
- presenting an interface so that you can swap different loggers.
- providing logger implementations for logging to the most common backends
	- streams (`stderr`, or files, via `io.WriteCloser`)
	- syslog
- avoid taking on any non-essential responsibilities.

By using an interface, you can write programs that aren't married to a particular
logging system. More importantly, you can switch between logging to stderr and
syslog by simply instantiating an appropriate logger.

# Usage/examples:

As seen in the godoc, the interface is limited to:

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

If you have a file you wish to log to, you should open the file and instantiate a logger
using the file handle, like so:

```golang
if fileHandle, err := ioutil.Tempfile("/tmp", "distillog-test"); err == nil {
        fileLogger := distillog.NewStreamLogger("test", fileHandle)
}
```

If you need a logger that manages the rotation of its own files, use `lumberjack`, like so:

```golang
lumberjackHandle := &lumberjack.Logger{
        Filename:   "/var/log/myapp/foo.log",
        MaxSize:    500,                       // megabytes
        MaxBackups: 3,
        MaxAge:     28,                        // days
}

logger := distillog.NewStreamLogger("tag", lumberjackHandle)
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
