package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {

	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "leetcode example 1",
			intervals: [][]int{{1, 3}, {6, 7}, {1, 5}},
			expected:  [][]int{{1, 5}, {6, 7}},
		},
		{
			name:      "leetcode example 2",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := merge(tt.intervals)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("merge(%v) = %v, want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}
