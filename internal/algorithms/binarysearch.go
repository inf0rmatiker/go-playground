package algorithms

// search searches a sorted array of ints for an element,
// returning a bool for its existence.
func binarySearch(element int, sortedArray []int) bool {

	if len(sortedArray) == 0 {
		return false
	}

	if len(sortedArray) == 1 {
		if element == sortedArray[0] {
			return true
		}
		return false
	}

	midIndex := len(sortedArray) / 2
	midElement := sortedArray[midIndex]
	if element == midElement {
		return true
	}
	if element > midElement {
		return binarySearch(element, sortedArray[midIndex+1:])
	}
	return binarySearch(element, sortedArray[:midIndex])
}

/*
You are given an m x n 2-D integer array matrix and an integer target.

Each row in matrix is sorted in non-decreasing order.
The first integer of every row is greater than the last integer of the previous row.
Return true if target exists within matrix or false otherwise.

Can you write a solution that runs in O(log(m * n)) time?

Input: matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]], target = 10

Output: true
*/

// Does binary search to find the row which *could* contain the target value.
// If a row is found which could contain the target, that row is returned.
// Otherwise, nil is returned.
func getRow(matrix [][]int, target int) []int {
	top, bottom := 0, len(matrix)-1

	for top <= bottom {
		// Grab the index of the row halfway between top/bottom.
		midpoint := int((top + bottom) / 2)
		firstEntry := matrix[midpoint][0]

		// If we're on the last row, just return it.
		if midpoint == len(matrix)-1 {
			return matrix[midpoint]
		}

		// Grab the first element of the next row so we have an
		// idea of what elements belong to the midpoint row.
		nextEntry := matrix[midpoint+1][0]
		if target < nextEntry && target >= firstEntry {
			return matrix[midpoint]
		} else if target >= nextEntry {
			// search bottom half
			top = midpoint + 1
		} else { // target < firstEntry
			// search top half
			bottom = midpoint - 1
		}
	}
	return nil
}

// Performs a matrix binary search to see if a target exists.
func searchMatrix(matrix [][]int, target int) bool {
	row := getRow(matrix, target)
	if row == nil {
		return false
	}

	// Perform binary search on row
	l, r := 0, len(row)-1
	for l <= r {
		mid := int((l + r) / 2)
		if row[mid] == target {
			return true
		}

		if target < row[mid] {
			// search left half
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
