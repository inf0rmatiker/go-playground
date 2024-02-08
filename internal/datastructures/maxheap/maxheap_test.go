package maxheap

import "testing"

/*
--- Helper functions for testing ---
*/

func arraysEqual(a []int, b []int) bool {
	if len(a) == len(b) {
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	return false
}

/*
--- Unit tests ---
*/

func TestLchild(t *testing.T) {
	inputs := []int{0, 1, 2, 3, 4}
	expectedResults := []int{1, 3, 5, 7, 9}
	for i, input := range inputs {
		actualResult := lchild(input)
		if actualResult != expectedResults[i] {
			t.Fail()
		}
	}
}

func TestRchild(t *testing.T) {
	inputs := []int{0, 1, 2, 3, 4}
	expectedResults := []int{2, 4, 6, 8, 10}
	for i, input := range inputs {
		actualResult := rchild(input)
		if actualResult != expectedResults[i] {
			t.Fail()
		}
	}
}

func TestSwap(t *testing.T) {
	testArray := []int{1, 2, 3, 4, 5, 6, 7}
	i, j := 0, len(testArray)-1
	expectedArray := []int{7, 2, 3, 4, 5, 6, 1}
	swap(i, j, testArray)
	if !arraysEqual(testArray, expectedArray) {
		t.Fail()
	}
}

func TestHasLeftChild(t *testing.T) {
	testArray := []int{1, 2, 3, 4, 5, 6, 7}
	if hasLeftChild(3, testArray) {
		t.Fail()
	}
}

func TestHasRightChild(t *testing.T) {
	testArray := []int{1, 2, 3, 4, 5, 6, 7}
	if hasRightChild(3, testArray) {
		t.Fail()
	}
}
