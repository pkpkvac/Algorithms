package rottingfruit

import (
	"testing"
)

func TestRottingFruit(t *testing.T) {

	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "leetcode example 1",
			grid: [][]byte{{1, 1, 0},
				{0, 1, 1},
				{0, 1, 2}},
			expected: 4,
		},
		{
			name: "leetcode example 2",
			grid: [][]byte{
				{1, 0, 1},
				{0, 2, 0},
				{1, 0, 1},
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rottingFruit(tt.grid)
			if result != tt.expected {
				t.Errorf("rottingFruit(%v) = %v, want %v", tt.grid, result, tt.expected)
			}
		})
	}
}
