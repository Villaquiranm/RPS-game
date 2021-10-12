package moves

import (
	"RPS_gaming/server/game/logic"
	"testing"
)

func TestAleatory(t *testing.T) {
	for i := 0; i < 20; i++ {
		aleatoryIndex := NewAleatory().index
		_, ok := logic.ResultsTable[aleatoryIndex]
		if !ok {
			t.Errorf("Aleatory generated a non existent index: %d", aleatoryIndex)
		}
	}

}
