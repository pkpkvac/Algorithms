package common

import "fmt"

type DoublyLinkedNode struct {
	Val  int
	Next *DoublyLinkedNode
	Prev *DoublyLinkedNode
}

func NewDoublyLinkedNode(val int) *DoublyLinkedNode {
	return &DoublyLinkedNode{Val: val, Next: nil, Prev: nil}
}

func NewDoublyLinkedNodeWithNext(val int, next *DoublyLinkedNode) *DoublyLinkedNode {
	return &DoublyLinkedNode{Val: val, Next: next, Prev: nil}
}

func NewDoublyLinkedNodeWithPrev(val int, prev *DoublyLinkedNode) *DoublyLinkedNode {
	return &DoublyLinkedNode{Val: val, Next: nil, Prev: prev}
}

func NewDoublyLinkedNodeWithNextAndPrev(val int, next *DoublyLinkedNode, prev *DoublyLinkedNode) *DoublyLinkedNode {
	return &DoublyLinkedNode{Val: val, Next: next, Prev: prev}
}

func CreateDoublyLinkedList(nums []int) *DoublyLinkedNode {
	if len(nums) == 0 {
		return nil
	}
	head := NewDoublyLinkedNode(nums[0])
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = NewDoublyLinkedNode(nums[i])
		current.Next.Prev = current
		current = current.Next
	}
	return head
}

func PrintDoublyLinkedList(head *DoublyLinkedNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		current = current.Next
	}
	fmt.Println()
}

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
		fmt.Printf("%d", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func ListsEqual(list1 *ListNode, list2 *ListNode) bool {
	if list1 == nil && list2 == nil {
		return true
	}
	if list1 == nil || list2 == nil {
		return false
	}
	return list1.Val == list2.Val && ListsEqual(list1.Next, list2.Next)
}
