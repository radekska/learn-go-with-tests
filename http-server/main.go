package main

import (
	server2 "learn-go-with-tests/http-server/server"
	"learn-go-with-tests/http-server/stores"
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
	store, err := stores.NewFileSystemPlayerStore(jsonDB)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := server2.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":8080", server))
}
