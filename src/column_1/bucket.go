package column_1

import (
	"sort"
)


type Sortable []int

func (v Sortable) ExSort() {
	sort.Ints(v)
}






