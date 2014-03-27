package focust

import (
	"code.google.com/p/go.net/websocket"
	//"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"os"
	"path"
	"strconv"
)

const VERSION = "0.0.1"

func init() {

	// init focust core
	Logger = logs.NewLogger(1000)
	Logger.SetLogger("console", "")

	// default config value
	AppName = "focust"
	AppPath = path.Dir(os.Args[0])
	AppConfigPath = path.Join(AppPath, "conf", "app.yaml")
	WsPort = 8080
	AppId = 1
	LogLevel = Trace
	Logger.SetLevel(LogLevel)

	// init others in ParseConfig func
	if err := ParseConfig(); err != nil {
		Log(Error, "Parse config error. %s", err.Error())
		panic(err.Error())
	}
}

func Serve(path string, handler func(*websocket.Conn)) {
	http.Handle(path, websocket.Server{Handler: handler})
}

func ServeHTTP(path string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(path, handler)
}

func Run() {
	if err := http.ListenAndServe(":"+strconv.Itoa(WsPort), nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
