package wshub

import (
	"bytes"
	"log"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
}

func (c *Client) ReadMsg() {
	_, message, err := c.Conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error client read: %v", err)
		}
		return
	}
	message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

	c.Hub.Broadcast <- string(message)
}

func (c *Client) WriteMsg(msg string) {
	err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Printf("error client write: %v", err)
	}
}
