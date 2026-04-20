package islands_and_treasure

import (
	"reflect"
	"testing"
)

const INF = 2147483647

func TestIslandsAndTreasure(t *testing.T) {

	tests := []struct {
		name     string
		grid     [][]int
		expected [][]int
	}{
		{
			name: "leetcode example 1",
			grid: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			expected: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "leetcode example 2",
			grid: [][]int{
				{0, -1},
				{INF, INF},
			},
			expected: [][]int{
				{0, -1},
				{1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			islandsAndTreasure(tt.grid)
			if !reflect.DeepEqual(tt.grid, tt.expected) {
				t.Errorf("got %v, want %v", tt.grid, tt.expected)
			}
		})
	}
}
