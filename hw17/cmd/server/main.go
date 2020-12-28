package main

import (
	"net/http"
	"pkg/jwtauth"
)

func main() {
	jwtserver := &jwtauth.AuthServer{}
	jwtserver.Endpoints()
	http.ListenAndServe(":8000", jwtserver.Handler)
}
