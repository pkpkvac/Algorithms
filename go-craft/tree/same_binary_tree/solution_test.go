package same_binary_tree

import (
	"algorithms/tree/common"
	"testing"
)

func TestIsSameTree(t *testing.T) {

	tests := []struct {
		name     string
		p        *common.TreeNode
		q        *common.TreeNode
		expected bool
	}{
		{
			name:     "leetcode example 1",
			p:        common.CreateBinaryTree([]int{1, 2, 3, -999, -999, 4}),
			q:        common.CreateBinaryTree([]int{1, 2, 3, -999, -999, 4}),
			expected: true,
		},
		{
			name:     "leetcode example 2",
			p:        common.CreateBinaryTree([]int{3, 2, 1}),
			q:        common.CreateBinaryTree([]int{3, 2, 1}),
			expected: true,
		},
		{
			name:     "empty tree",
			p:        common.CreateBinaryTree([]int{}),
			q:        common.CreateBinaryTree([]int{}),
			expected: true,
		},
		{
			name:     "different trees",
			p:        common.CreateBinaryTree([]int{1, 2, 3}),
			q:        common.CreateBinaryTree([]int{1, 3, 2}),
			expected: false,
		},
		{
			name:     "different trees with different values",
			p:        common.CreateBinaryTree([]int{4, 7}),
			q:        common.CreateBinaryTree([]int{4, -999, 7}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSameTree(tt.p, tt.q)
			if result != tt.expected {
				t.Errorf("isSameTree() = %v, want %v", result, tt.expected)
			}
		})
	}

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		result := DiameterOfBinaryTreeIterativeBFSA(tt.root)
	// 		if !reflect.DeepEqual(result, tt.expected) {
	// 			t.Errorf("DiameterOfBinaryTreeIterative() = %v, want %v", result, tt.expected)
	// 		}
	// 	})
	// }

}
