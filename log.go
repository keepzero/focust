package focust

import (
	"github.com/astaxie/beego/logs"
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
	Logger.SetLevel(l)
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
