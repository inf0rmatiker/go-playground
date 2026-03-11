package linkedlists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseListIterative(t *testing.T) {

	values := []int{1, 2, 3, 4, 5}
	head := &ListNode{Val: values[0]}
	curr := head
	for i := 1; i < len(values); i++ {
		curr.Next = &ListNode{Val: values[i]}
		curr = curr.Next
	}

	fmt.Print("Before list: ")
	printList(head)

	newHead := ReverseListRecursive(head)
	curr = newHead
	i := len(values) - 1
	for curr != nil {
		assert.Equal(t, values[i], curr.Val)
		curr = curr.Next
		i--
	}
	assert.Equal(t, -1, i)

	fmt.Print("After list:  ")
	printList(newHead)
}

func TestReverseListRecursive(t *testing.T) {

	values := []int{1, 2, 3, 4, 5}
	head := &ListNode{Val: values[0]}
	curr := head
	for i := 1; i < len(values); i++ {
		curr.Next = &ListNode{Val: values[i]}
		curr = curr.Next
	}

	fmt.Print("Before list: ")
	printList(head)

	newHead := ReverseListRecursive(head)
	curr = newHead
	i := len(values) - 1
	for curr != nil {
		assert.Equal(t, values[i], curr.Val)
		curr = curr.Next
		i--
	}
	assert.Equal(t, -1, i)

	fmt.Print("After list:  ")
	printList(newHead)
}
