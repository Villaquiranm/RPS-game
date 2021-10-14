package moves

import (
	"RPS_gaming/game/logic"
	"RPS_gaming/model"
	"math/rand"
)

// Move represents one of the options to chose on the game
type Move struct {
	index int
}

func NewWithIndex(index int) *Move {
	return &Move{index: index}
}

func NewAleatory() *Move {
	// generate an aleatory index 0 -> len possible moves
	index := rand.Intn(len(logic.PossibleMoves))
	return &Move{index: index}
}

func (m *Move) Index() int {
	return m.index
}

/* Compare Using the clockwise & counter-clockwise logic.
   0   Draw
   1-2 Loss
   3-4 Win
*/
func (m *Move) Compare(other *Move) model.Result {
	if m.index == other.index {
		return model.Draw
	}
	lenMoves := len(logic.PossibleMoves)
	foundIndex := 0
	for i := 1; i < lenMoves; i++ {
		currentIndex := (m.index + i) % lenMoves
		if currentIndex == other.index {
			foundIndex = i
		}
	}
	if foundIndex <= 2 {
		return model.Loss
	}
	return model.Win
}
