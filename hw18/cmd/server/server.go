package main

import (
	api "hw18/pkg/api"
	wsserver "hw18/pkg/ws-server"
	"log"
	"net/http"
)

func main() {
	wa := &api.Wsapi{
		Srv: &wsserver.Server{
			Broadcast: make(chan string),
		},
	}
	wa.Endpoints()
	wa.Srv.Run()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
