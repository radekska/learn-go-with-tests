package main

import (
	"learn-go-with-tests/http-server"
	"log"
	"net/http"
	"os"
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
	jsonDB, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, err := http_server.NewFileSystemPlayerStore(jsonDB)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := http_server.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":8080", server))
}
