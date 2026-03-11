package linkedlists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseListIterative reverses a singly linked list in-place and returns the new head of the reversed list.
// Iterative approach: O(n) time complexity, O(1) space complexity.
func ReverseListIterative(head *ListNode) *ListNode {
	var prev, curr, next *ListNode
	curr = head

	for curr != nil {
		next = curr.Next // save reference to next node, before we change the ptr
		curr.Next = prev // change the ptr to point to prev node
		prev = curr      // move prev to current node
		curr = next      // move curr to saved next node
	}
	return prev
}

// ReverseListRecursive reverses a singly linked list in-place and returns the new head of the reversed list.
// Recursive approach: O(n) time complexity, O(n) space complexity due to call stack.
func ReverseListRecursive(head *ListNode) *ListNode {
	// Base case: if head is nil or only one node, return head
	if head == nil || head.Next == nil {
		return head
	}

	tail := ReverseListRecursive(head.Next) // reverse the rest of the list

	head.Next.Next = head // make the next node point back to current node
	head.Next = nil       // set current node's next to nil to avoid cycle

	return tail
}

// Helper function to print the linked list, used in tests for visualization.
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
