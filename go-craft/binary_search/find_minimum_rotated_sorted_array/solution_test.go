package find_minimum_rotated_sorted_array

import (
	"testing"
)

func TestFindMin(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{3, 4, 5, 6, 1, 2},
			expected: 1,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{4, 5, 0, 1, 2, 3},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinProper(tt.nums)
			// result := findMin(tt.nums)
			if result != tt.expected {
				t.Errorf("findMin() = %v, want %v", result, tt.expected)
			}
		})
	}
}
