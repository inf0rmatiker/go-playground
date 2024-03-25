package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Size(t *testing.T) {
	q := Queue[int]{}
	if q.Size() != 0 {
		t.Fail()
	}

	q = Queue[int]{
		array: []int{1, 2, 3},
	}
	if q.Size() != 3 {
		t.Fail()
	}
}

func TestQueue_Enqueue(t *testing.T) {
	inputs := []int{1, 2, 3}
	q := Queue[int]{}

	for _, input := range inputs {
		q.Enqueue(input)
	}

	if q.Size() != len(inputs) {
		t.Fail()
	}

	for _, input := range inputs {
		if output, ok := q.Dequeue(); !ok || input != output {
			fmt.Printf("input=%d, output=%d\n", input, output)
			t.Fail()
		}
	}
}
