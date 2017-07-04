package column_2

/**
 int型配列に対する二分探索の実装
 真ん中より大きい整数と
 小さい整数の二つのグループにわける
 数字が格納されている数の少ないグループから
 指定された範囲にない数字を出す
 max 15
 range [1, 4, 11, 7, 9, 10, 8, 12]
 */
func BinarySort(target []int, max int) int {
	// 少数は切り捨て ex.) 15 / 2 => 7
	middle := max / 2
	minGroup := []int{}
	maxGroup := []int{}

	/**
	 二つのグループに分ける
	 */
	for _, val := range target {
		if val >= middle {
			maxGroup = append(maxGroup, val)
		} else {
			minGroup = append(minGroup, val)
		}
	}

	/**
	 個数の少ないグループに対して処理を行う
	 */
	if len(maxGroup) > len(minGroup) {
		i := 1
		for _, val := range minGroup {
			if i != val {
				return i
			}
			i++
		}
	} else {
		i := middle
		for _, val := range maxGroup {
			if i != val {
				return i
			}
			i++
		}
	}

	return 0
}
