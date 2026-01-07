package minimumsize

import (
	"testing"
)

func TestMinSubArrayLenBruteForce(t *testing.T) {

	tests := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			target:   10,
			nums:     []int{2, 1, 5, 1, 5, 3},
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			name:     "leetcode example 3",
			target:   5,
			nums:     []int{1, 2, 1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSubArrayLenCleaner(tt.target, tt.nums)
			if result != tt.expected {
				t.Errorf("minSubArrayLenCleaner() = %v, want %v", result, tt.expected)
			}
		})
	}
}
