package game

import (
	"RPS_gaming/server/game/logic"
	"RPS_gaming/server/model"
	"RPS_gaming/server/moves"
	"fmt"
	"net/http"
)

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) compare(p1 moves.Move, p2 moves.Move) model.Result {
	return logic.ResultsTable[p1.Index()][p2.Index()]
}

func (g *Game) PlayerVsComputer(w http.ResponseWriter, r *http.Request) {
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
	result := g.compare(p1, moves.NewAleatory())
	if result == model.Loss {
		w.Write([]byte("You loss"))
		w.WriteHeader(http.StatusOK)
	}
	if result == model.Win {
		w.Write([]byte("You WIN"))
		w.WriteHeader(http.StatusOK)
	}
	if result == model.Draw {
		w.Write([]byte("a draw"))
		w.WriteHeader(http.StatusOK)
	}
}

func (g *Game) ComputerVsComputer(w http.ResponseWriter, r *http.Request) {
	c1 := moves.NewAleatory()
	c2 := moves.NewAleatory()
	result := g.compare(c1, c2)
	if result == model.Loss {
		w.Write([]byte("Computer 2 wins"))
		w.WriteHeader(http.StatusOK)
	}
	if result == model.Win {
		w.Write([]byte("Computer 1 wins"))
		w.WriteHeader(http.StatusOK)
	}
	if result == model.Draw {
		w.Write([]byte("a draw"))
		w.WriteHeader(http.StatusOK)
	}
}
