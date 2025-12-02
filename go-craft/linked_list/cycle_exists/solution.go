package cycle_exists

import (
	"algorithms/linked_list/common"
)

func hasCycle(head *common.ListNode) bool {

	var slow *common.ListNode
	var fast *common.ListNode

	slow = head
	fast = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}

	return false
}

// func ReverseList(head *common.ListNode) *common.ListNode {

// 	if head == nil {
// 		return nil
// 	}

// 	prev := (*common.ListNode)(nil)
// 	current := head

// 	for current != nil {

// 		next := current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next

// 	}

// 	return prev
// }
