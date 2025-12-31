package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelfBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "leetcode example 1",
			nums:     []int{1, 2, 4, 6},
			expected: []int{48, 24, 12, 8},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{-1, 0, 1, 2, 3},
			expected: []int{0, -6, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.expected)
			}
		})
	}
}
