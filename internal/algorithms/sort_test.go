package algorithms

import (
	"fmt"
	"testing"
)

func isSorted(actual []int) bool {
	if len(actual) < 2 {
		return true
	}

	prev := actual[0]
	for i := 1; i < len(actual); i++ {
		if actual[i] < prev {
			return false
		}
	}
	return true
}

func TestMergeSort(t *testing.T) {
	input := []int{5, 1, 2, 3, 4}
	output := mergeSort(input)

	fmt.Println(output)
	if !isSorted(output) {
		t.Fail()
	}
}
