package courseschedule_ii

import (
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {

	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{
			name:          "leetcode example 1",
			numCourses:    2,
			prerequisites: [][]int{{0, 1}},
			expected:      []int{0, 1},
		},
		{
			name:          "leetcode example 2",
			numCourses:    2,
			prerequisites: [][]int{{0, 1}, {1, 0}},
			expected:      []int{},
		},
		{
			name:          "leetcode example 3",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1, 2, 3},
		},
		{
			name:          "leetcode example 4",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findOrder(tt.numCourses, tt.prerequisites)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findOrder(%v, %v) = %v, want %v", tt.numCourses, tt.prerequisites, result, tt.expected)
			}
		})
	}
}
