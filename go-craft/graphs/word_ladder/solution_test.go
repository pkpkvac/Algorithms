package wordladder

import (
	"testing"
)

func TestWordLadder(t *testing.T) {

	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			name:      "leetcode example 1",
			beginWord: "cat",
			endWord:   "sag",
			wordList:  []string{"bat", "bag", "sag", "dag", "dot"},
			expected:  4,
		},
		{
			name:      "leetcode example 2",
			beginWord: "cat",
			endWord:   "sag",
			wordList:  []string{"bat", "bag", "sat", "dag", "dot"},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if result != tt.expected {
				t.Errorf("ladderLength(%v, %v, %v) = %v, want %v", tt.beginWord, tt.endWord, tt.wordList, result, tt.expected)
			}
		})
	}
}
