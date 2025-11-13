package mergetwosortedlists

import (
	"algorithms/linked_list/common"
)

func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// Attach remaining nodes from whichever list still has nodes
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}
