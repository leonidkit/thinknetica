package main

import (
	api "hw18/pkg/api"
	wshub "hw18/pkg/ws-hub"
	"log"
	"net/http"
)

func main() {
	wa := &api.Wsapi{
		Hub: &wshub.Hub{
			Broadcast: make(chan string),
		},
	}
	wa.Endpoints()
	wa.Hub.Run()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
