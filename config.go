package focust

import (
	//"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"os"
	"path"
)

var (
	AppName       string
	AppPath       string
	AppConfigPath string

	WsPort   int
	LogLevel int
)

func init() {
	// default value
	AppName = "focust"
	AppPath = path.Dir(os.Args[0])
	AppConfigPath = path.Join(AppPath, "conf", "app.yaml")
}

func ParseConfig() error {

	// read from config file
	WsPort = 8080
	LogLevel = LevelTrace

	// Init
	Logger = logs.NewLogger(100)
	Logger.SetLogger("console", "")
	Logger.SetLevel(LogLevel)

	return nil
}
