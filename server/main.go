package main

import (
	"RPS_gaming/server/game"
	"net/http"
)

func main() {
	game := game.NewGame()
	http.HandleFunc("/play", game.PlayerVsComputer)
	http.HandleFunc("/watch", game.ComputerVsComputer)
	http.ListenAndServe(":8080", nil)
}
