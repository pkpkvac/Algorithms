package decode_ways

import (
	"testing"
)

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "leetcode example 1",
			s:        "12",
			expected: 2,
		},
		{
			name:     "leetcode example 2",
			s:        "226",
			expected: 3,
		},
		{
			name:     "leetcode example 3",
			s:        "06",
			expected: 0,
		},
		{
			name:     "single digit",
			s:        "1",
			expected: 1,
		},
		{
			name:     "leading zero",
			s:        "0",
			expected: 0,
		},
		{
			name:     "multiple ways",
			s:        "111",
			expected: 3,
		},
		{
			name:     "zero in middle",
			s:        "10",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDecodings(tt.s)
			if got != tt.expected {
				t.Errorf("numDecodings(%q) = %v, want %v", tt.s, got, tt.expected)
			}
		})
	}
}
