package reorderlinkedlist

import (
	"algorithms/linked_list/common"
	reverselinkedlist "algorithms/linked_list/reverse_linked_list"
)

func ReorderListArrayMethod(head *common.ListNode) *common.ListNode {

	var arr []*common.ListNode

	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	l := 0
	r := len(arr) - 1

	for l < r {
		// Connect left to right
		arr[l].Next = arr[r]
		l++

		// Check if pointers meet before connecting
		if l >= r {
			break
		}

		// Connect right to the new left
		arr[r].Next = arr[l]
		r--
	}

	// Terminate the list (important!)
	arr[l].Next = nil

	head = arr[0]

	common.PrintLinkedList(head)

	return head
}

func ReorderListReversalMethod(head *common.ListNode) *common.ListNode {

	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// Save the second half before disconnecting
	secondHalf := slow.Next
	slow.Next = nil // Disconnect first half from second half

	// Reverse the second half and capture the new head
	l2 := reverselinkedlist.ReverseList(secondHalf)

	l1 := head

	for l1 != nil && l2 != nil {

		l1tmp := l1.Next
		l1.Next = l2

		l2tmp := l2.Next
		l2.Next = l1tmp

		l2 = l2tmp
		l1 = l1tmp

	}

	common.PrintLinkedList(head)
	return head
}
