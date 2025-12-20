package lowestcommonancestorinbst

import (
	"algorithms/tree/common"
	"testing"
)

func TestLowestCommonAncestorInBST(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		p        int
		q        int
		expected *common.TreeNode
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{5, 3, 8, 1, 4, 7, 9, -999, 2}),
			p:        3,
			q:        8,
			expected: common.NewTreeNode(5),
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{5, 3, 8, 1, 4, 7, 9, -999, 2}),
			p:        3,
			q:        4,
			expected: common.NewTreeNode(3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lowestCommonAncestor(tt.root, common.NewTreeNode(tt.p), common.NewTreeNode(tt.q))
			if result.Val != tt.expected.Val {
				t.Errorf("lowestCommonAncestor() = %d, want %d", result.Val, tt.expected.Val)
			}
		})
	}

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		result := PostorderTraversalIterative(tt.root)
	// 		if !reflect.DeepEqual(result, tt.expected) {
	// 			t.Errorf("PostorderTraversalIterative() = %v, want %v", result, tt.expected)
	// 		}
	// 	})
	// }
}
