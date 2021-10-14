package moves

import (
	"RPS_gaming/game/logic"
	"RPS_gaming/model"
	"testing"
)

func TestCompare(t *testing.T) {
	tests := map[string]struct {
		index1         int
		index2         int
		expectedResult model.Result
	}{
		"p1 wins": {
			index1:         logic.PAPER,
			index2:         logic.ROCK,
			expectedResult: model.Win,
		},
		"p2 wins": {
			index1:         logic.PAPER,
			index2:         logic.SCISSORS,
			expectedResult: model.Loss,
		},
		"draw": {
			index1:         logic.SCISSORS,
			index2:         logic.SCISSORS,
			expectedResult: model.Draw,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m1 := NewWithIndex(tc.index1)
			m2 := NewWithIndex(tc.index2)
			result := m1.Compare(m2)
			if result != tc.expectedResult {
				t.Errorf("Unexpected result: got :%d, \n wanted: %d", result, tc.expectedResult)
			}
		})
	}
}
