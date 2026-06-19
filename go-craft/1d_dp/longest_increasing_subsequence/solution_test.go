package longest_increasing_subsequence

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			name:     "leetcode example 3",
			nums:     []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "strictly decreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "strictly increasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			if got != tt.expected {
				t.Errorf("lengthOfLIS(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
