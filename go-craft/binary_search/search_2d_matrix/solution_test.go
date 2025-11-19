package search2dmatrix

import (
	"testing"
)

func TestSearch2DMatrix(t *testing.T) {

	tests := []struct {
		name     string
		nums     [][]int
		target   int
		expected bool
	}{
		{
			name:     "leetcode example 1",
			nums:     [][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}},
			target:   14,
			expected: true,
		},
		{
			name:     "leetcode example 2",
			nums:     [][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}},
			target:   15,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Search2DMatrix(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("Search2DMatrix() = %v, want %v", result, tt.expected)
			}
		})
	}
}
