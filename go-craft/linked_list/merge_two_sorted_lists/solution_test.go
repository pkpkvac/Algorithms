package mergetwosortedlists

import (
	"algorithms/linked_list/common"
	"reflect"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {

	tests := []struct {
		name     string
		list1    *common.ListNode
		list2    *common.ListNode
		expected *common.ListNode
	}{
		{
			name:     "leetcode example 1",
			list1:    common.CreateLinkedList([]int{1, 2, 4}),
			list2:    common.CreateLinkedList([]int{1, 3, 5}),
			expected: common.CreateLinkedList([]int{1, 1, 2, 3, 4, 5}),
		},
		{
			name:     "empty list",
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
		{
			name:     "single node",
			list1:    common.CreateLinkedList([]int{1}),
			list2:    common.CreateLinkedList([]int{2}),
			expected: common.CreateLinkedList([]int{1, 2}),
		},
		{
			name:     "multiple nodes",
			list1:    common.CreateLinkedList([]int{}),
			list2:    common.CreateLinkedList([]int{1, 2}),
			expected: common.CreateLinkedList([]int{1, 2}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeTwoLists(tt.list1, tt.list2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeTwoLists() = %v, want %v", result, tt.expected)
			}
		})
	}
}
