package stack

import "fmt"

// Stack - implementation of a LIFO structure using generics.
type Stack[K any] struct {

	// array - our underlying array
	array []K
}

// Size returns the size of the stack.
func (s *Stack[K]) Size() int {
	return len(s.array)
}

// Print prints the stack's contents.
func (s *Stack[K]) Print() {
	fmt.Println(s.array)
}

// Put adds an item to the top of the stack.
func (s *Stack[K]) Put(item K) {
	s.array = append(s.array, item)
}

// Pop removes and returns an item from the top of the stack.
func (s *Stack[K]) Pop() (item K, ok bool) {
	if len(s.array) == 0 {
		item, ok = *new(K), false
	} else {
		item, ok = s.array[len(s.array)-1], true
		s.array = s.array[:len(s.array)-1]
	}
	return
}
