package twopointers

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func generateSumNumbers(a, b, count int) ([]int, int, int) {
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = a + b + 1
	}
	mid := count / 2
	nums[mid] = a
	nums[mid+1] = b
	return nums, mid, mid + 1
}

func TestTwoSum(t *testing.T) {
	// Default test case
	inputNums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}

	result := twoSumNaive(inputNums, target)
	assert.NotNil(t, result)
	assert.Equal(t, expected, result)

	// Generate a difficult test case
	a, b := 12345, 67890
	nums, idx1, idx2 := generateSumNumbers(a, b, 100000)
	target = a + b
	expected = []int{idx1, idx2}

	// Test twoSumNaive
	start := time.Now()
	result = twoSumNaive(nums, target)
	fmt.Printf("Time for naive two sum: %s\n", time.Since(start).Truncate(10*time.Millisecond).String())
	assert.NotNil(t, result)
	assert.Equal(t, expected, result)

	// Test twoSumOptimized
	start = time.Now()
	result = twoSumOptimized(nums, target)
	fmt.Printf("Time for optimized two sum: %s\n", time.Since(start).Truncate(10*time.Millisecond).String())
	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}
