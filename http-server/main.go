package main

import (
	"log"
	"net/http"
)

func main() {
	db, err := GetDatabase()
	if err != nil {
		log.Fatal(err)
	}
	store := NewPostgresPlayerStore(db)
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":8080", server))
}
