package examples

import "fmt"

func ArrayExamples() {

	fmt.Println("=== Single-dimensional Arrays ===")

	// Initialize array of known size to default values
	var defaultArray [3]int
	for i := 0; i < len(defaultArray); i++ {
		fmt.Printf("defaultArray[%d]=[%d] ", i, defaultArray[i])
		fmt.Println()
	}

	// Initialized array using custom values, no need to specify size, can use '...' instead
	var initializedArray2 = [...]int{5, 6, 7, 8}
	for i := 0; i < len(initializedArray2); i++ {
		fmt.Printf("initializedArray2[%d]=[%d] ", i, initializedArray2[i])
		fmt.Println()
	}

	fmt.Println("=== Multi-dimensional Arrays ===")

	var multiDimArray = [2][3]int{
		{1, 2, 3},
		{2, 3, 4},
	}
	for i := 0; i < len(multiDimArray); i++ {
		for j := 0; j < len(multiDimArray[i]); j++ {
			fmt.Printf("multiDimArray[%d][%d]=[%d] ", i, j, multiDimArray[i][j])
			fmt.Println()
		}
	}

	// Change value of row
	multiDimArray[1] = [...]int{5, 6, 7}

	// Another way to iterate over arrays, using the range keyword
	for index, row := range multiDimArray {
		fmt.Println(index, row)
	}

	fmt.Println("=== Pass-by-value Array Example ===")

	passByValueExample(defaultArray)
	fmt.Println(defaultArray)

	fmt.Println("=== Pass-by-reference Array Example ===")

	passByReferenceExample(&defaultArray) // note the '&'
	fmt.Println(defaultArray)
}

// Arg declared as [3]int means a copy is made and passed by value.
// Note: We have to specify exact size because the size is part of the type for arrays.
// Use slices if you want them to be more dynamic.
func passByValueExample(myarray [3]int) {
	for index := range myarray {
		myarray[index]++
	}
	fmt.Println(myarray)
}

// Arg declared as *[3]int means the reference is passed.
// Note: We have to specify exact size because the size is part of the type for arrays.
// Use slices if you want them to be more dynamic.
func passByReferenceExample(myarray *[3]int) {
	for index := range myarray {
		myarray[index]++
	}
	fmt.Println(myarray)
}
