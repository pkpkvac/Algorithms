package remove_element

// RemoveElement removes all instances of val from nums in-place.
// Returns the new length of the array after removal.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	// Two pointers approach
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
