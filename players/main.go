package main

import (
	v1 "github.com/pmatsinopoulos/players/v1"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(v1.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
