package main

import (
	"log"
	"net/http"

	"github.com/dip-dev/go-tutorial/internal/chapter1"
	"github.com/dip-dev/go-tutorial/internal/chapter2"
)

func main() {
	mux := http.NewServeMux()

	// EchoAPI
	mux.HandleFunc("/echo", chapter1.GetEcho)

	mux.HandleFunc("/users/get", chapter2.Get)
	mux.HandleFunc("/users/post", chapter2.Post)

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatalf("failed to launch service: %+v", err)
	}
}
