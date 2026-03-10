package dp

import "fmt"

/*
https://www.youtube.com/watch?v=_i4Yxeh5ceQ&t=65s

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. How many distinct ways can
you reach the top?

Example:
steps n = 3
         _3_
      _2_|  |
   _1_|     |
_0_|________|
*/

func StaircaseRecursiveSolution(n int) int {
	// Base case
	if n == 0 {
		// no steps left, we are at the top already
		return 1
	}

	if n-2 >= 0 {
		// We can either take 1 or 2 steps without overshooting
		return StaircaseRecursiveSolution(n-1) + StaircaseRecursiveSolution(n-2)
	}
	// We can only take 1 step without overshooting
	return StaircaseRecursiveSolution(n - 1)
}

func StaircaseDPIterativeBottomUp(n int) int {
	// Initialize a memoization structure to remember previously computed steps.
	// Each index will represent the solution found for that number of steps.
	memo := make([]int, n+1)

	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	}

	memo[n], memo[n-1] = 1, 1

	for i := n - 2; i >= 0; i-- {
		memo[i] = memo[i+1] + memo[i+2]
	}
	return memo[0]
}

func StaircaseDPIterativeBottomUpMinimalMemory(n int) int {
	// Instead of using a full memo array to store every value, we only have
	// to store the last two in variables.
	one, two := 1, 1
	for ; n > 1; n-- {
		temp := one
		one = one + two
		two = temp
	}
	return one
}

func StaircaseDPRecursiveBottomUp(n int) int {
	// Initialize a memoization structure to remember previously computed steps.
	// Each index will represent the solution found for that number of steps.
	memo := make([]int, n+1)

	// Anonymous helper function for recursion
	var recursiveHelper func(int, int, []int) int
	recursiveHelper = func(current int, n int, memo []int) int {
		// Corner base case: n < 2, there's only 1 step to take
		// n = 2
		// [ 1 , 1 ]
		if n < 2 {
			return 1
		}

		if current >= n-1 {
			memo[current] = 1
		} else {
			// current < n-1
			memo[current] = memo[current+1] + memo[current+2]
			if current == 0 {
				return memo[0] // Base case, we've reached the answer for the start of the staircase
			}
		}
		return recursiveHelper(current-1, n, memo)
	}

	return recursiveHelper(n, n, memo)
}

/*
https://www.youtube.com/watch?v=_i4Yxeh5ceQ&t=1143s

Min Cost Climbing Stairs

You are given an integer array cost where cost[i] is the cost of the ith step on the staircase.
Once you pay the cost, you can either climb one or two steps. You can either start from the step
with index 0, or the step with index 1.

What's the minimum cost to climb the staircase?

Example:

cost = [10, 15, 20]
*/

func min(a int, b int) int {
	if a < b {
		return a
	} else if a > b {
		return b
	}
	return a
}

func MinCostStaircaseRecursive(cost []int) int {

	var recursiveHelper func([]int) int
	recursiveHelper = func(cost []int) int {
		// Base case: there is only two steps left.
		if len(cost) <= 2 {
			// Base case: there are only 2 steps left
			return cost[0]
		}

		// Pay current step and return whichever costs less, taking one step or two
		return cost[0] + min(recursiveHelper(cost[1:]), recursiveHelper(cost[2:]))
	}

	// Determine whether to start at first or second step
	return min(recursiveHelper(cost), recursiveHelper(cost[1:]))
}

func MinCostStaircaseMemoization(cost []int) int {

	// cost: [10, 15, 20]
	// memo: [10, 15,  0]
	// cost: [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
	// memo: [5, 103, 4, 4, 3, 101, 2, 1, 100, 0]
	memo := make([]int, len(cost))
	last := len(cost) - 1
	memo[last] = cost[last]
	memo[last-1] = cost[last-1] + cost[last]
	memo[last-2] = cost[last-2] + cost[last]

	for i := last - 3; i >= 0; i-- {
		memo[i] = cost[i] + min(memo[i+1], memo[i+2])
	}
	fmt.Println(memo)
	return min(memo[0], memo[1])
}
