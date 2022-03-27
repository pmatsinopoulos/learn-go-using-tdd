package main

import (
	v1 "github.com/pmatsinopoulos/players/v1"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Problem opening file: %v, Error: %v", dbFileName, err)
	}

	playerStore, err := v1.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("Problem creating file system player store. Error: %v", err)
	}

	server := v1.NewPlayerServer(playerStore)
	defer db.Close()

	log.Fatal(http.ListenAndServe(":5000", server))
}
