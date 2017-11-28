/*
Package distillog provides a minimalistic logging interface (inspired by the
stdlib) that also supports levelled logging.

You can instantiate a logger (that adheres to the Logger interface) or use the
pkg level log functions (like the stdlib).

You can set where the log messages are sent to by either either instantiating
your own logger using the appropriate constructor function, or by using
SetOutput to configure the pkg level logger.

You may also use `lumberjack` or a similar library in conjunction with this
one to manage (rotate) your log files.

Visit the README at https://github.com/amoghe/distillog
*/
package distillog
