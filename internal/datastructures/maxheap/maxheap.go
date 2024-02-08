package maxheap

import "fmt"

type MaxHeap struct {

	// The underlying array for the heap structure.
	underlyingArray []int
}

// New creates a blank MaxHeap initialized with default values
func New() MaxHeap {
	return MaxHeap{
		underlyingArray: []int{},
	}
}

// Print prints the underlying array of the heap
func (h *MaxHeap) Print() {
	fmt.Println("underlyingArray:", h.underlyingArray)
}

// GetSize returns the size of the heap
func (h *MaxHeap) GetSize() int {
	return len(h.underlyingArray)
}

func (h *MaxHeap) get(index int) int {
	return h.underlyingArray[index]
}

// parent returns the parent index of a given index
func parent(index int) int {
	return (index - 1) / 2
}

// lchild returns the left child of a parent index
func lchild(index int) int {
	return (index * 2) + 1
}

// lchild returns the right child of a parent index
func rchild(index int) int {
	return (index * 2) + 2
}

// bubbleUp bubbles up the last element to its destination in the heap
// by continuously swapping with its parent
func bubbleUp(index int, underlyingArray []int) bool {

	// Sanity check for index
	if index < 0 || index >= len(underlyingArray) {
		return false
	}

	// Base case: we've reached the root, return
	if index == 0 {
		return true
	}

	parentIndex := parent(index)
	currentKey := underlyingArray[index]
	parentKey := underlyingArray[parentIndex]

	// We're now smaller or equal to our parent, so return
	if currentKey <= parentKey {
		return true
	}

	// Swap with parent and make recursive call with parent index
	underlyingArray[index], underlyingArray[parentIndex] = underlyingArray[parentIndex], underlyingArray[index]
	return bubbleUp(parentIndex, underlyingArray)
}

// Insert adds a new key to the MaxHeap and sorts it into place
func (h *MaxHeap) Insert(key int) bool {
	// Add the key to the end of the array (bottom of heap)
	h.underlyingArray = append(h.underlyingArray, key)
	return bubbleUp(len(h.underlyingArray)-1, h.underlyingArray)
}

func hasRightChild(parentIndex int, underlyingArray []int) bool {
	return rchild(parentIndex) < len(underlyingArray)
}

func hasLeftChild(parentIndex int, underlyingArray []int) bool {
	return lchild(parentIndex) < len(underlyingArray)
}

func swap(i int, j int, underlyingArray []int) {
	underlyingArray[i], underlyingArray[j] = underlyingArray[j], underlyingArray[i]
}

func heapify(parentIndex int, underlyingArray []int) {
	parentKey := underlyingArray[parentIndex]
	if !hasLeftChild(parentIndex, underlyingArray) {
		// We have no children, so we're done.
		return
	} else if !hasRightChild(parentIndex, underlyingArray) {
		// We only have a left child, so see if we need to swap, otherwise we're done.
		if underlyingArray[parentIndex] < underlyingArray[lchild(parentIndex)] {
			swap(parentIndex, lchild(parentIndex), underlyingArray)
		}
		return
	}

	// We have both children, see if we're larger than either of them
	lchildKey := underlyingArray[lchild(parentIndex)]
	rchildKey := underlyingArray[rchild(parentIndex)]
	if parentKey < lchildKey || parentKey < rchildKey {
		// Find which of the children is largest.
		// Swap with the largest child, then recursively heapify from that child's index.
		if rchildKey >= lchildKey {
			swap(parentIndex, rchild(parentIndex), underlyingArray)
			heapify(rchild(parentIndex), underlyingArray)
		} else {
			swap(parentIndex, lchild(parentIndex), underlyingArray)
			heapify(lchild(parentIndex), underlyingArray)
		}
	}
}

// Extract pops the root key off the heap and reorders the tree
func (h *MaxHeap) Extract() (bool, int) {

	// Sanity check: is someone trying to extract from an empty heap?
	if len(h.underlyingArray) < 1 {
		return false, -1
	}

	root := h.underlyingArray[0]
	size := h.GetSize()

	if size == 1 {
		// Only one thing in the heap so just remove it, and we're done.
		h.underlyingArray = h.underlyingArray[:0]
	} else {
		// Set new root to last element in heap, remove last element, then heapify down
		h.underlyingArray[0] = h.get(size - 1)
		h.underlyingArray = h.underlyingArray[:size-1]
		heapify(0, h.underlyingArray)
	}

	return true, root
}
