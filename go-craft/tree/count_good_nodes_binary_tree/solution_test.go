package countgoodnodesbinarytree

import (
	"algorithms/tree/common"
	"reflect"
	"testing"
)

func TestCountGoodNodesBinaryTree(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		expected int
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{2, 1, 1, 3, -999, 1, 5}),
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{1, 2, -1, 3, 4}),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goodNodes(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("goodNodes() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// --- Benchmarks ---

// buildCompleteTree builds a complete binary tree with values 1..n in level-order
// using the existing CreateBinaryTree helper.
func buildCompleteTree(n int) *common.TreeNode {
	if n <= 0 {
		return nil
	}
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = i + 1
	}
	return common.CreateBinaryTree(vals)
}
