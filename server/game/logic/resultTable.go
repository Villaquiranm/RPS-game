package logic

import (
	"RPS_gaming/server/model"
)

const (
	ROCK int = iota
	PAPER
	SCISSORS
)

//Return the index of every posible move
var PossibleMoves = map[string]int{
	"rock":     ROCK,
	"paper":    PAPER,
	"scissors": SCISSORS,
}

var MovesNames = map[int]string{
	ROCK:     "rock",
	PAPER:    "paper",
	SCISSORS: "scissors",
}

// Results table will have the Result of every move index each possible move have an index
// this variable will have map[int] inside another map[int]model.Result
// The litlest index correspond to the first map key and the biggest to the other one

var ResultsTable map[int]map[int]model.Result = map[int]map[int]model.Result{
	ROCK:     rockTable,
	PAPER:    paperTable,
	SCISSORS: scissorsTable,
}

//Rock table
var rockTable map[int]model.Result = map[int]model.Result{
	ROCK:     model.Draw,
	SCISSORS: model.Win,
	PAPER:    model.Loss,
}

//Paper table
var paperTable map[int]model.Result = map[int]model.Result{
	SCISSORS: model.Loss,
	PAPER:    model.Draw,
	ROCK:     model.Win,
}

//Scissors table
var scissorsTable map[int]model.Result = map[int]model.Result{
	SCISSORS: model.Draw,
	PAPER:    model.Win,
	ROCK:     model.Loss,
}
