package subsets_ii

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "leetcode example 1",
			nums: []int{1, 2, 2},
			expected: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name: "multiple duplicates",
			nums: []int{4, 4, 4, 1, 4},
			expected: [][]int{
				{},
				{1},
				{1, 4},
				{1, 4, 4},
				{1, 4, 4, 4},
				{1, 4, 4, 4, 4},
				{4},
				{4, 4},
				{4, 4, 4},
				{4, 4, 4, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.nums)
			
			// Sort both slices for comparison
			sortSubsets := func(subsets [][]int) {
				for _, subset := range subsets {
					sort.Ints(subset)
				}
				sort.Slice(subsets, func(i, j int) bool {
					if len(subsets[i]) != len(subsets[j]) {
						return len(subsets[i]) < len(subsets[j])
					}
					for k := 0; k < len(subsets[i]); k++ {
						if subsets[i][k] != subsets[j][k] {
							return subsets[i][k] < subsets[j][k]
						}
					}
					return false
				})
			}
			
			sortSubsets(got)
			sortSubsets(tt.expected)
			
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("subsetsWithDup(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
