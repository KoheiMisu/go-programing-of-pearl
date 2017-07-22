package column_2

import (
	"testing"
	"reflect"
)

var applyDataSet = []string{"pot", "top"}

var expectedDataSet = []string{"pot", "pto", "otp", "opt", "top", "tpo"}

func TestGenerateAnagram(t *testing.T) {
	word := "top"

	// assert generated anagram result
	generated := GenerateAnagram(word, 0, len(word))

	if !reflect.DeepEqual(generated, expectedDataSet) {
		t.Errorf("diff %v and %v", generated, expectedDataSet)
	}
}