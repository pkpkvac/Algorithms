package mergestringsalternately

import (
	"reflect"
	"testing"
)

func TestMergeStringsAlternately(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{
			name:     "empty array",
			word1:    "",
			word2:    "",
			expected: "",
		},
		{
			name:     "single element",
			word1:    "a",
			word2:    "b",
			expected: "ab",
		},
		{
			name:     "multiple elements",
			word1:    "abc",
			word2:    "xyz",
			expected: "axbycz",
		},
		{
			name:     "leetcode example 1",
			word1:    "abc",
			word2:    "xyz",
			expected: "axbycz",
		},
		{
			name:     "leetcode example 2",
			word1:    "ab",
			word2:    "pqrs",
			expected: "apbqrs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeAlternately(tt.word1, tt.word2)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MergeStringsAlternately() = %v, want %v", got, tt.expected)
			}
		})
	}
}
