package moves

import (
	"RPS_gaming/server/game/logic"
	"math/rand"
)

type Aleatory struct {
	index int
}

func NewAleatory() *Aleatory {
	// generate an aleatory index 0 -> len possible moves
	index := rand.Intn(len(logic.ResultsTable))
	return &Aleatory{index: index}
}

func (p *Aleatory) Index() int {
	return p.index
}
