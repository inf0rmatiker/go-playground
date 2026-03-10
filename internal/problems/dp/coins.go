package dp

import (
	"fmt"
	"strings"
)

/*
Given an amount and the denominations of coins available, determine how many UNIQUE ways change can be made for amount.
There is a limitless supply of each coin type.

Example:

Input: n = 4, c = [1, 2, 3]

Total ways to make change for n = 4 are 4: [1, 1, 1, 1], [2, 2], [1, 1, 2], and [1, 3].
*/

// Recursive helper function for naive approach.
func helper(n int32, c []int64) int64 {

	// Base cases:
	if n == 0 {
		// if n == 0, we found a valid combination, so return 1.
		return 1
	} else if n < 0 || len(c) == 0 {
		// If n < 0 or no coins left, return 0.
		return 0
	}

	lastCoin := int32(c[len(c)-1])

	// Recursive case:
	// 1.) Use the coin: subtract it from sum and keep the same set of coins (since we can reuse coins).
	// 2.) Don't use the coin: keep the sum the same but remove the last coin from the set.
	return helper(n-lastCoin, c) + helper(n, c[:len(c)-1])
}

// O(2^n) time complexity, O(n) space complexity.
func GetWaysNaive(n int32, c []int64) int64 {
	return helper(n, c)
}

// This is a top-down dynamic programming approach using a 2D memoization table.
// O(n*m) time complexity, O(n*m) space complexity where n is the target sum and m is the number of coins.
func GetWaysDpSquare(n int32, c []int64) int64 {

	fmt.Printf("\nMemoization approach: n=%d c=%v\n", n, c)

	// First: Initialize a 2D memo table where the rows are the coins,
	// and columns are the sums from 0 to n.
	memo := make([][]int64, 0, len(c)+1)
	for range len(c) + 1 {
		memo = append(memo, make([]int64, n+1))
	}

	// Ignore this; just a function to show the memo table at various stages.
	printMemo := func() {
		fmt.Print("\n ")
		for i := range memo[0] {
			fmt.Printf("  %d", i)
		}
		fmt.Println("\n  ", strings.Repeat("-", 3*int(n+1)))
		for i := range memo {
			fmt.Printf("%d |", i)
			for j := range memo[i] {
				fmt.Printf(" %d ", memo[i][j])
			}
			fmt.Println()
		}
		fmt.Println()
	}

	for i := 0; i < len(c)+1; i++ {
		memo[i][0] = 1 // Base case: 1 way to make a sum of 0 with any number of coins (by using no coins).
	}

	// Start after the base case row and column.
	for coinSet := 1; coinSet < len(c)+1; coinSet++ {
		for sum := 1; sum < int(n+1); sum++ {
			// Ways to make sum without using this coin:
			// just inherit what it was without this coin.
			memo[coinSet][sum] = memo[coinSet-1][sum]

			coinValue := int(c[coinSet-1])
			canUse := sum >= coinValue
			fmt.Printf("Can use coin %d for sum %d: %t\n", c[coinSet-1], sum, canUse)

			if canUse {
				// If we can use the current coin, we add the ways to
				// make (sum - coinValue) using the same set of coins (coinSet).
				memo[coinSet][sum] += memo[coinSet][sum-coinValue]
			}

		}
	}

	printMemo()

	// The bottom-right cell of the table has our aggregated answer.
	return memo[len(c)][n]
}
