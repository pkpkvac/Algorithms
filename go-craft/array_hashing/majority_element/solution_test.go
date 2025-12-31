package majorityelement

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {

	t.Run("leetcode example 1", func(t *testing.T) {
		nums := []int{5, 5, 1, 1, 1, 5, 5}
		expected := 5
		result := majorityElementHashMap(nums)
		if result != expected {
			t.Errorf("majorityElement(%v) = %v, want %v", nums, result, expected)
		}

		nums = []int{2, 2, 2}
		expected = 2
		result = majorityElementHashMap(nums)
		if result != expected {
			t.Errorf("majorityElement(%v) = %v, want %v", nums, result, expected)
		}
	})
}
