package main

import (
	"hw21/pkg/api"
	"hw21/pkg/storage/postgres"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	strg, err := postgres.New("127.0.0.1", "5432", "admin", "admin", "store", false)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	defer strg.Close()

	api := api.New(router, strg)

	log.Fatal(http.ListenAndServe(":8000", api.Router))
}
