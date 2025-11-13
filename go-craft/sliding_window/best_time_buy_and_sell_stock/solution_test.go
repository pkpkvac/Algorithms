package besttimebuyandsellstock

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {

	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{

		{
			name:     "example 1",
			numbers:  []int{10, 1, 5, 6, 7, 1},
			expected: 6,
		},
		{
			name:     "example 2",
			numbers:  []int{10, 8, 7, 5, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProfit(tt.numbers)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MaxProfit() = %v, want %v", result, tt.expected)
			}
		})
	}

}
