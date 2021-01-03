package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type wsapi struct {
	srv *Server
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (a *wsapi) send(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}

		if string(message) == "password" {
			cl := &client{
				srv:  a.srv,
				conn: ws,
			}

			cl.writeMsg("OK")
			cl.readMsg()
			cl.writeMsg("CONNECTION CLOSED")
			cl.conn.Close()
			break
		}
	}

}

func (a *wsapi) messages(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	cl := &client{
		srv:  a.srv,
		conn: ws,
	}

	cl.conn.SetCloseHandler(func(code int, text string) error {
		cl.srv.Unregister <- *cl
		return nil
	})

	a.srv.Clients = append(a.srv.Clients, cl)
}

func (a *wsapi) endpoints() {
	http.HandleFunc("/send", a.send)
	http.HandleFunc("/messages", a.messages)
}

func main() {
	wa := &wsapi{
		srv: &Server{
			Broadcast: make(chan string),
		},
	}
	wa.endpoints()
	wa.srv.run()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
