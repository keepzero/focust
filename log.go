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

func init() {
	Logger = logs.NewLogger(100)
	Logger.SetLogger("console", "")
}
