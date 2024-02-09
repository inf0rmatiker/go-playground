package linkedlist

import "fmt"

type node struct {
	// The actual data of the linked list node
	data int

	// Pointer to next node in our list
	next *node

	// Pointer to previous node in our list
	prev *node
}

// LinkedList is a doubly-linked list,
// traversable either forwards or backwards.
type LinkedList struct {

	// First node in our list
	head *node

	// Last node in our list
	tail *node

	// Size of our list
	size int
}

func New() LinkedList {
	return LinkedList{
		size: 0,
	}
}

// Print traverses the list and prints all the data items.
func (l *LinkedList) Print() {
	fmt.Println("List size:", l.size)
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

/* --- Retrieval functions --- */

// getNode is a helper function to retrieve a node at
// a specified index. Returns true and the node pointer
// if successful, false and nil otherwise.
func (l *LinkedList) getNode(index int) (bool, *node) {
	// Out-of-bounds check
	if index >= l.size || index < 0 {
		return false, nil
	}

	if index == 0 {
		return true, l.head
	} else if index == l.size-1 {
		return true, l.tail
	}

	// Loop through list until i matches index
	i := 0
	current := l.head
	for i < index {
		current = current.next
		i++
	}

	return true, current
}

// Get retrieves a data item at a specified index.
// Returns true and the data item if successful,
// false, -1 otherwise.
func (l *LinkedList) Get(index int) (bool, int) {
	if ok, value := l.getNode(index); ok {
		return ok, value.data
	}
	return false, -1
}

/* --- Addition functions --- */

// Append adds an element to the tail of the list.
func (l *LinkedList) Append(data int) {
	element := node{data: data}
	if l.size == 0 {
		// We'll also have to update the head to be the new element
		l.head = &element
	} else {
		// Make current tail's next ptr point to new element,
		// and element's prev ptr point to old tail
		l.tail.next = &element
		element.prev = l.tail
	}

	// Change tail to the last element
	l.tail = &element
	l.size++
}

// Prepend adds an element to the beginning of the list
func (l *LinkedList) Prepend(data int) {
	element := node{data: data}
	if l.size == 0 {
		// We'll also have to update the tail to be the new element
		l.tail = &element
	} else {
		// Make current head's prev ptr point to new element,
		// and element's next ptr point to old head
		l.head.prev = &element
		element.next = l.head
	}

	// Change head to the new first element
	l.head = &element
	l.size++
}

// Insert inserts a data item at a specified index.
// Returns true if the operation was successful, false otherwise.
func (l *LinkedList) Insert(data int, index int) bool {
	element := node{data: data}
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		// Inserting as new head
		l.Prepend(data)
	} else if index == l.size {
		// Inserting as new tail
		l.Append(data)
	} else {
		// Inserting somewhere in the middle
		_, prevNode := l.getNode(index - 1)
		_, nextNode := l.getNode(index + 1)
		prevNode.next = &element
		nextNode.prev = &element
		element.prev = prevNode
		element.next = nextNode
		l.size++
	}
	return true
}

/* Removal functions */

// Remove removes an item from the linked list at a specified index.
// Returns a bool true for success, false for failure.
func (l *LinkedList) Remove(index int) bool {

	// Bounds check. This also covers an empty linked list.
	if index < 0 || index >= l.size {
		return false
	}

	// We're deleting the only node in the list, which is
	// both the tail and the head
	if l.size == 1 {
		l.tail = nil
		l.head = nil
	} else if index == l.size-1 {
		// Deleting the tail: set the tail's prev next ptr to nil
		newTail := l.tail.prev
		newTail.next = nil
		l.tail = newTail
	} else if index == 0 {
		// Deleting the head: set the head's next prev ptr to nil
		newHead := l.head.next
		newHead.prev = nil
		l.head = newHead
	} else {
		// Deleting something in the middle. Make the two nodes
		// on either side "skip" the node slated for removal.
		_, removalNode := l.getNode(index)
		before := removalNode.prev
		after := removalNode.next
		before.next = after
		after.prev = before
	}

	l.size--
	return true
}
