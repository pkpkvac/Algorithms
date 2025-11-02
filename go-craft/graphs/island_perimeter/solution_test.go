package islandperimeter

import (
	"testing"
)

func TestIslandPerimeter(t *testing.T) {

	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "leetcode example 1",
			grid: [][]int{{1, 1, 0, 0},
				{1, 0, 0, 0},
				{1, 1, 1, 0},
				{0, 0, 1, 1}},
			expected: 18,
		},
		{
			name:     "leetcode example 2",
			grid:     [][]int{{1, 0}},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := islandPerimeter(tt.grid)
			if result != tt.expected {
				t.Errorf("islandPerimeter(%v) = %v, want %v", tt.grid, result, tt.expected)
			}
		})
	}
}
