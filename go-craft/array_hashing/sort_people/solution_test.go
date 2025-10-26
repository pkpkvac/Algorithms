package sortpeople

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortPeople(t *testing.T) {

	tests := []struct {
		name          string
		names         []string
		heights       []int
		expectedNames []string
	}{
		{
			name:          "empty",
			names:         []string{},
			heights:       []int{},
			expectedNames: []string{},
		},
		{
			name:          "Proper heights",
			names:         []string{"Homer", "Achilles", "Hector"},
			heights:       []int{180, 200, 190},
			expectedNames: []string{"Achilles", "Hector", "Homer"},
		},
		{
			name:          "Single person",
			names:         []string{"Alice"},
			heights:       []int{170},
			expectedNames: []string{"Alice"},
		},
		{
			name:          "Same heights",
			names:         []string{"Alice", "Bob"},
			heights:       []int{170, 170},
			expectedNames: []string{"Alice", "Bob"}, // Should maintain original order
		},
	}

	// Test input validation separately
	validationTests := []struct {
		name        string
		names       []string
		heights     []int
		shouldPanic bool
	}{
		{
			name:        "Mismatched lengths - more names",
			names:       []string{"Alice", "Bob", "Charlie"},
			heights:     []int{170, 180},
			shouldPanic: true,
		},
		{
			name:        "Mismatched lengths - more heights",
			names:       []string{"Alice", "Bob"},
			heights:     []int{170, 180, 190},
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SortPeople(tt.names, tt.heights)
			if err != nil {
				t.Errorf("SortPeople() unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tt.expectedNames) {
				t.Errorf("SortPeople() = %v, want %v", result, tt.expectedNames)
			}
		})
	}

	// Test input validation
	for _, tt := range validationTests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SortPeople(tt.names, tt.heights)
			if tt.shouldPanic {
				// Now we expect an error, not a panic
				if err == nil {
					t.Errorf("Expected error for mismatched lengths, but got none")
				}
				if result != nil {
					t.Errorf("Expected nil result for error case, got %v", result)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}

}

// Benchmark tests for performance comparison
func BenchmarkSortPeople_Current(b *testing.B) {
	// Create test data
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Henry"}
	heights := []int{170, 180, 165, 175, 160, 185, 155, 190}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create copies to avoid modifying original data
		testNames := make([]string, len(names))
		testHeights := make([]int, len(heights))
		copy(testNames, names)
		copy(testHeights, heights)

		SortPeople(testNames, testHeights) // Ignore error for benchmark
	}
}

// Benchmark for the optimized implementation
func BenchmarkSortPeople_Optimized(b *testing.B) {
	// Create test data
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Henry"}
	heights := []int{170, 180, 165, 175, 160, 185, 155, 190}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create copies to avoid modifying original data
		testNames := make([]string, len(names))
		testHeights := make([]int, len(heights))
		copy(testNames, names)
		copy(testHeights, heights)

		SortPeopleOptimized(testNames, testHeights) // Ignore error for benchmark
	}
}

// Benchmark with larger dataset
func BenchmarkSortPeople_LargeDataset(b *testing.B) {
	// Create larger test data
	names := make([]string, 1000)
	heights := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		names[i] = fmt.Sprintf("Person%d", i)
		heights[i] = 150 + (i % 50) // Heights from 150 to 199
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testNames := make([]string, len(names))
		testHeights := make([]int, len(heights))
		copy(testNames, names)
		copy(testHeights, heights)

		SortPeopleOptimized(testNames, testHeights) // Ignore error for benchmark
	}
}
