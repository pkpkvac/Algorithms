package maxdepthbinarytree

import (
	"algorithms/tree/common"
	"reflect"
	"testing"
)

func TestMaxDepthBinaryTree(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		expected int
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{1, 2, 3, -999, -999, 4}),
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{3, 2, 1}),
			expected: 2,
		},
		{
			name:     "empty tree",
			root:     common.CreateBinaryTree([]int{}),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxDepthBinaryTreeRecursive(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MaxDepthBinaryTreeRecursive() = %v, want %v", result, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxDepthBinaryTreeIterativeBFSA(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MaxDepthBinaryTreeIterative() = %v, want %v", result, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxDepthBinaryTreeIterativeBFSB(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MaxDepthBinaryTreeIterative() = %v, want %v", result, tt.expected)
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

func BenchmarkMaxDepthBinaryTreeRecursive(b *testing.B) {
	// Use a reasonably large complete tree
	root := buildCompleteTree(2047) // height ~11
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MaxDepthBinaryTreeRecursive(root)
	}
}

// func BenchmarkMaxDepthBinaryTreeIterative(b *testing.B) {
// 	root := buildCompleteTree(2047)
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		_ = MaxDepthBinaryTreeIterative(root)
// 	}
// }
