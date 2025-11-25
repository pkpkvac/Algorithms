package nonoverlappingintervals

import (
	"testing"
)

func TestNonOverlappingIntervals(t *testing.T) {

	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "leetcode example 1",
			intervals: [][]int{{1, 2}, {2, 4}, {1, 4}},
			expected:  1,
		},
		{
			name:      "leetcode example 2",
			intervals: [][]int{{1, 2}, {2, 4}},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := eraseOverlapIntervals(tt.intervals)
			if result != tt.expected {
				t.Errorf("eraseOverlapIntervals(%v) = %v, want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}
