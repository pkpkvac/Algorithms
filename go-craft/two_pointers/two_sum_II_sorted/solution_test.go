package twosumiisorted

import (
	"reflect"
	"testing"
)

func TestTwoSumIISorted(t *testing.T) {

	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{
			name:     "empty array",
			numbers:  []int{},
			target:   0,
			expected: []int{},
		},
		{
			name:     "four elements",
			numbers:  []int{1, 2, 3, 4},
			target:   3,
			expected: []int{1, 2},
		},

		{
			name:     "four elements, target 7",
			numbers:  []int{2, 3, 4, 5},
			target:   7,
			expected: []int{1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSumIISorted(tt.numbers, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TwoSumIISorted() = %v, want %v", result, tt.expected)
			}
		})
	}

}
