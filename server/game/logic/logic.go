package logic

const (
	ROCK int = iota
	SPOCK
	PAPER
	LIZARD
	SCISSORS
)

// PossibleMoves Return the index of every posible move
var PossibleMoves = map[string]int{
	"Rock":     ROCK,
	"Paper":    PAPER,
	"Scissors": SCISSORS,
	"Lizard":   LIZARD,
	"Spock":    SPOCK,
}

// MovesNames Return the name of the move with the index of a move
var MovesNames = map[int]string{
	ROCK:     "Rock",
	PAPER:    "Paper",
	SCISSORS: "Scissors",
	LIZARD:   "Lizard",
	SPOCK:    "Spock",
}
