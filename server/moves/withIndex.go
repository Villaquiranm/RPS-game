package moves

type WithIndex struct {
	index int
}

func NewWithIndex(index int) *WithIndex {
	return &WithIndex{index: index}
}

func (p *WithIndex) Index() int {
	return p.index
}
