package focust

import "code.google.com/p/go.net/websocket"

type Parser struct{}

type ParserInterface interface {
	Parse(*websocket.Conn) (interface{}, interface{}, error)
	ParseContent(string) (interface{}, interface{}, error)
}

func (p *Parser) Parse(ws *websocket.Conn) (interface{}, string, error) {
	// default return nil
	return nil, "", nil
}

func (p *Parser) ParseContent(content string) (interface{}, interface{}, error) {
	// default return nil and content
	return nil, content, nil
}
