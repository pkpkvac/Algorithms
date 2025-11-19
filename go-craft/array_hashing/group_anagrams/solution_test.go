package groupanagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "empty array",
			strs:     []string{},
			expected: [][]string{},
		},
		{
			name:     "single element",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "multiple elements",
			strs:     []string{"act", "pots", "tops", "cat", "stop", "hat"},
			expected: [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			name:     "leetcode example 1",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input slice to avoid modifying the original
			strs := make([]string, len(tt.strs))
			copy(strs, tt.strs)

			got := groupAnagrams(strs)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.expected)
			}
		})
	}
}
