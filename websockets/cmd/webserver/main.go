package main

import (
	"log"
	"net/http"

	poker "github.com/rinwaowuogba/learn-go-with-tests/websockets"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewTexasHoldemGame(poker.BlindAlerterFunc(poker.Alerter), store)
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("problem creating new player server %v", err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}
