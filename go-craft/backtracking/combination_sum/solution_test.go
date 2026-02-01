package combination_sum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{2, 5, 6, 9},
			target:   9,
			expected: [][]int{{2, 2, 5}, {9}},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{3, 4, 5},
			target:   16,
			expected: [][]int{{3, 3, 3, 3, 4}, {3, 3, 5, 5}, {4, 4, 4, 4}, {3, 4, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("combinationSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
