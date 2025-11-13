package laststoneweight

import (
	"testing"
)

func TestLastStoneWeight(t *testing.T) {

	tests := []struct {
		name     string
		stones   []int
		expected int
	}{

		{
			name:     "leetcode example 1",
			stones:   []int{2, 3, 6, 2, 4},
			expected: 1,
		},
		{
			name:     "leetcode example 2",
			stones:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "leetcode example 3",
			stones:   []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lastStoneWeight(tt.stones)
			if result != tt.expected {
				t.Errorf("lastStoneWeight(%v) = %d, want %d", tt.stones, result, tt.expected)
			}
		})
	}
}
