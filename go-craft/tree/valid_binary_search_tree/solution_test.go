package valid_binary_search_tree

import (
	"algorithms/tree/common"
	"testing"
)

func TestIsValidBST(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		expected bool
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{2, 1, 3}),
			expected: true,
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{3, 2, 1}),
			expected: false,
		},
		{
			name:     "different trees",
			root:     common.CreateBinaryTree([]int{1, 2, 3}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidBST(tt.root)
			if result != tt.expected {
				t.Errorf("isValidBST() = %v, want %v", result, tt.expected)
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
