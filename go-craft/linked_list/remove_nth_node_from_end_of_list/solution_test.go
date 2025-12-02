package remove_nth_node_from_end_of_list

import (
	"algorithms/linked_list/common"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {

	tests := []struct {
		name     string
		head     *common.ListNode
		n        int
		expected *common.ListNode
	}{
		{
			name:     "leetcode example 1",
			head:     common.CreateLinkedList([]int{0, 1, 2, 3}),
			n:        2,
			expected: common.CreateLinkedList([]int{0, 1, 3}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeNthFromEnd(tt.head, tt.n)
			if !common.ListsEqual(result, tt.expected) {
				t.Errorf("removeNthFromEnd() = %v, want %v", result, tt.expected)
			}
		})
	}
}
