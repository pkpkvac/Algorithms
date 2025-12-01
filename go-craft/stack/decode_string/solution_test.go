package decode_string

import (
	"testing"
)

func TestDecodeString(t *testing.T) {

	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "leetcode example 1",
			s:        "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			name:     "leetcode example 2",
			s:        "3[a2[c]]",
			expected: "accaccacc",
		},
		{
			name:     "leetcode example 3",
			s:        "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			name:     "leetcode example 4",
			s:        "2[3[a]2[bc]]",
			expected: "aaabcbcaaabcbc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := decodeString(tt.s)
			if result != tt.expected {
				t.Errorf("decodeString() = %v, want %v", result, tt.expected)
			}
		})
	}
}
