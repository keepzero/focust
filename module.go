package focust

import (
	"code.google.com/p/go.net/websocket"
	"errors"
	"fmt"
)

type Module struct {
	Handlers map[interface{}]func(*websocket.Conn, interface{}) (string, error)
	Parser   ParserInterface
}

type ModuleInterface interface {
	setHandlers(map[interface{}]func(*websocket.Conn, interface{}) (string, error))
	GetHandlers() map[interface{}]func(*websocket.Conn, interface{}) (string, error)
	Handler(*websocket.Conn)

	Init()
}

func (m *Module) Handler(ws *websocket.Conn) {

	m.Parser.Parse(ws)
	for {
		var content string
		if err := websocket.Message.Receive(ws, &content); err != nil {
			Logger.Error("Can't receive message. %v", err)
			return
		}

		index, request, err := m.Parser.ParseContent(content)
		if err != nil {
			Logger.Error("Parse client request error. %v", err)
			return
		}

		response, err := m.getHandler(index)(ws, request)
		if err != nil {
			Logger.Error("Exec command:%v error. %v", index, err)
		}

		if err := websocket.Message.Send(ws, response); err != nil {
			Logger.Error("Can't send msg. %v", err)
			return
		}
	}
}

func (m *Module) getHandler(index interface{}) func(*websocket.Conn, interface{}) (string, error) {
	if fun, ok := m.Handlers[index]; ok {
		return fun
	} else {
		return func(*websocket.Conn, interface{}) (string, error) {
			return "", errors.New(fmt.Sprintf("No handler map to command:%v", index))
		}
	}
}

func (m *Module) setHandlers(hs map[interface{}]func(*websocket.Conn, interface{}) (string, error)) {
	if m.Handlers == nil {
		m.Handlers = make(map[interface{}]func(*websocket.Conn, interface{}) (string, error))
	}
	for k, v := range hs {
		m.Handlers[k] = v
	}
}

func (m *Module) SetParser(parser ParserInterface) {
	m.Parser = parser
}

func (m *Module) Init() {
	// init module
}
