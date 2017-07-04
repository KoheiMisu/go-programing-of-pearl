package util

func Max(target []int) int {
	var result int

	for _, val := range target {
		if result < val {
			result = val
		}
	}

	return result
}

func Min(target []int) int {
	var result int

	for _, val := range target {
		if result > val {
			result = val
		}
	}

	return result
}
