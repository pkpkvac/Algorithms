package permutationsinastring

import (
	"testing"
)

func TestPermutationsInAString(t *testing.T) {

	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "leetcode example 1",
			s1:       "abc",
			s2:       "lecabee",
			expected: true,
		},
		{
			name:     "leetcode example 2",
			s1:       "abc",
			s2:       "lecaabee",
			expected: false,
		},
	}

	// Test input validation separately
	// validationTests := []struct {
	// 	name     string
	// 	s        string
	// 	k        int
	// 	expected int
	// }{
	// 	{
	// 		name:     "empty string",
	// 		s:        "",
	// 		k:        0,
	// 		expected: 0,
	// 	},
	// }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkInclusion(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("checkInclusion() = %v, want %v", result, tt.expected)
			}
		})
	}

	// Test input validation
	// for _, tt := range validationTests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// result, _ := LongestRepeatingCharacterReplacement(tt.s)
	// if err != nil {
	// 	t.Errorf("ValidPalindromeII() error: %v", err)
	// }
	// if result != tt.expected {
	// 	t.Errorf("LongestRepeatingCharacterReplacement() = %v, want %v", result, tt.expected)
	// }
	// })
	// }

}

// Benchmark tests for performance comparison
// func BenchmarkLongestRepeatingCharacterReplacement_Small(b *testing.B) {
// 	testCases := []string{
// 		"racecar",  // Already palindrome
// 		"abc",      // Not a palindrome
// 		"racecarr", // Needs one deletion
// 		"aba",      // Short palindrome
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		for _, s := range testCases {
// 			ValidPalindromeII(s) // Ignore error for benchmark
// 		}
// 	}
// }

// func BenchmarkValidPalindromeII_Medium(b *testing.B) {
// 	testCases := []string{
// 		"raceacaraba", // Medium sized with one char to remove
// 		"level",       // Palindrome
// 		"deified",     // Palindrome
// 		"rotator",     // Palindrome
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		for _, s := range testCases {
// 			ValidPalindromeII(s) // Ignore error for benchmark
// 		}
// 	}
// }

// func BenchmarkValidPalindromeII_Large(b *testing.B) {
// 	// Create a large palindrome-like string
// 	base := "abcdefghijklmnopqrstuvwxyz"
// 	testString := base + "x" + reverse(base)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		ValidPalindromeII(testString) // Ignore error for benchmark
// 	}
// }

// // Helper function to reverse a string
// func reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }
