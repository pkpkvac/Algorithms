package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{-1, 0, 2, 4, 6, 8},
			target:   4,
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{-1, 0, 2, 4, 6, 8},
			target:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch() = %v, want %v", result, tt.expected)
			}
		})
	}
}
