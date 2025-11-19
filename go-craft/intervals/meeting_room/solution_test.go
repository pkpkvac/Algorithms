package meetingroom

import (
	"reflect"
	"testing"
)

func TestCanAttendMeetings(t *testing.T) {

	tests := []struct {
		name      string
		intervals []Interval
		expected  bool
	}{
		{
			name:      "leetcode example 1",
			intervals: []Interval{{5, 10}, {15, 20}, {0, 30}},
			expected:  false,
		},
		{
			name:      "leetcode example 2",
			intervals: []Interval{{5, 8}, {9, 15}},
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canAttendMeetings(tt.intervals)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("canAttendMeetings(%v) = %v, want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}
