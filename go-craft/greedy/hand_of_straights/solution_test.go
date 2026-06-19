package hand_of_straights

import (
	"testing"
)

func TestIsNStraightHand(t *testing.T) {
	tests := []struct {
		name      string
		hand      []int
		groupSize int
		expected  bool
	}{
		{
			name:      "leetcode example 1",
			hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize: 3,
			expected:  true,
		},
		{
			name:      "leetcode example 2",
			hand:      []int{1, 2, 3, 4, 5},
			groupSize: 4,
			expected:  false,
		},
		{
			name:      "single group",
			hand:      []int{1, 2, 3},
			groupSize: 3,
			expected:  true,
		},
		{
			name:      "impossible with gaps",
			hand:      []int{1, 2, 4, 5},
			groupSize: 2,
			expected:  false,
		},
		{
			name:      "group size of 1",
			hand:      []int{1, 1, 2, 2, 3, 3},
			groupSize: 1,
			expected:  true,
		},
		{
			name:      "duplicates form multiple groups",
			hand:      []int{1, 1, 2, 2, 3, 3},
			groupSize: 3,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNStraightHand(tt.hand, tt.groupSize)
			if got != tt.expected {
				t.Errorf("isNStraightHand(%v, %v) = %v, want %v", tt.hand, tt.groupSize, got, tt.expected)
			}
		})
	}
}
