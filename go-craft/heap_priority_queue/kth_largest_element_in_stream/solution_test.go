package kthlargestelementinastream

import (
	"testing"
)

func TestKthLargestElementInStream(t *testing.T) {

	tests := []struct {
		name     string
		k        int
		nums     []int
		adds     []int
		expected []int
	}{

		{
			name:     "leetcode example 1",
			k:        3,
			nums:     []int{1, 2, 3, 3},
			adds:     []int{3, 5, 6, 7, 8},
			expected: []int{3, 3, 3, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kl := Constructor(tt.k, tt.nums)
			for i, add := range tt.adds {
				result := kl.Add(add)
				if result != tt.expected[i] {
					t.Errorf("Add(%d) = %v, want %v", add, result, tt.expected[i])
				}
			}
		})
	}
}
