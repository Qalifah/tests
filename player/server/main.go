package main

import (
	"log"
	"net/http"
	"github.com/Qalifah/tests/player"
)



func main() {
	handler := &player.Server{Store : player.NewInMemoryStore()}
	log.Fatal(http.ListenAndServe(":5000", handler))
}