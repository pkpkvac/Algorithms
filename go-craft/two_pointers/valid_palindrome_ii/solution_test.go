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
			name:     "is a palindrome",
			s:        "racecar",
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
			name:     "palindrome 2 bad character",
			s:        "bracecarr",
			expected: false,
		},
	}

	// Test input validation separately
	validationTests := []struct {
		name     string
		s        string
		expected bool
	}{
		// Constraints:
		// 1 <= s.length <= 100,000
		// s is made up of only lowercase English letters.
		{
			name:     "empty input",
			s:        "",
			expected: false,
		},
		{
			name:     "bad input with spaces",
			s:        " ra ce c a r ",
			expected: false,
		},
		{
			name:     "input with non lowercase inputs (numbers)",
			s:        "racecar1",
			expected: false,
		},
		{
			name:     "input with non lowercase inputs (capitalized)",
			s:        "Racecar",
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
func BenchmarkValidPalindromeII_Current(b *testing.B) {
	// Create test data
	// names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Henry"}
	// heights := []int{170, 180, 165, 175, 160, 185, 155, 190}

	// b.ResetTimer()
	// for i := 0; i < b.N; i++ {
	// 	// Create copies to avoid modifying original data
	// 	testNames := make([]string, len(names))
	// 	testHeights := make([]int, len(heights))
	// 	copy(testNames, names)
	// 	copy(testHeights, heights)

	// 	ValidPalindromeII(testNames, testHeights) // Ignore error for benchmark
	// }
}

// // Benchmark for the optimized implementation
// func BenchmarkSortPeople_Optimized(b *testing.B) {
// 	// Create test data
// 	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Henry"}
// 	heights := []int{170, 180, 165, 175, 160, 185, 155, 190}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		// Create copies to avoid modifying original data
// 		testNames := make([]string, len(names))
// 		testHeights := make([]int, len(heights))
// 		copy(testNames, names)
// 		copy(testHeights, heights)

// 		SortPeopleOptimized(testNames, testHeights) // Ignore error for benchmark
// 	}
// }

// Benchmark with larger dataset
// func BenchmarkValidPalindromeII_LargeDataset(b *testing.B) {
// 	// Create larger test data
// 	names := make([]string, 1000)
// 	heights := make([]int, 1000)
// 	for i := 0; i < 1000; i++ {
// 		names[i] = fmt.Sprintf("Person%d", i)
// 		heights[i] = 150 + (i % 50) // Heights from 150 to 199
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		testNames := make([]string, len(names))
// 		testHeights := make([]int, len(heights))
// 		copy(testNames, names)
// 		copy(testHeights, heights)

// 		SortPeopleOptimized(testNames, testHeights) // Ignore error for benchmark
// 	}
// }
