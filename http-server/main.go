package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":8080", server))
}

// TODO
// - Pick a store Postgres DB store
// - Make PostgresPlayerStore implement PlayerStore
// - TDD the functionality so you're sure it works
// - Plug it into the integration test, check it's still ok
// - Finally plug it into main
