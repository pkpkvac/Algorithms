package lemonadechange

import "testing"

func TestLemonadeChange(t *testing.T) {

	tests := []struct {
		name     string
		bills    []int
		expected bool
	}{
		{
			name:     "leetcode example 1",
			bills:    []int{5, 10, 5, 5, 20},
			expected: true,
		},
		{
			name:     "leetcode example 2",
			bills:    []int{5, 20, 10, 5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lemonadeChange(tt.bills)
			if result != tt.expected {
				t.Errorf("lemonadeChange(%v) = %v, want %v", tt.bills, result, tt.expected)
			}
		})
	}
}
