package focust

import (
	"github.com/astaxie/beego/logs"
)

// Log levels to control the logging output.
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

// SetLogLevel sets the global log level used by the simple
// logger.
func SetLevel(l int) {
	Logger.SetLevel(l)
}

// logger references the used application logger.
var Logger *logs.BeeLogger

func Log(level int, format string, v ...interface{}) {
	if level == LevelTrace {
		Logger.Trace(format, v...)
	} else if level == LevelDebug {
		Logger.Info(format, v...)
	} else if level == LevelInfo {
		Logger.Info(format, v...)
	} else if level == LevelWarning {
		Logger.Warn(format, v...)
	} else if level == LevelError {
		Logger.Error(format, v...)
	} else if level == LevelCritical {
		Logger.Critical(format, v...)
	}
}
