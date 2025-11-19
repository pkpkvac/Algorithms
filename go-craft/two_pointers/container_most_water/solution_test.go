package containermostwater

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			heights:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "leetcode example 2",
			heights:  []int{1, 7, 2, 5, 4, 7, 3, 6},
			expected: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.heights)
			if got != tt.expected {
				t.Errorf("maxArea() = %v, want %v", got, tt.expected)
			}
		})
	}
}
