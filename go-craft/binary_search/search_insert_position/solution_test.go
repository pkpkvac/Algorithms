package searchinsertposition

import (
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{-1, 0, 2, 4, 6, 8},
			target:   5,
			expected: 4,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{-1, 0, 2, 4, 6, 8},
			target:   10,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchInsert(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("searchInsert() = %v, want %v", result, tt.expected)
			}
		})
	}
}
