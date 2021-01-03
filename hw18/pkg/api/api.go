package wsapi

import (
	"log"
	"net/http"

	wsserver "hw18/pkg/ws-server"

	"github.com/gorilla/websocket"
)

type Wsapi struct {
	Srv *wsserver.Server
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (a *Wsapi) send(w http.ResponseWriter, r *http.Request) {
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
			cl := &wsserver.Client{
				Srv:  a.Srv,
				Conn: ws,
			}

			cl.WriteMsg("OK")
			cl.ReadMsg()
			cl.WriteMsg("CONNECTION CLOSED")
			cl.Conn.Close()
			break
		}
	}

}

func (a *Wsapi) messages(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	cl := &wsserver.Client{
		Srv:  a.Srv,
		Conn: ws,
	}

	cl.Conn.SetCloseHandler(func(code int, text string) error {
		cl.Srv.Unregister <- *cl
		return nil
	})

	a.Srv.Clients = append(a.Srv.Clients, cl)
}

func (a *Wsapi) Endpoints() {
	http.HandleFunc("/send", a.send)
	http.HandleFunc("/messages", a.messages)
}
