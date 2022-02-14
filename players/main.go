package main

import (
	v1 "github.com/pmatsinopoulos/players/v1"
	"log"
	"net/http"
)

func main() {
	server := v1.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
