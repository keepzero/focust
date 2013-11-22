package focust

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
	"strconv"
)

const VERSION = "0.0.1"

var Modules map[string]ModuleInterface

func init() {
	Modules = make(map[string]ModuleInterface)
	if err := ParseConfig(); err != nil {
		panic(err.Error())
	}
}

func Serve(path string, module ModuleInterface) {
	if _, ok := Modules[path]; ok {
		panic(fmt.Sprintf("Aleady exist module to serve path:\"%s\"", path))
	}
	Modules[path] = module
}

func Run() {
	SetLevel(LogLevel)

	for path, module := range Modules {
		module.Init()
		http.Handle(path, websocket.Server{Handler: module.Handler})
	}

	if err := http.ListenAndServe(":"+strconv.Itoa(WsPort), nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
