package algorithms

func mergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	middle := len(input) / 2
	sortedFirstHalf := mergeSort(input[0:middle])
	sortedSecondHalf := mergeSort(input[middle:])

	return merge(sortedFirstHalf, sortedSecondHalf)
}

func merge(first []int, second []int) []int {
	combined := make([]int, len(first)+len(second))

	// Indices for combined, first, second
	f, s := 0, 0
	for c := 0; c < len(combined); c++ {
		if f < len(first) && s < len(second) {
			// We get to choose between first and second sorted arrays
			if first[f] <= second[s] {
				// Choose from the first array
				combined[c] = first[f]
				f++
			} else {
				// Choose from second array
				combined[c] = second[s]
				s++
			}
		} else if f < len(first) {
			// We've run out of second items
			combined[c] = first[f]
			f++
		} else {
			// We've run out of first items
			combined[c] = second[s]
			s++
		}
	}

	return combined
}
