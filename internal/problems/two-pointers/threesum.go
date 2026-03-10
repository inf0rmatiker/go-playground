package twopointers

import "slices"

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
*/

// When scanning, eliminate possibility of generating duplicate triplets.
// One idea: sort the array, then skip over duplicates. Use the two pointer method to find pairs that
// sum to the complement of the current number.
// O(nlogn) time for sorting, + O(n^2) time for scanning, O(1) space (not counting output).
// Total: O(n^2) time, O(1) space.
func threeSum(nums []int) [][]int {

	// Results
	triplets := [][]int{}

	slices.Sort(nums) // O(nlogn) time
	// [-1, 0, 1, 2, -1, -4] -> [-4, -1, -1, 0, 1, 2]

	// Maintain three pointers: i, j, k. i is the first number, j and k are the second and third numbers.
	// Example:
	//  i  j     k
	// [0, 0, 0, 0]
	// results = [[0, 0, 0]]

	for i := 0; i < len(nums)-2; i++ {
		// Skip over duplicates for i: If the current number is the same as the last number, continue.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Use two-pointers method to find pairs that, when combined with nums[i], sum to 0.
		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				// Skip over duplicates for j and k: If the current number is the same as the last number, continue.
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return triplets
}
