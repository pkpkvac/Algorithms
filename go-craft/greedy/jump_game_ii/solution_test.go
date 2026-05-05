package jump_game_ii

import (
	"testing"
)

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		
		{
			name:     "two elements",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "all max jumps",
			nums:     []int{5, 4, 3, 2, 1, 0},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jump(tt.nums)
			if got != tt.expected {
				t.Errorf("jump(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
