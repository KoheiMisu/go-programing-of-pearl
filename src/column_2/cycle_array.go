package column_2

// Todo できればポインタでやりたい
func CycleSlice(target []interface{}, val int) []interface{} {
	result := append(target[val:], target[:val]...)

	return result
}
