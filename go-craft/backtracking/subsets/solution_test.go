package subsets

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{7},
			expected: [][]int{{}, {7}},
		},
		{
			name:     "leetcode example 4",
			nums:     []int{1, 2, 3, 4},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}, {4}, {1, 4}, {2, 4}, {1, 2, 4}, {3, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsets(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("subsets(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
