package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	// Default test case
	inputNums := []int{0, 0, 0, 0}
	expected := [][]int{{0, 0, 0}}

	result := threeSum(inputNums)
	assert.NotNil(t, result)
	assert.Equal(t, expected, result)

	// Test twoSumNaive
	result = threeSum(inputNums)
	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}
