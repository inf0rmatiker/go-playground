package stack

import (
	"testing"
)

func TestStack_Size(t *testing.T) {
	s := Stack[int]{}
	if s.Size() != 0 {
		t.Fail()
	}

	s = Stack[int]{
		array: []int{1, 2, 3},
	}
	if s.Size() != 3 {
		t.Fail()
	}
}

func TestStack_Put(t *testing.T) {
	inputs := []int{1, 2, 3}
	s := Stack[int]{}

	for _, input := range inputs {
		s.Put(input)
	}
	if s.Size() != len(inputs) {
		t.Fail()
	}

	for i := len(inputs) - 1; i >= 0; i-- {
		expected := inputs[i]
		output, ok := s.Pop()
		if !ok || output != expected {
			t.Fail()
		}
	}
}
