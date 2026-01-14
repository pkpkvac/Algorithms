package levelordertraversal

import (
	"algorithms/tree/common"
	"reflect"
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		expected [][]int
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{1, 2, 3, 4, 5, 6, 7}),
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrderTraversal() = %v, want %v", result, tt.expected)
			}
		})
	}
}
