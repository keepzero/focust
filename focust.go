package focust

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
)

const VERSION = "0.0.1"

func init() {
	fmt.Println("Init focust")
}

func Router(path string, ids [2]int, module *ModuleInterface) error {
	return nil
}

func Run() {
	fmt.Println("Running...")
	// run websocket here
	http.Handle("/", websocket.Server{Handler: handler})
}

func handler(ws *websocket.Conn) {

	for {

		break
	}
}
