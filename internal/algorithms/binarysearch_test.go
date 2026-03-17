package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		assert.True(t, binarySearch(searchElements[i], arr))
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
		assert.False(t, binarySearch(searchElements[i], arr))
	}
}

func TestBinarySearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 2, 4, 8},
		{10, 11, 12, 13},
		{14, 20, 30, 40},
	}

	searchElementsTrue := []int{10, 4, 20, 30}
	searchElementsFalse := []int{21, 45, 0, -1}

	for i := range searchElementsTrue {
		assert.True(t, searchMatrix(matrix, searchElementsTrue[i]))
		assert.False(t, searchMatrix(matrix, searchElementsFalse[i]))
	}
}
