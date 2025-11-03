package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	left := 0
	right := 0
	intMap := make(map[int]bool) // Initialize map with make()

	for right < len(nums) {
		// Shrink window if it exceeds k
		if right-left > k {
			// Remove leftmost element from window
			delete(intMap, nums[left])
			left++
		}

		// Check if current element already exists in window
		if intMap[nums[right]] {
			return true
		}

		// Add current element to window
		intMap[nums[right]] = true
		right++
	}

	return false
}
