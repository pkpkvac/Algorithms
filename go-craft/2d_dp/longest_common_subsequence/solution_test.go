package longest_common_subsequence

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{
			name:     "leetcode example 1",
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			name:     "leetcode example 2",
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			name:     "leetcode example 3",
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			name:     "empty strings",
			text1:    "",
			text2:    "abc",
			expected: 0,
		},
		{
			name:     "partial match",
			text1:    "abcdef",
			text2:    "fedcba",
			expected: 1,
		},
		{
			name:     "interleaved match",
			text1:    "programming",
			text2:    "gaming",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonSubsequence(tt.text1, tt.text2)
			if got != tt.expected {
				t.Errorf("longestCommonSubsequence(%q, %q) = %v, want %v", tt.text1, tt.text2, got, tt.expected)
			}
		})
	}
}
