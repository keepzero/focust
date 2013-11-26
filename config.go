package focust

import (
	"github.com/astaxie/beego/config"
)

var (
	AppName       string
	AppPath       string
	AppConfigPath string
	WsPort        int
	LogLevel      int
)

// Read config file and init
func ParseConfig() error {

	AppConfig, err := config.NewConfig("yaml", AppConfigPath)
	if err != nil {
		return err
	}

	if v, err := AppConfig.Int("wsport"); err == nil {
		WsPort = v
	}

	//LogLevel = LevelTrace
	//Logger.SetLevel(LogLevel)

	return nil
}
