package kokoeatingbananas

import (
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {

	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "leetcode example 1",
			piles:    []int{1, 4, 3, 2},
			h:        9,
			expected: 2,
		},
		{
			name:     "leetcode example 2",
			piles:    []int{25, 10, 23, 4},
			h:        4,
			expected: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minEatingSpeed(tt.piles, tt.h)
			if result != tt.expected {
				t.Errorf("minEatingSpeed() = %v, want %v", result, tt.expected)
			}
		})
	}
}
