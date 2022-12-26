package main

import (
	db2 "learn-go-with-tests/http-server/db"
	server2 "learn-go-with-tests/http-server/server"
	"learn-go-with-tests/http-server/stores"
	"log"
	"net/http"
)

func main() {
	db, err := db2.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}
	store := stores.NewPostgresPlayerStore(db)
	server := &server2.PlayerServer{Store: store}
	log.Fatal(http.ListenAndServe(":8080", server))
}
