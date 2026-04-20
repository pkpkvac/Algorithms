package max_subarray

import "testing"

func TestMaxSubarray(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{2, -3, 4, -2, 2, 1, -1, 4},
			expected: 8,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{-1},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSubArray(tt.nums)
			if result != tt.expected {
				t.Errorf("maxSubArray(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
