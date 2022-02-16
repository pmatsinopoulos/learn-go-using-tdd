package main

import (
	v1 "github.com/pmatsinopoulos/players/v1"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
}

func (i InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	playerStore := InMemoryPlayerStore{}
	server := v1.PlayerServer{PlayerStore: playerStore}
	log.Fatal(http.ListenAndServe(":5000", server))
}
