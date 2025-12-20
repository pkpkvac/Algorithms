package taskscheduler

import (
	"testing"
)

func TestLeastInterval(t *testing.T) {

	tests := []struct {
		name     string
		tasks    []byte
		n        int
		expected int
	}{

		{
			name:     "leetcode example 1",
			tasks:    []byte{'X', 'X', 'Y', 'Y'},
			n:        2,
			expected: 5,
		},
		{
			name:     "leetcode example 2",
			tasks:    []byte{'A', 'A', 'A', 'B', 'C'},
			n:        3,
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leastInterval(tt.tasks, tt.n)
			if result != tt.expected {
				t.Errorf("leastInterval(%v, %d) = %d, want %d", tt.tasks, tt.n, result, tt.expected)
			}
		})
	}
}
