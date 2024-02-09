package maxheap

import (
	"fmt"
	"testing"
)

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

// isValidMaxHeap is a recursive function to validate the properties
// are maintained of a max heap.
func isValidMaxHeap(parentIndex int, underlyingArray []int) bool {
	// Base case
	if parentIndex >= len(underlyingArray) {
		// We're beyond the array, or array is empty
		return true
	}

	parentItem := underlyingArray[parentIndex]
	if hasLeftChild(parentIndex, underlyingArray) {
		lchildItem := underlyingArray[lchild(parentIndex)]
		if parentItem < lchildItem {
			fmt.Printf("Invalid MaxHeap: %d parent < %d child\n", parentItem, lchildItem)
			return false
		}
		if hasRightChild(parentIndex, underlyingArray) {
			// There's a potential to have more subtrees when we have both children
			rchildItem := underlyingArray[rchild(parentIndex)]
			if parentItem < rchildItem {
				fmt.Printf("Invalid MaxHeap: %d parent < %d child\n", parentItem, rchildItem)
				return false
			}

			// Parent was indeed larger-or-equal-to both lchild and rchild, so recursively evaluate children
			return isValidMaxHeap(lchild(parentIndex), underlyingArray) &&
				isValidMaxHeap(rchild(parentIndex), underlyingArray)
		}
	} else if hasRightChild(parentIndex, underlyingArray) {
		fmt.Println("Cannot have a right child without a left child")
		return false
	}
	return true // only child, which is valid
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

func TestNew(t *testing.T) {
	heap := New()
	if heap.GetSize() != 0 {
		t.Fail()
	}
}

func TestGetSize(t *testing.T) {
	size := 5
	testHeap := MaxHeap{
		underlyingArray: make([]int, size),
	}
	if testHeap.GetSize() != size {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	testHeap := MaxHeap{
		underlyingArray: []int{5, 3, 2},
	}
	if testHeap.get(1) != 3 {
		t.Fail()
	}
}

func TestBubbleUpEmpty(t *testing.T) {
	underlyingArray := []int{}
	if bubbleUp(0, underlyingArray) ||
		bubbleUp(-1, underlyingArray) ||
		bubbleUp(1, underlyingArray) {
		t.Fail()
	}
}

func TestBubbleUpSingle(t *testing.T) {
	underlyingArray := []int{1}
	if !bubbleUp(0, underlyingArray) {
		t.Fail()
	}
}

func TestBubbleUpMultiple(t *testing.T) {
	underlyingArray := []int{6, 5, 4, 3, 3, 7}
	if !bubbleUp(len(underlyingArray)-1, underlyingArray) {
		t.Fail()
	}
	expectedArray := []int{7, 5, 6, 3, 3, 4}
	if !arraysEqual(underlyingArray, expectedArray) {
		t.Fail()
	}
}

func TestMaxHeap_Extract(t *testing.T) {
	testHeap := MaxHeap{
		underlyingArray: []int{7, 5, 6, 3, 3, 4},
	}
	extractionOrder := []int{7, 6, 5, 4, 3, 3}
	expectedSize := 6
	for _, v := range extractionOrder {
		ok, next := testHeap.Extract()
		expectedSize -= 1
		if !ok {
			t.Fail()
		} else if next != v {
			t.Fail()
		} else if !isValidMaxHeap(0, testHeap.underlyingArray) {
			t.Fail()
		} else if testHeap.GetSize() != expectedSize {
			t.Fail()
		}
	}
}

func TestMaxHeap_Insert(t *testing.T) {
	testHeap := New()
	testElements := []int{
		9, 3, 10, 15, 2, 1, 45, 4, 7, 8, 12, 0,
	}
	expectedSize := 0
	for _, v := range testElements {
		testHeap.Insert(v)
		expectedSize += 1
		if testHeap.GetSize() != expectedSize {
			t.Fail()
		} else if !isValidMaxHeap(0, testHeap.underlyingArray) {
			t.Fail()
		}
	}
}
