package baseballgame

import (
	"testing"
)

func TestBaseballGame(t *testing.T) {

	tests := []struct {
		name       string
		operations []string
		expected   int
	}{
		{
			name:       "leetcode example 1",
			operations: []string{"1", "2", "+", "C", "5", "D"},
			expected:   18,
		},
		{
			name:       "leetcode example 2",
			operations: []string{"5", "D", "+", "C"},
			expected:   15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := BaseballGame(tt.operations)
			if result != tt.expected {
				t.Errorf("BaseballGame() = %v, want %v", result, tt.expected)
			}
		})
	}
}
