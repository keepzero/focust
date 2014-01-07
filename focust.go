package focust

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"os"
	"path"
	"strconv"
)

const VERSION = "0.0.1"

var Modules map[string]ModuleInterface

func init() {
	// init focust core
	Modules = make(map[string]ModuleInterface)
	Logger = logs.NewLogger(1000)
	Logger.SetLogger("console", "")

	// default config value
	AppName = "focust"
	AppPath = path.Dir(os.Args[0])
	AppConfigPath = path.Join(AppPath, "conf", "app.yaml")
	WsPort = 8080
	LogLevel = Trace
	Logger.SetLevel(LogLevel)

	// init others in ParseConfig func
	if err := ParseConfig(); err != nil {
		Log(Error, "Parse config error. %v", err)
	}
}

func Serve(path string, module ModuleInterface) {
	if _, ok := Modules[path]; ok {
		panic(fmt.Sprintf("Aleady exist module to serve path:\"%s\"", path))
	}
	Modules[path] = module
}

func ServeHTTP(path string, handler func(http.ResponseWriter, *http.Request)) {
	if _, ok := Modules[path]; ok {
		panic(fmt.Sprintf("Aleady exist module to serve path:\"%s\"", path))
	}
	http.HandleFunc(path, handler)
}

func Run() {
	for path, module := range Modules {
		module.Init()
		module.setHandlers(module.GetHandlers())
		module.setDefaultHandler(module.Error404)
		http.Handle(path, websocket.Server{Handler: module.Handler})
	}

	if err := http.ListenAndServe(":"+strconv.Itoa(WsPort), nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
