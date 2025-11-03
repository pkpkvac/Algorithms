package containsduplicateii

import (
	"reflect"
	"testing"
)

func TestContainsDuplicateII(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		k        int
		expected bool
	}{
		{
			name:     "empty string",
			nums:     []int{},
			k:        0,
			expected: false,
		},
		{
			name:     "true case ",
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			name:     "false case",
			nums:     []int{2, 1, 2},
			k:        1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsNearbyDuplicate(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ContainsDuplicateII() = %v, want %v", result, tt.expected)
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
