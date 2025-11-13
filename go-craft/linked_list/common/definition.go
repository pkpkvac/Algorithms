package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func NewListNodeWithNext(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := NewListNode(nums[0])
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = NewListNode(nums[i])
		current = current.Next
	}
	return head
}

func PrintLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
