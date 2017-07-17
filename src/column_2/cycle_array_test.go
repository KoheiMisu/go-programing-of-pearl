package column_2

import (
	"testing"
)

func TestCycleSlice(t *testing.T) {
	dataSet_2 := []interface{}{"first", "second", 3, "fourth", 5, 6, "seven"}
	result := CycleSlice(dataSet_2, 2)

	expected := []interface{}{3, "fourth", 5, 6, "seven", "first", "second"}

	for i, k := range result {
		if expected[i] != k {
			t.Errorf("%d != %d", expected[i], k)
		}
	}
}





