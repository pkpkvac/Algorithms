package word_break

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "leetcode example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "leetcode example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "leetcode example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			name:     "single word match",
			s:        "dog",
			wordDict: []string{"dog"},
			expected: true,
		},
		{
			name:     "no match",
			s:        "abc",
			wordDict: []string{"a", "b"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			if got != tt.expected {
				t.Errorf("wordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.expected)
			}
		})
	}
}
