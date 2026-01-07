package minimumsize

func minSubArrayLen(target int, nums []int) int {

	currSum := 0
	l := 0
	r := 0
	minLen := 100001
	for r < len(nums) {

		currSum += nums[r]
		if currSum >= target {
			minLen = min(r-l+1, minLen)
			currSum -= nums[l]
			if l < r {
				currSum -= nums[r]
				// will re-add r.
				l++
			} else {
				r++
			}
		} else {
			r++
		}

	}

	if minLen == 100001 {
		return 0
	} else {
		return minLen
	}

}

func minSubArrayLenCleaner(target int, nums []int) int {

	currSum := 0
	l := 0
	minLen := 100001

	for r := 0; r < len(nums); r++ {
		currSum += nums[r] // Always expand right first

		// Shrink left while condition is met
		for currSum >= target {
			minLen = min(r-l+1, minLen)
			currSum -= nums[l]
			l++
		}
	}
	if minLen == 100001 {
		return 0
	} else {
		return minLen
	}

}
