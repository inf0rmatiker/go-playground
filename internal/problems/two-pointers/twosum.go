package twopointers

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

// Naive implementation of TwoSum (unsorted variant). O(n^2) time, O(1) space.
func twoSumNaive(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Optimized implementation of TwoSum (unsorted variant): O(n) time, O(n) space.
func twoSumOptimized(nums []int, target int) []int {
	// Remember the number's we've seen and their indices.
	seen := make(map[int]int) // number -> index
	for i, num := range nums {
		complement := target - num
		// If we've seen the complement before, we have our answer.
		if idx, found := seen[complement]; found {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that
they add up to a specific target number. Let these two numbers be numbers[index1] and
numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers index1 and index2, each incremented by one, as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
*/

// Two pointers implementation of TwoSum (sorted variant). O(n) time, O(1) space.
// We could also use the twoPointersNaive approach because that satisfies the O(1) space requirement,
// but the twoPointersOptimized approach is more efficient.
func twoSumSorted(numbers []int, target int) []int {

	// Initialize two pointers at the start and end of the array.
	left, right := 0, len(numbers)-1

	// Idea: Converge the pointers towards the solution.
	// If the sum is two small, need to increase it by moving the left pointer right.
	// If the sum is too large, we need to decrease it by moving the right pointer left.
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1} // Return 1-indexed positions.
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
