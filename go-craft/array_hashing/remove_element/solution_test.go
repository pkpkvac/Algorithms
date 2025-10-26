package remove_element

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
		wantNums []int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			val:      3,
			expected: 0,
			wantNums: []int{},
		},
		{
			name:     "single element to remove",
			nums:     []int{3},
			val:      3,
			expected: 0,
			wantNums: []int{},
		},
		{
			name:     "single element not to remove",
			nums:     []int{1},
			val:      3,
			expected: 1,
			wantNums: []int{1},
		},
		{
			name:     "multiple elements, some to remove",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			wantNums: []int{2, 2},
		},
		{
			name:     "all elements to remove",
			nums:     []int{3, 3, 3, 3},
			val:      3,
			expected: 0,
			wantNums: []int{},
		},
		{
			name:     "no elements to remove",
			nums:     []int{1, 2, 3, 4},
			val:      5,
			expected: 4,
			wantNums: []int{1, 2, 3, 4},
		},
		{
			name:     "leetcode example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			wantNums: []int{2, 2},
		},
		{
			name:     "leetcode example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input slice to avoid modifying the original
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			got := RemoveElement(nums, tt.val)

			if got != tt.expected {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.expected)
			}

			// Check that the first 'got' elements match the expected array
			if got > 0 && !reflect.DeepEqual(nums[:got], tt.wantNums) {
				t.Errorf("RemoveElement() modified array = %v, want %v", nums[:got], tt.wantNums)
			}
		})
	}
}

// Benchmark tests for performance
func BenchmarkRemoveElement(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i % 3 // Creates pattern: 0,1,2,0,1,2...
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		RemoveElement(testNums, 1)
	}
}
