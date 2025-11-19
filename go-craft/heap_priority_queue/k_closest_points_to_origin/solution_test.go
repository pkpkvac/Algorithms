package kclosestpointstoorigin

import (
	"reflect"
	"testing"
)

func TestKClosestPointsToOrigin(t *testing.T) {

	tests := []struct {
		name     string
		k        int
		points   [][]int
		expected [][]int
	}{

		{
			name:     "leetcode example 1",
			k:        1,
			points:   [][]int{{0, 2}},
			expected: [][]int{{0, 2}},
		},
		{
			name:     "leetcode example 2",
			k:        1,
			points:   [][]int{{0, 2}, {2, 1}},
			expected: [][]int{{0, 2}},
		},
		{
			name:     "leetcode example 3",
			k:        2,
			points:   [][]int{{0, 2}, {2, 0}, {2, 2}},
			expected: [][]int{{0, 2}, {2, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kClosestPointsToOrigin(tt.k, tt.points)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("kClosestPointsToOrigin(%v) = %v, want %v", tt.points, result, tt.expected)
			}
		})
	}
}
