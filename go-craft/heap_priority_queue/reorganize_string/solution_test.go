package reorganizestring

import (
	"testing"
)

func TestReorganizeString(t *testing.T) {

	tests := []struct {
		name     string
		s        string
		expected string
	}{

		{
			name:     "leetcode example 1",
			s:        "axyy",
			expected: "xyay",
		},
		{
			name:     "leetcode example 2",
			s:        "abbccdd",
			expected: "abcdbcd",
		},
		{
			name:     "leetcode example 3",
			s:        "ccccd",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reorganizeString(tt.s)
			if result != tt.expected {
				t.Errorf("reorganizeString(%v) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
