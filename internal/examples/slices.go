package examples

import "fmt"

func SliceExamples() {

	fmt.Println("=== Single-dimension Slice ===")

	// Empty slice
	var defaultSlice []int
	fmt.Printf("%T\n", defaultSlice) // []int

	// Append to slice
	newSlice := append(defaultSlice, 1, 2) // Append elements
	fmt.Println(newSlice)                  // [1 2]

	y := []int{3, 4, 5}
	newSlice = append(newSlice, y...) // Append another slice using '...'
	fmt.Println(newSlice)             // [1 2 3 4 5]

	fmt.Println("=== Pass-by-reference Slice Example ===")

	passByReference(newSlice) // Changes the value of the underlying array
	fmt.Println(newSlice)     // [2 3 4 5 6]

	fmt.Println("=== Copy Slice Example ===")
	slice1 := []int{9, 9, 9}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Printf("slice1: %v, slice2: %v\n", slice1, slice2)
	fmt.Printf("Length of slice1: %d, capacity of slice1: %d\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 3)
	fmt.Printf("Length of slice1: %d, capacity of slice1: %d\n", len(slice1), cap(slice1))

	slicingSlices()
}

// Slices are just headers referencing an underlying array.
// So while the *header* is copied, the underlying array remains the same,
// no copy is created.
func passByReference(myslice []int) {
	for i := range myslice {
		myslice[i] += 1
	}
	fmt.Println(myslice)
}

func slicingSlices() {

	fmt.Println("=== Slicing Slices Example ===")
	a := []int{1, 2, 3, 4, 5}
	b := a[:2]
	fmt.Printf("a=%v, len(a)=%d, cap(a)=%d\n", a, len(a), cap(a))
	fmt.Printf("b=%v, len(b)=%d, cap(b)=%d\n", b, len(b), cap(b))

}
