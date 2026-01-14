package maxpathsum

import (
	"algorithms/tree/common"
	"reflect"
	"testing"
)

func TestMaxPathSum(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		expected int
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{1, 2, 3}),
			expected: 6,
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{-15, 10, 20, -999, -999, 15, 5, -5}),
			expected: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxPathSum(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("maxPathSum() = %v, want %v", result, tt.expected)
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
