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

	playerStore := v1.FileSystemPlayerStore{Database: db}

	server := v1.NewPlayerServer(playerStore)
	log.Fatal(http.ListenAndServe(":5000", server))
}
