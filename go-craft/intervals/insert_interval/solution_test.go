package insertinterval

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {

	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name:        "leetcode example 1",
			intervals:   [][]int{{1, 3}, {4, 6}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 6}},
		},
		{
			name:        "leetcode example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {9, 10}},
			newInterval: []int{6, 7},
			expected:    [][]int{{1, 2}, {3, 5}, {6, 7}, {9, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, result, tt.expected)
			}
		})
	}
}
