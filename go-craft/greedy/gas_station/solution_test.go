package gas_station

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{
			name:     "leetcode example 1",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			name:     "single station with enough gas",
			gas:      []int{5},
			cost:     []int{4},
			expected: 0,
		},
		{
			name:     "single station not enough gas",
			gas:      []int{2},
			cost:     []int{3},
			expected: -1,
		},
		{
			name:     "start at beginning",
			gas:      []int{3, 1, 1},
			cost:     []int{1, 2, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canCompleteCircuit(tt.gas, tt.cost)
			if got != tt.expected {
				t.Errorf("canCompleteCircuit(%v, %v) = %v, want %v", tt.gas, tt.cost, got, tt.expected)
			}
		})
	}
}
