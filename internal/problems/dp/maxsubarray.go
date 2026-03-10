package dp

import "fmt"

func MaxSubarray(arr []int32) []int32 {
	// Write your code here

	maxSubSeqSum := arr[0]
	for _, val := range arr[1:] {
		if val > maxSubSeqSum {
			maxSubSeqSum = val
		} else if val > 0 {
			maxSubSeqSum += val
		}
	}

	return []int32{maxSubSeqSum, subhelper(arr)}
}

// Kadane's algorithm: O(n) time, O(1) space.
// This is a classic algorithm for finding the maximum sum of a contiguous subarray.
// The idea is to iterate through the array, keeping track of the current sum of the subarray and the maximum sum found so far.
func subhelper(arr []int32) int32 {
	currentSubArray := arr[:1] // first element only
	currentSum := arr[0]
	maxSum := arr[0]

	extend := func(val int32) {
		currentSubArray = append(currentSubArray, val)
		currentSum += val
	}

	startover := func(val int32) {
		currentSubArray = []int32{val}
		currentSum = val
	}

	// There are four scenarios:
	// 1. The current subarray sum is positive, and the current value is positive. Action: Extend the subarray.
	// 2. The current subarray is negative, and the current value is positive. Action: Start a new subarray.
	// 3. The current subarray is positive, and the current value is negative. Action: Extend the subarray.
	// 4. The current subarray is negative, and the current value is negative. Action: Start a new subarray.
	for _, val := range arr[1:] {
		if currentSum >= 0 {
			extend(val)
		} else {
			startover(val)
		}
		if currentSum > maxSum {
			fmt.Printf("New max sum: %d, subarray: %v\n", currentSum, currentSubArray)
			maxSum = currentSum
		}
	}

	return maxSum
}
