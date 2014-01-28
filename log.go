package focust

import (
	"github.com/astaxie/beego/logs"
	"strings"
)

// Log levels to control the logging output.
const (
	Trace = iota
	Debug
	Info
	Warning
	Error
	Critical
)

// SetLogLevel sets the global log level used by the simple logger.
func SetLevel(l int) {
	LogLevel = l
	Logger.SetLevel(LogLevel)
}

// logger references the used application logger.
var Logger *logs.BeeLogger

// Used for unify log format
func Log(level int, format string, v ...interface{}) {
	if level == Trace {
		Logger.Trace(format, v...)
	} else if level == Debug {
		Logger.Debug(format, v...)
	} else if level == Info {
		Logger.Info(format, v...)
	} else if level == Warning {
		Logger.Warn(format, v...)
	} else if level == Error {
		Logger.Error(format, v...)
	} else if level == Critical {
		Logger.Critical(format, v...)
	}
}

// Trace logs a message at trace level.
func T(v ...interface{}) {
	Logger.Trace(generateFmtStr(len(v)), v...)
}

// Debug logs a message at debug level.
func D(v ...interface{}) {
	Logger.Debug(generateFmtStr(len(v)), v...)
}

// Info logs a message at info level.
func I(v ...interface{}) {
	Logger.Info(generateFmtStr(len(v)), v...)
}

// Warning logs a message at warning level.
func W(v ...interface{}) {
	Logger.Warn(generateFmtStr(len(v)), v...)
}

// Error logs a message at error level.
func E(v ...interface{}) {
	Logger.Error(generateFmtStr(len(v)), v...)
}

// Critical logs a message at critical level.
func C(v ...interface{}) {
	Logger.Critical(generateFmtStr(len(v)), v...)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
