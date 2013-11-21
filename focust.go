package focust

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
)

const VERSION = "0.0.1"

var modules map[string]ModuleInterface

func init() {
	modules = make(map[string]ModuleInterface)
}

func Serve(path string, module ModuleInterface) {
	if _, ok := modules[path]; ok {
		panic(fmt.Sprintf("Aleady exist module to serve path:\"%s\"", path))
	}
	module.Init()
	modules[path] = module

	fmt.Printf("Use module:%s to serve path:\"%s\"\n", module.GetName(), path)
}

func Run() {
	for path, module := range modules {
		http.Handle(path, websocket.Server{Handler: module.Handler})
	}

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
