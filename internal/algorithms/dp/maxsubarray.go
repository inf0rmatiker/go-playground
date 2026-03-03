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
