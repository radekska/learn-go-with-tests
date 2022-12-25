package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewPostgresPlayerStore()
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":8080", server))
}
