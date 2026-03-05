package jump_game

import "testing"

func TestJumpGame(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.nums)
			if result != tt.expected {
				t.Errorf("canJump(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
