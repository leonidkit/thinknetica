package wsapi

import (
	"log"
	"net/http"

	wshub "hw18/pkg/ws-hub"

	"github.com/gorilla/websocket"
)

type Wsapi struct {
	Hub *wshub.Hub
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
			cl := &wshub.Client{
				Hub:  a.Hub,
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

	cl := &wshub.Client{
		Hub:  a.Hub,
		Conn: ws,
	}

	cl.Conn.SetCloseHandler(func(code int, text string) error {
		cl.Hub.Unregister <- *cl
		return nil
	})

	a.Hub.Clients = append(a.Hub.Clients, cl)
}

func (a *Wsapi) Endpoints() {
	http.HandleFunc("/send", a.send)
	http.HandleFunc("/messages", a.messages)
}
