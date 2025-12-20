package subtreeofanothertree

import (
	"algorithms/tree/common"
	"testing"
)

func TestSubtreeOfAnotherTree(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		subRoot  *common.TreeNode
		expected bool
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{1, 2, 3, 4, 5, -999, -999, 6}),
			subRoot:  common.CreateBinaryTree([]int{2, 4, 5}),
			expected: false,
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{1, 2, 3, 4, 5}),
			subRoot:  common.CreateBinaryTree([]int{2, 4, 5}),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubtree(tt.root, tt.subRoot)
			if result != tt.expected {
				t.Errorf("isSubtree() = %v, want %v", result, tt.expected)
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
