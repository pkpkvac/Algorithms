package validanagram

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {

	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{

		{
			name:     "Proper heights",
			s:        "racecar",
			t:        "carrace",
			expected: true,
		},
		{
			name:     "Single person",
			s:        "alice",
			t:        "alice",
			expected: true,
		},
		{
			name:     "Different lengths",
			s:        "alice",
			t:        "alicebob",
			expected: false,
		},
		{
			name:     "Different characters",
			s:        "alice",
			t:        "bob",
			expected: false,
		},
	}

	// Test input validation separately

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("isAnagram(%v, %v) = %v, want %v", tt.s, tt.t, result, tt.expected)
			}
		})
	}

}
