package dailytemperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {

	tests := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{
			name:         "leetcode example 1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "leetcode example 2",
			temperatures: []int{30, 38, 30, 36, 35, 40, 28},
			expected:     []int{1, 4, 1, 2, 1, 0, 0},
		},
		{
			name:         "leetcode example 3",
			temperatures: []int{22, 21, 20},
			expected:     []int{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("dailyTemperatures() = %v, want %v", result, tt.expected)
			}
		})
	}
}
