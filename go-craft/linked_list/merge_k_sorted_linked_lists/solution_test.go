package mergetwosortedlists

import (
	"algorithms/linked_list/common"
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {

	tests := []struct {
		name     string
		lists    []*common.ListNode
		expected *common.ListNode
	}{
		{name: "leetcode example 1",
			lists: []*common.ListNode{common.CreateLinkedList([]int{1, 2, 4}),
				common.CreateLinkedList([]int{1, 3, 5}),
				common.CreateLinkedList([]int{3, 6}),
			},
			expected: common.CreateLinkedList([]int{1, 1, 2, 3, 3, 4, 5, 6})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeKLists(tt.lists)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeKLists() = %v, want %v", result, tt.expected)
			}
		})
	}
}
