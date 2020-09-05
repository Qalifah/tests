package main

import (
	"log"
	"net/http"
	"github.com/Qalifah/tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

    if err != nil {
        log.Fatal(err)
    }
    defer close()

    server := poker.NewServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}