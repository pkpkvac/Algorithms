package topkfrequentelements

import (
	"reflect"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{

		{
			name:     "Leetcode example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{2, 1},
		},
		{
			name:     "Leetcode example 2",
			nums:     []int{1, 2, 2, 3, 3, 3},
			k:        2,
			expected: []int{2, 3},
		},
		{
			name:     "Leetcode example 3",
			nums:     []int{7, 7},
			k:        1,
			expected: []int{7},
		},
	}

	// Test input validation separately

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequentHeap(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("topKFrequent(%v, %v) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}

}
