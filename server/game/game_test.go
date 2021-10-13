package game

import (
	"RPS_gaming/game/constants"
	"RPS_gaming/game/logic"
	"RPS_gaming/moves"
	"testing"
)

func TestCompare(t *testing.T) {
	tests := map[string]struct {
		index1         int
		index2         int
		expectedResult string
	}{
		"p1 wins": {
			index1:         logic.PAPER,
			index2:         logic.ROCK,
			expectedResult: constants.PLAYER_1_WINS,
		},
		"p2 wins": {
			index1:         logic.PAPER,
			index2:         logic.SCISSORS,
			expectedResult: constants.PLAYER_2_WINS,
		},
		"draw": {
			index1:         logic.SCISSORS,
			index2:         logic.SCISSORS,
			expectedResult: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			game := NewGame()
			result := game.compare(moves.NewWithIndex(tc.index1), moves.NewWithIndex(tc.index2))
			if result != tc.expectedResult {
				t.Errorf("Unexpected result: got :%s, \n wanted: %s", result, tc.expectedResult)
			}
		})
	}
}
