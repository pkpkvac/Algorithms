package clonegraph

import (
	"algorithms/graphs/common"
	"testing"
)

func TestCloneGraph(t *testing.T) {

	tests := []struct {
		name     string
		adjList  [][]int
		expected *common.GraphNode
	}{
		{
			name:     "leetcode example 1",
			adjList:  [][]int{{2}, {1, 3}, {2}},
			expected: common.BuildGraphFromAdjList([][]int{{2}, {1, 3}, {2}}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cloneGraph(common.BuildGraphFromAdjList(tt.adjList))
			if !common.GraphsEqual(result, tt.expected) {
				t.Errorf("cloneGraph(%v) = %v, want %v", tt.adjList, result.Val, tt.expected.Val)
			}
		})
	}
}
