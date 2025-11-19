package meetingroomii

import (
	"reflect"
	"testing"
)

func TestMinMeetingRooms(t *testing.T) {

	tests := []struct {
		name      string
		intervals []Interval
		expected  int
	}{
		{
			name:      "leetcode example 1",
			intervals: []Interval{{0, 40}, {5, 10}, {15, 20}},
			expected:  2,
		},
		{
			name:      "leetcode example 2",
			intervals: []Interval{{0, 30}, {5, 10}, {15, 20}, {30, 40}},
			expected:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minMeetingRoomsWithHeap(tt.intervals)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("minMeetingRooms(%v) = %v, want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}
