package cycle_exists

import (
	"algorithms/linked_list/common"
	"testing"
)

func TestHasCycle(t *testing.T) {

	cycleHead := common.CreateLinkedList([]int{1, 2, 3, 4})
	cycleHead.Next.Next.Next.Next = cycleHead.Next

	tests := []struct {
		name     string
		head     *common.ListNode
		expected bool
	}{
		{
			name:     "leetcode example 1",
			head:     cycleHead,
			expected: true,
		},

		{
			name:     "single node",
			head:     common.CreateLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasCycle(tt.head)
			if result != tt.expected {
				t.Errorf("hasCycle() = %v, want %v", result, tt.expected)
			}
		})
	}
}
