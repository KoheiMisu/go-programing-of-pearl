package column_2

import (
	"fmt"
	"strings"
)

func GenerateAnagram(word string) (results []string) {
	splitted := strings.Split(word,"")

	//func setAnagram() {
	//
	//}
	//
	//setAnagram()

	for i, val := range splitted {
		index := len(splitted) - 1 - i

		// 先頭の一文字
		lead := splitted[index]
		fmt.Println(i, val)
		fmt.Println(lead)
		results = append(results, lead)
	}

	return
}