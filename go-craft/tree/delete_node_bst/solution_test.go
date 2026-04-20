package deletenodebst

import (
	"algorithms/tree/common"
	"reflect"
	"testing"
)

func TestDeleteNodeBST(t *testing.T) {

	tests := []struct {
		name     string
		root     *common.TreeNode
		key      int
		expected *common.TreeNode
	}{
		{
			name:     "leetcode example 1",
			root:     common.CreateBinaryTree([]int{5, 3, 9, 1, 4}),
			key:      3,
			expected: common.CreateBinaryTree([]int{5, 4, 9, 1}),
		},
		{
			name:     "leetcode example 2",
			root:     common.CreateBinaryTree([]int{5, 3, 6, -999, 4, -999, 10, -999, -999, 7}),
			key:      3,
			expected: common.CreateBinaryTree([]int{5, 4, 6, -999, -999, -999, 10, 7}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := deleteNode(tt.root, tt.key)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("deleteNode() = %v, want %v", result, tt.expected)
			}
		})
	}

}
