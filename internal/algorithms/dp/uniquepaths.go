package dp

import "fmt"

/*
https://leetcode.com/problems/unique-paths/description/

Solution walkthrough: https://youtu.be/qMky6D6YtXU?si=sXYjSz971kvx9hKv&t=71

There is a robot on an m x n grid.
The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique
paths that the robot can take to reach the bottom-right corner.

Example:
Input: m = 3, n = 7
Output: 28

Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
*/

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Printf("[ ")
		for _, num := range row {
			fmt.Printf("%3d ", num)
		}
		fmt.Printf("]\n")
	}
}

func UniquePaths(m int, n int) int {

	// Some corner cases
	if m == 1 || n == 1 {
		return 1
	}

	// Make a grid for answers
	// I.e. m=4, n=3:
	//     0 1 2 3
	// 0 [ 0 0 0 0 ]
	// 1 [ 0 0 0 0 ]
	// 2 [ 0 0 0 0 ]
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	// Bottom-up DP solution, start at bottom-right
	// end and work our way back to start of top left
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == 0 && j == 0 {
				// Base case, we're at the beginning
				printGrid(grid)
				return grid[1][0] + grid[0][1]
			}
			if i == m-1 && j == n-1 {
				// We're at the end, mark it a 1
				grid[i][j] = 1
			} else if i < m-1 && j < n-1 {
				// We're somewhere off the right and bottom walls, so we have 2 movement choices
				grid[i][j] = grid[i+1][j] + grid[i][j+1]
			} else if i < m-1 {
				// We're on the right wall, but off the bottom, so only have 1 movement choice
				grid[i][j] = grid[i+1][j]
			} else {
				// We're on the bottom, but off the right wall, so only have 1 movement choice
				grid[i][j] = grid[i][j+1]
			}
		}
	}

	return -1
}
