package game

import (
	"RPS_gaming/game/constants"
	"RPS_gaming/game/logic"
	"RPS_gaming/model"
	"RPS_gaming/moves"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Game struct{}

func NewGame() *Game {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Game{}
}

func (g *Game) PlayerVsComputer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	option := r.URL.Query().Get("option")
	index, ok := logic.PossibleMoves[option]
	if !ok {
		res := "Unkown option, possible options are:"
		for option := range logic.PossibleMoves {
			res += fmt.Sprintf(" %s", option)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res))
		return
	}
	p1 := moves.NewWithIndex(index)
	p2 := moves.NewAleatory()
	result := p1.Compare(p2)
	winner := winnerFromResult(result)
	response := model.GameResult{
		P1Move: logic.MovesNames[p1.Index()],
		P2Move: logic.MovesNames[p2.Index()],
		Winner: winner,
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
	result := c1.Compare(c2)
	winner := winnerFromResult(result)
	response := model.GameResult{
		P1Move: logic.MovesNames[c1.Index()],
		P2Move: logic.MovesNames[c2.Index()],
		Winner: winner,
	}
	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"err":"%s"}`, err.Error())))
		return
	}
	w.Write(data)
}

func winnerFromResult(result model.Result) string {
	var winner string
	if result == model.Win {
		winner = constants.PLAYER_1_WINS
	}
	if result == model.Loss {
		winner = constants.PLAYER_2_WINS
	}
	return winner
}
