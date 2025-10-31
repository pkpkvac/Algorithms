package reverselinkedlist

import (
	"algorithms/linked_list/common"
)

func ReverseList(head *common.ListNode) *common.ListNode {

	if head == nil {
		return nil
	}

	prev := (*common.ListNode)(nil)
	current := head

	for current != nil {

		next := current.Next
		current.Next = prev
		prev = current
		current = next

	}

	return prev
}
