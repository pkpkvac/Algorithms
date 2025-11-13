package invertbinarytree

import (
	"algorithms/tree/common"
	"reflect"
	"testing"
)

func TestInvertBinaryTree(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		expected *common.TreeNode
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{1, 2, 3, 4, 5, 6, 7}),
			expected: common.CreateBinaryTree([]int{1, 3, 2, 7, 6, 5, 4}),
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{3, 2, 1}),
			expected: common.CreateBinaryTree([]int{3, 1, 2}),
		},
		{
			name:     "empty tree",
			root:     nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := invertTreeRecursive(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("invertTreeRecursive() = %v, want %v", result, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := invertTreeIterative(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("invertTreeIterative() = %v, want %v", result, tt.expected)
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

func BenchmarkInvertTreeRecursive(b *testing.B) {
	// Use a reasonably large complete tree
	root := buildCompleteTree(2047) // height ~11
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = invertTreeRecursive(root)
	}
}

func BenchmarkInvertTreeIterative(b *testing.B) {
	root := buildCompleteTree(2047)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = invertTreeIterative(root)
	}
}
