package remove_nth_node_from_end_of_list

import (
	"algorithms/linked_list/common"
)

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	var prev *common.ListNode
	var curr *common.ListNode

	prev = head
	curr = head

	for i := 0; i < n; i++ {
		curr = curr.Next
	}

	if curr == nil {
		// then move head ahead once
		// bc n == length list
		head = head.Next
		return head
	}

	for curr.Next != nil {
		prev = prev.Next
		curr = curr.Next
	}

	prev.Next = prev.Next.Next

	return head
}
