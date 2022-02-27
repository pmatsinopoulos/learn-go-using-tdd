package main

import (
	v1 "github.com/pmatsinopoulos/players/v1"
	"log"
	"net/http"
)

func main() {
	playerStore := v1.NewInMemoryPlayerStore()
	server := v1.NewPlayerServer(playerStore)
	log.Fatal(http.ListenAndServe(":5000", server))
}
