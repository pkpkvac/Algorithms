package courseschedule

import (
	"testing"
)

func TestCanFinish(t *testing.T) {

	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:          "leetcode example 1",
			numCourses:    2,
			prerequisites: [][]int{{0, 1}},
			expected:      true,
		},
		{
			name:          "leetcode example 2",
			numCourses:    2,
			prerequisites: [][]int{{0, 1}, {1, 0}},
			expected:      false,
		},
		{
			name:          "leetcode example 3",
			numCourses:    4,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}, {0, 4}},
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canFinish(tt.numCourses, tt.prerequisites)
			if result != tt.expected {
				t.Errorf("canFinish(%v, %v) = %v, want %v", tt.numCourses, tt.prerequisites, result, tt.expected)
			}
		})
	}
}
