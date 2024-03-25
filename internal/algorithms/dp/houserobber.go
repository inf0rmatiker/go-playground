package dp

/*
https://www.youtube.com/watch?v=_i4Yxeh5ceQ&t=2296s

You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping
you from robbing each of them is that adjacent houses have security systems connected,
and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.

Example:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
*/

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func RobHousesRecursive(nums []int) int {
	// Base cases:
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// Return the max of either:
	//  1) robbing the first house, skipping the second
	//  2) skipping the first house
	return max(nums[0]+RobHousesRecursive(nums[2:]), RobHousesRecursive(nums[1:]))
}

func RobHousesMemoization(nums []int) int {

	// nums: [1,2,3,1]
	// memo: [1,2,4,4]
	// nums: [1,3,2,2,5]
	// memo: [1,3,3,5,8]
	memo := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			memo[i] = v
		} else if i == 1 {
			memo[i] = max(memo[0], v)
		} else {
			memo[i] = max(memo[i-1], v+memo[i-2])
		}
	}
	return memo[len(memo)-1]
}

func RobHousesMemoryEfficient(nums []int) int {

	// nums: [1,2,3,1]
	// nums: [1,3,2,2,5]
	// [first, second, n, n+1, n+2, ...]
	first, second := 0, 0
	for _, v := range nums {
		temp := max(v+first, second)
		first = second
		second = temp
	}
	return second
}

/*
After robbing those houses on that street, the thief has found himself
a new place for his thievery so that he will not get too much attention.
This time, all houses at this place are arranged in a circle. That means
the first house is the neighbor of the last one. Meanwhile, the security
system for these houses remain the same as for those in the previous street.

Given a list of non-negative integers representing the amount of money of
each house, determine the maximum amount of money you can rob tonight without
alerting the police.
*/

func RobHouses2(nums []int) int {
	return max(RobHousesMemoryEfficient(nums[1:]), RobHousesMemoryEfficient(nums[:len(nums)-1]))
}
