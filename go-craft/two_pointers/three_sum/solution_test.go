package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "Leetcode example 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Leetcode example 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "Leetcode example 3",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("threeSum() = %v, want %v", result, tt.expected)
			}
		})
	}

}
