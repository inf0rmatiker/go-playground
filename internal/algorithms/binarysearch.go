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
