package permutations

import (
	"sort"
	"testing"
)

func sortPerms(perms [][]int) {
	for _, p := range perms {
		sort.Ints(p)
	}
	sort.Slice(perms, func(i, j int) bool {
		for k := 0; k < len(perms[i]); k++ {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}
		return false
	})
}

func TestPermute(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permute(tt.nums)
			sortPerms(result)
			sortPerms(tt.expected)
			if len(result) != len(tt.expected) {
				t.Fatalf("permute(%v) returned %d results, want %d", tt.nums, len(result), len(tt.expected))
			}
			for i := range result {
				for j := range result[i] {
					if result[i][j] != tt.expected[i][j] {
						t.Errorf("permute(%v)[%d] = %v, want %v", tt.nums, i, result[i], tt.expected[i])
						break
					}
				}
			}
		})
	}
}
