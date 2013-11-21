package focust

import (
	"code.google.com/p/go.net/websocket"
	"errors"
	"fmt"
)

type Module struct {
	Name     string
	Handlers map[interface{}]func(unique interface{}, requestStr string) (string, error)
	Parser   ParserInterface
}

type ModuleInterface interface {
	GetName() string
	getHandler(interface{}) func(interface{}, string) (string, error)
	Handler(*websocket.Conn)
	Init()
}

func (m *Module) GetName() string {
	return m.Name
}

func (m *Module) Handler(ws *websocket.Conn) {

	for {
		//index, request := m.parser.Parse(ws)
		//_, _ = index, request
		//handle := m.getHandler(index)
		//handle(index, request)

		//break
	}
}

func (m *Module) getHandler(index interface{}) func(interface{}, string) (string, error) {
	if fun, ok := m.Handlers[index]; ok {
		return fun
	} else {
		return func(interface{}, string) (string, error) {
			return "", errors.New(fmt.Sprintf("No handler map to command:%v", index))
		}
	}
}

func (m *Module) Init() {
	fmt.Println("Init module...")
}
