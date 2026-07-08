package unique_paths

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			name:     "leetcode example 1",
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			name:     "leetcode example 2",
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			name:     "1x1 grid",
			m:        1,
			n:        1,
			expected: 1,
		},
		{
			name:     "1xn grid",
			m:        1,
			n:        10,
			expected: 1,
		},
		{
			name:     "mx1 grid",
			m:        10,
			n:        1,
			expected: 1,
		},
		{
			name:     "square grid",
			m:        5,
			n:        5,
			expected: 70,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.m, tt.n)
			if got != tt.expected {
				t.Errorf("uniquePaths(%v, %v) = %v, want %v", tt.m, tt.n, got, tt.expected)
			}
		})
	}
}
