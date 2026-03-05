package coin_change

import (
	"reflect"
	"testing"
)

func TestCoinChange(t *testing.T) {

	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "leetcode example 1",
			coins:    []int{1, 5, 10},
			amount:   12,
			expected: 2,
		},
		{
			name:     "leetcode example 2",
			coins:    []int{2},
			amount:   3,
			expected: 3,
		},
		{
			name:     "leetcode example 3",
			coins:    []int{1},
			amount:   0,
			expected: 5,
		},
		{
			name:     "leetcode example 4",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("coinChange(%v, %v) = %v, want %v", tt.coins, tt.amount, got, tt.expected)
			}
		})
	}
}
