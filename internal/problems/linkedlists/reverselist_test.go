package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to build a linked list from a slice of values
func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	curr := head
	for i := 1; i < len(values); i++ {
		curr.Next = &ListNode{Val: values[i]}
		curr = curr.Next
	}
	return head
}

// Helper function to convert a linked list to a slice of values
func listToSlice(head *ListNode) []int {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}

// Helper function to reverse a slice
func reverseSlice(s []int) []int {
	reversed := make([]int, len(s))
	for i, v := range s {
		reversed[len(s)-1-i] = v
	}
	return reversed
}

func TestReverseList(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "standard list",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "two elements",
			input: []int{1, 2},
			want:  []int{2, 1},
		},
		{
			name:  "empty list",
			input: []int{},
			want:  []int{},
		},
	}

	implementations := []struct {
		name string
		fn   func(*ListNode) *ListNode
	}{
		{"Iterative", ReverseListIterative},
		{"Recursive", ReverseListRecursive},
	}

	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					head := buildList(tc.input)
					reversed := impl.fn(head)
					got := listToSlice(reversed)
					assert.Equal(t, tc.want, got)
				})
			}
		})
	}
}
