package util

import (
	"testing"
)

type testSet struct {
	target interface{}
	arr []interface{}
	expect bool
}

var testSets = []testSet{
	{"test", []interface{}{"_test_", "test"}, true},
	{1, []interface{}{3, 2, 1}, true},
	{1, []interface{}{"1", "2", "3"}, false},
	{1.1, []interface{}{3.3, 1.1}, true},
}

func TestInArray(t *testing.T)  {
	for i, data := range testSets {
		if InArray(data.target, data.arr) != data.expect {
			t.Errorf("failed index: %v", i)
		}
	}
}
