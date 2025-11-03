package concatenationofarray

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: []int{1, 1},
		},
		{
			name:     "multiple elements",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3, 1, 2, 3},
		},
		{
			name:     "leetcode example 1",
			nums:     []int{1, 4, 1, 2},
			expected: []int{1, 4, 1, 2, 1, 4, 1, 2},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{22, 21, 20, 1},
			expected: []int{22, 21, 20, 1, 22, 21, 20, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getConcatenation(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GetConcatenation() = %v, want %v", got, tt.expected)
			}
		})
	}
}
