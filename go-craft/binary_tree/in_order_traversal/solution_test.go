package inordertraversal

import (
	"algorithms/binary_tree/common"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		expected []int
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{1, 2, 3, 4, 5, 6, 7}),
			expected: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{1, 2, 3, -999, 4, 5, -999}),
			expected: []int{2, 4, 1, 5, 3},
		},
		{
			name:     "empty tree",
			root:     nil,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InorderTraversalRecursive(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InorderTraversal() = %v, want %v", result, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InorderTraversalIterative(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InorderTraversal() = %v, want %v", result, tt.expected)
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

func BenchmarkInorderTraversalRecursive(b *testing.B) {
	// Use a reasonably large complete tree
	root := buildCompleteTree(2047) // height ~11
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = InorderTraversalRecursive(root)
	}
}

func BenchmarkInorderTraversalIterative(b *testing.B) {
	root := buildCompleteTree(2047)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = InorderTraversalIterative(root)
	}
}
