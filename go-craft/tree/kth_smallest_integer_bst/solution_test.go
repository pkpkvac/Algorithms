package kth_smallest_integer_bst

import (
	"algorithms/tree/common"
	"testing"
)

func TestKthSmallest(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		k        int
		expected int
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{2, 1, 3}),
			k:        1,
			expected: 1,
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{4, 3, 5, 2, -999}),
			k:        4,
			expected: 5,
		},
		{
			name:     "leetcode example 3",
			root:     common.CreateBinaryTree([]int{4, 3, 5, 2, -999}),
			k:        2,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kthSmallest(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("kthSmallest() = %v, want %v", result, tt.expected)
			}
		})
	}
}
