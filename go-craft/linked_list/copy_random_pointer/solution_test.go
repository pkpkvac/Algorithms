package copyrandompointer

import (
	"algorithms/linked_list/common"
	"testing"
)

func TestCopyRandomList(t *testing.T) {

	// Example: [[3,null],[7,3],[4,0],[5,1]]
	// Node values: 3 -> 7 -> 4 -> 5
	// Random:      3->nil, 7->node[3]=5, 4->node[0]=3, 5->node[1]=7

	node0 := &common.ListNodeWithRandom{Val: 3}
	node1 := &common.ListNodeWithRandom{Val: 7}
	node2 := &common.ListNodeWithRandom{Val: 4}
	node3 := &common.ListNodeWithRandom{Val: 5}

	node0.Next = node1
	node1.Next = node2
	node2.Next = node3

	// node0.Random = nil  (zero value, no action needed)
	node1.Random = node3 // 7 -> 5
	node2.Random = node0 // 4 -> 3
	node3.Random = node1 // 5 -> 7

	tests := []struct {
		name     string
		head     *common.ListNodeWithRandom
		expected *common.ListNodeWithRandom
	}{
		{
			name:     "leetcode example 1",
			head:     node0,
			expected: node0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copyRandomList(tt.head)
			if !common.ListNodeWithRandomEqual(result, tt.expected) {
				t.Errorf("copyRandomList() = %v, want %v", result.Val, tt.expected.Val)
			}
		})
	}
}
