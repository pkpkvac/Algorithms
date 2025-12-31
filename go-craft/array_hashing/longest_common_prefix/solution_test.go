package longestcommonprefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	t.Run("leetcode example 1", func(t *testing.T) {
		strs := []string{"bat", "bag", "bank", "band"}
		expected := "ba"
		result := longestCommonPrefix(strs)
		if result != expected {
			t.Errorf("longestCommonPrefix(%v) = %v, want %v", strs, result, expected)
		}

		strs = []string{"dance", "dag", "danger", "damage"}
		expected = "da"
		result = longestCommonPrefix(strs)
		if result != expected {
			t.Errorf("longestCommonPrefix(%v) = %v, want %v", strs, result, expected)
		}

		strs = []string{"neet", "feet"}
		expected = ""
		result = longestCommonPrefix(strs)
		if result != expected {
			t.Errorf("longestCommonPrefix(%v) = %v, want %v", strs, result, expected)
		}
	})
}
