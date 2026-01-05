package trappingrainwater

import (
	"testing"
)

func TestTrapBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			heights:  []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got := trapBruteForce(tt.heights)
			got := trap(tt.heights)
			if got != tt.expected {
				t.Errorf("trapBruteForce() = %v, want %v", got, tt.expected)
			}
		})
	}
}
