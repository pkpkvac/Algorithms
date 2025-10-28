package validpalindromeii

import (
	"reflect"
	"testing"
)

func TestValidPalindromeII(t *testing.T) {

	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "empty string",
			s:        "",
			expected: false,
		},
		{
			name:     "single character",
			s:        "a",
			expected: true,
		},
		{
			name:     "two identical characters",
			s:        "aa",
			expected: true,
		},
		{
			name:     "is a palindrome",
			s:        "racecar",
			expected: true,
		},
		{
			name:     "already a palindrome",
			s:        "aba",
			expected: true,
		},
		{
			name:     "not a palindrome",
			s:        "abc",
			expected: false,
		},
		{
			name:     "palindrome 1 bad character",
			s:        "racecarr",
			expected: true,
		},
		{
			name:     "can become palindrome by removing one char (remove right)",
			s:        "racecarr", // Remove last 'r' -> "racecar"
			expected: true,
		},
		{
			name:     "can become palindrome by removing one char (remove left)",
			s:        "raceacar", // Remove 'a' -> "racecar"
			expected: true,
		},
		{
			name:     "cannot become palindrome",
			s:        "race",
			expected: false,
		},
	}

	// Test input validation separately
	validationTests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "input with non lowercase inputs (capitalized)",
			s:        "Racecar",
			expected: false,
		},
		{
			name:     "input with numbers",
			s:        "racecar1",
			expected: false,
		},
		{
			name:     "input with uppercase",
			s:        "RACECAR",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := ValidPalindromeII(tt.s)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ValidPalindromeII() = %v, want %v", result, tt.expected)
			}
		})
	}

	// Test input validation
	for _, tt := range validationTests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := ValidPalindromeII(tt.s)
			// if err != nil {
			// 	t.Errorf("ValidPalindromeII() error: %v", err)
			// }
			if result != tt.expected {
				t.Errorf("ValidPalindromeII() = %v, want %v", result, tt.expected)
			}
		})
	}

}

// Benchmark tests for performance comparison
func BenchmarkValidPalindromeII_Small(b *testing.B) {
	testCases := []string{
		"racecar",  // Already palindrome
		"abc",      // Not a palindrome
		"racecarr", // Needs one deletion
		"aba",      // Short palindrome
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testCases {
			ValidPalindromeII(s) // Ignore error for benchmark
		}
	}
}

func BenchmarkValidPalindromeII_Medium(b *testing.B) {
	testCases := []string{
		"raceacaraba", // Medium sized with one char to remove
		"level",       // Palindrome
		"deified",     // Palindrome
		"rotator",     // Palindrome
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testCases {
			ValidPalindromeII(s) // Ignore error for benchmark
		}
	}
}

func BenchmarkValidPalindromeII_Large(b *testing.B) {
	// Create a large palindrome-like string
	base := "abcdefghijklmnopqrstuvwxyz"
	testString := base + "x" + reverse(base)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidPalindromeII(testString) // Ignore error for benchmark
	}
}

// Helper function to reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
