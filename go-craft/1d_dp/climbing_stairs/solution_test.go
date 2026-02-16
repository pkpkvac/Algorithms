package climbing_stairs

import (
	"reflect"
	"testing"
)

func TestClimbingStairs(t *testing.T) {

	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "leetcode example 1",
			n:        2,
			expected: 2,
		},
		{
			name:     "leetcode example 2",
			n:        3,
			expected: 3,
		},
		{
			name:     "leetcode example 3",
			n:        4,
			expected: 5,
		},
		{
			name:     "leetcode example 4",
			n:        5,
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("climbStairs(%v) = %v, want %v", tt.n, got, tt.expected)
			}
		})
	}
}
