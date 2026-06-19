package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] = length of LIS ending at index i
	dp := make([]int, len(nums))

	// base case
	// every element is a subsequence of length 1 (to start)
	for i := range dp {
		dp[i] = 1
	}

	// fill dp array
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	maxLen := 0
	for _, v := range dp {
		maxLen = max(maxLen, v)
	}

	return maxLen
}
