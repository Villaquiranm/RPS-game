package main

import (
	"RPS_gaming/game"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	game := game.NewGame()
	router.HandleFunc("/play", game.PlayerVsComputer)
	router.HandleFunc("/watch", game.ComputerVsComputer)

	log.Fatal(http.ListenAndServe(":3000", router))
}
