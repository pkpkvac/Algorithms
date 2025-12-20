package sumofallsubsetsxortotals

import (
	"testing"
)

func TestSubsetXORTotal(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{3, 1, 1},
			expected: 12,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{2, 4},
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsetXORTotal(tt.nums)
			if result != tt.expected {
				t.Errorf("subsetXORSum(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
