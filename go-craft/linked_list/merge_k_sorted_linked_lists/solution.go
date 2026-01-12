package mergetwosortedlists

import (
	"algorithms/linked_list/common"
	"container/heap"
	"fmt"
)

type MinHeap []*common.ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*common.ListNode))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MinHeap) Peek() int {
	return h[0].Val
}

func mergeKLists(lists []*common.ListNode) *common.ListNode {
	// the strategy here for merging to two lists will not work
	// K is an undetermined number.

	// we could throw all the nodes into a min heap
	// sorting the nodes based on value
	// then continue to pop and assign next, eventually
	// getting the ascending list
	h := &MinHeap{}
	heap.Init(h)

	for _, listItem := range lists {
		for listItem != nil {
			heap.Push(h, listItem)
			listItem = listItem.Next
		}
	}

	var head *common.ListNode
	var curr *common.ListNode
	for h.Len() > 0 {

		node := heap.Pop(h).(*common.ListNode)
		fmt.Println("VALUE", node.Val)

		if head == nil {
			head = node
		}
		if curr == nil {
			curr = node
		} else {
			curr.Next = node
			curr = curr.Next
		}

	}

	return head
}

// func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
// 	dummy := &common.ListNode{}
// 	tail := dummy

// 	for list1 != nil && list2 != nil {
// 		if list1.Val <= list2.Val {
// 			tail.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			tail.Next = list2
// 			list2 = list2.Next
// 		}
// 		tail = tail.Next
// 	}

// 	// Attach remaining nodes from whichever list still has nodes
// 	if list1 != nil {
// 		tail.Next = list1
// 	}
// 	if list2 != nil {
// 		tail.Next = list2
// 	}

// 	return dummy.Next
// }
