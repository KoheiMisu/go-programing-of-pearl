package column_2

import (
	"testing"
)

var dataSet = []int{5, 10, 3, 15, 11, 1, 2, 6, 4}

// minGroup[5, 3, 1, 2, 6, 4]
// maxGroup[10, 15, 11]

func TestBinarySort(t *testing.T) {
	expect := 7
	actual := BinarySort(dataSet, 15)

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}