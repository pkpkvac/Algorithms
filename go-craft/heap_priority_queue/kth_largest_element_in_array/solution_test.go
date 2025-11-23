package kthlargestelementinarray

import (
	"testing"
)

func TestKthLargestElementInArray(t *testing.T) {

	tests := []struct {
		name     string
		k        int
		nums     []int
		expected int
	}{

		{
			name:     "leetcode example 1",
			k:        2,
			nums:     []int{2, 3, 1, 5, 4},
			expected: 4,
		},
		{
			name:     "leetcode example 2",
			k:        3,
			nums:     []int{2, 3, 1, 1, 5, 5, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("findKthLargest(%v, %v) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
