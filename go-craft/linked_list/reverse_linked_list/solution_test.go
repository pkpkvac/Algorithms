package reverselinkedlist

import (
	"algorithms/linked_list/common"
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {

	tests := []struct {
		name     string
		head     *common.ListNode
		expected *common.ListNode
	}{
		{
			name:     "leetcode example 1",
			head:     common.CreateLinkedList([]int{0, 1, 2, 3}),
			expected: common.CreateLinkedList([]int{3, 2, 1, 0}),
		},
		{
			name:     "empty list",
			head:     nil,
			expected: nil,
		},
		{
			name:     "single node",
			head:     common.CreateLinkedList([]int{1}),
			expected: common.CreateLinkedList([]int{1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseList(tt.head)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReverseList() = %v, want %v", result, tt.expected)
			}
		})
	}
}
