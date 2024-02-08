package algorithms

import "testing"

func TestBinarySearchFound(t *testing.T) {
	// Binary search only works with sorted lists
	inputArrays := [][]int{
		{1},
		{64, 72},
		{54, 87, 100},
		{7, 9, 12, 15, 53},
		{8, 16, 19, 25, 32, 45, 54, 67, 75, 92, 93, 94, 95},
	}

	searchElements := []int{
		1,
		72,
		54,
		7,
		95,
	}

	for i, arr := range inputArrays {
		if !binarySearch(searchElements[i], arr) {
			t.Fail()
		}
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	// Binary search only works with sorted lists
	inputArrays := [][]int{
		{0},
		{1},
		{64, 72},
		{54, 87, 100},
		{7, 9, 12, 15, 53},
		{8, 16, 19, 25, 32, 45, 54, 67, 75, 92, 93, 94, 95},
	}

	searchElements := []int{
		99,
		2,
		73,
		101,
		16,
		7,
	}

	for i, arr := range inputArrays {
		if binarySearch(searchElements[i], arr) {
			t.Fail()
		}
	}
}
