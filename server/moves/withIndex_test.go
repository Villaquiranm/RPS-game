package moves

import (
	"testing"
)

func TestWithIndex(t *testing.T) {
	expectedIndex := 10
	moveIndex := NewWithIndex(expectedIndex).Index()
	if moveIndex != expectedIndex {
		t.Errorf("index: %d is different than expected %d", moveIndex, expectedIndex)
	}

}
