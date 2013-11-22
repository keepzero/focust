package focust

import (
	//"github.com/astaxie/beego/config"
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
	AppName = "focust"
	AppPath = path.Dir(os.Args[0])
	AppConfigPath = path.Join(AppPath, "conf", "app.yaml")
}

func ParseConfig() error {
	// read from confi
	WsPort = 8080
	LogLevel = LevelTrace
	return nil
}
