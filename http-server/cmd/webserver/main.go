package main

import (
	poker "learn-go-with-tests/http-server"
	"log"
	"net/http"
)

func main() {

	// DB store
	//db, err := db2.GetDatabase()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//store := stores.NewPostgresPlayerStore(db)

	// File store
	const dbFileName = "game.db.json"
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()
	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":8080", server))
}
