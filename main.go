package main

import (
	"net/http"
)

func main() {
	multiplexer := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: multiplexer,
	}
	server.ListenAndServe()
}
