package column_1

import (
	"testing"
	"reflect"
	"fmt"
)

var dataSet Sortable = []int{5, 10, 3, 15, 11}

func TestBucket(t *testing.T) {
	var expect Sortable = []int{3, 5, 10, 11, 15}
	dataSet.ExSort()
	fmt.Println(expect)
	fmt.Println(dataSet)

	// expectはSortableのinterfaceとっていないとエラー
	if !reflect.DeepEqual(expect, dataSet) {
		t.Errorf("%s != %s", expect, dataSet)
	}
}