package main

import (
	"RPS_gaming/game"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	router := mux.NewRouter()
	game := game.NewGame()
	router.HandleFunc("/play", game.PlayerVsComputer)
	router.HandleFunc("/watch", game.ComputerVsComputer)

	log.Fatal(http.ListenAndServe(":3000", router))
}
