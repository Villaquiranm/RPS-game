package game

import (
	"RPS_gaming/server/game/logic"
	"RPS_gaming/server/model"
	"RPS_gaming/server/moves"
	"encoding/json"
	"fmt"
	"net/http"
)

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) compare(p1 moves.Move, p2 moves.Move) string {
	result := logic.ResultsTable[p1.Index()][p2.Index()]
	switch result {
	case model.Loss:
		return "Player 2"
	case model.Win:
		return "Player 1"
	default:
		return ""
	}
}

func (g *Game) PlayerVsComputer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	option := r.URL.Query().Get("option")
	index, ok := logic.PossibleMoves[option]
	if !ok {
		res := "Unkown option, possible options are:"
		for option, _ := range logic.PossibleMoves {
			res += fmt.Sprintf(" %s", option)
		}
		w.Write([]byte(res))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p1 := moves.NewWithIndex(index)
	p2 := moves.NewAleatory()
	response := model.GameResult{
		P1Move: logic.MovesNames[p1.Index()],
		P2Move: logic.MovesNames[p2.Index()],
		Winner: g.compare(p1, p2),
	}
	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"err":"%s"}`, err.Error())))
		return
	}
	w.Write(data)
}

func (g *Game) ComputerVsComputer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c1 := moves.NewAleatory()
	c2 := moves.NewAleatory()
	response := model.GameResult{
		P1Move: logic.MovesNames[c1.Index()],
		P2Move: logic.MovesNames[c2.Index()],
		Winner: g.compare(c1, c2),
	}
	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"err":"%s"}`, err.Error())))
		return
	}
	w.Write(data)
}
