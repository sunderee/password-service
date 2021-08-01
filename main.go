package main

import (
	"net/http"

	"github.com/sunderee/password-service/server"
)

func main() {
	http.HandleFunc("/generate", server.ServerHandler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
