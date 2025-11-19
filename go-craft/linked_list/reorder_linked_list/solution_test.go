package reorderlinkedlist

import (
	"algorithms/linked_list/common"
	"reflect"
	"testing"
)

func TestReorderLinkedList(t *testing.T) {

	tests := []struct {
		name     string
		head     *common.ListNode
		expected *common.ListNode
	}{
		{
			name:     "leetcode example 1",
			head:     common.CreateLinkedList([]int{0, 1, 2, 3, 4, 5, 6}),
			expected: common.CreateLinkedList([]int{0, 6, 1, 5, 2, 4, 3}),
		},
		{
			name:     "leetcode example 2",
			head:     common.CreateLinkedList([]int{2, 4, 6, 8, 10}),
			expected: common.CreateLinkedList([]int{2, 10, 4, 8, 6}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := ReorderListArrayMethod(tt.head)
			result := ReorderListReversalMethod(tt.head)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReorderList() = %v, want %v", result, tt.expected)
			}
		})
	}
}
