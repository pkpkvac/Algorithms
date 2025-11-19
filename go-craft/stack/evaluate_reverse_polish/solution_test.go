package evaluaterpn

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {

	tests := []struct {
		name     string
		tokens   []string
		expected int
	}{
		{
			name:     "leetcode example 1",
			tokens:   []string{"1", "2", "+", "3", "*", "4", "-"},
			expected: 5,
		},
		{
			name:     "leetcode example 2",
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalRPN(tt.tokens)
			if result != tt.expected {
				t.Errorf("evalRPN() = %v, want %v", result, tt.expected)
			}
		})
	}
}
