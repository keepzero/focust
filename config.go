package focust

import (
	"github.com/astaxie/beego/config"
	"strings"
)

var AppConfig config.ConfigContainer

var (
	AppName       string
	AppPath       string
	AppConfigPath string
	WsPort        int
	LogLevel      int
)

// Read config file and init
func ParseConfig() error {

	var err error
	AppConfig, err = config.NewConfig("yaml", AppConfigPath)
	if err != nil {
		return err
	}

	// Websocket Port
	if v, err := AppConfig.Int("wsport"); err == nil {
		WsPort = v
	}

	// LogLevel
	if v := AppConfig.String("level"); v != "" {
		switch strings.ToLower(v) {
		case "debug":
			LogLevel = Debug
		case "info":
			LogLevel = Info
		case "warning":
			LogLevel = Warning
		case "error":
			LogLevel = Error
		case "critical":
			LogLevel = Critical
		default:
			LogLevel = Trace
		}
		Logger.SetLevel(LogLevel)
	}

	return nil
}
