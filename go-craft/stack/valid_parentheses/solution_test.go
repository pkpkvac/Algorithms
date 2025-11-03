package validparentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "true case",
			input:    "()",
			expected: true,
		},
		{
			name:     "false case",
			input:    "([)",
			expected: false,
		},
		{
			name:     "true case",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "false case",
			input:    "(]",
			expected: false,
		},
		{
			name:     "true case",
			input:    "([{}])",
			expected: true,
		},
		{
			name:     "false case",
			input:    "([{}]]",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.input)
			if result != tt.expected {
				t.Errorf("IsValid() = %v, want %v", result, tt.expected)
			}
		})
	}
}
