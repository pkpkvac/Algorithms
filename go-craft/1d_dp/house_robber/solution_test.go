package house_robber

import (
	"reflect"
	"testing"
)

func TestHouseRobber(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{1, 1, 3, 3},
			expected: 2,
		},
		{
			name:     "leetcode example 2",
			nums:     []int{2, 9, 8, 3, 6},
			expected: 3,
		},
		{
			name:     "leetcode example 3",
			nums:     []int{1, 2, 3, 1},
			expected: 5,
		},
		{
			name:     "leetcode example 4",
			nums:     []int{2, 7, 9, 3, 1},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("rob(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
