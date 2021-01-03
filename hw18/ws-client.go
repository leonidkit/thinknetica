package main

import (
	"bytes"
	"log"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type client struct {
	srv  *Server
	conn *websocket.Conn
}

func (c *client) readMsg() {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error client read: %v", err)
		}
		return
	}
	message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

	c.srv.Broadcast <- string(message)
}

func (c *client) writeMsg(msg string) {
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Printf("error client write: %v", err)
	}
}
