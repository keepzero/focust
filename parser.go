package focust

import "code.google.com/p/go.net/websocket"

type Parser struct {
	Name string
}

type ParserInterface interface {
	GetName() string
	Parse(*websocket.Conn) (interface{}, string)
	ParseContent(string) (interface{}, string)
}

func (p *Parser) Parse(ws *websocket.Conn) (interface{}, string) {
	// default return nil
	return nil, ""
}

func (p *Parser) ParseContent(content string) (interface{}, string) {
	// default return nil and content
	return nil, content
}
