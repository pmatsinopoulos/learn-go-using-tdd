package main

import (
	"log"
	"net/http"
	"os"
)

func main__() {
	Greet(os.Stdout, "Elodie\n")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
