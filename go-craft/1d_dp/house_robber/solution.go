package house_robber

// this problem differs slightly from climbing stairs, but
// in an interesting way
// where the indexing can be very dynamic
// due to the recursive call stack paths

func rob(nums []int) int {

	// m := make(map[int]int)

	res := robHelper(nums, 0)
	// res := robHelperWithMemoization(nums, m, 0)
	// res := robTabulation(nums)

	return res
}

func robHelper(nums []int, idx int) int {

	if idx >= len(nums) {
		// base case
		// out of range
		return 0
	}

	// two COMPETING options, whose subtrees
	// need to be explored independently, taking the max
	// 1. take the current index, but that means
	// you cannot take the neighbor
	// 2. skip the current index, and make the same decision as 1 and 2 on the
	// next step.
	return max(nums[idx]+robHelper(nums, idx+2), robHelper(nums, idx+1))

}

func robHelperWithMemoization(nums []int, m map[int]int, idx int) int {

	if idx >= len(nums) {
		// base case
		// out of range
		return 0
	}

	if value, ok := m[idx]; ok {
		return value
	}

	res := max(nums[idx]+robHelperWithMemoization(nums, m, idx+2), robHelperWithMemoization(nums, m, idx+1))

	m[idx] = res

	return res

}

// tabulation is trickier here, the table is filled in one direction, so never recurse.
// 1. need to define what dp[i] means. dp[i] is the maximum amount that can be robbed from
// houses 0 through i

// e.g for [2,9,8,3,6]
// dp[0] = 2, dp[1] = 9, dp[2] = 2 + 8 = 10,
// dp[i] = max(nums[i] + dp[i-2], dp[i-1])
// dp[2] = max(8 + 2, 9) = 10

func robTabulation(nums []int) int {

	dp := make(map[int]int)
	dp[0] = 2
	dp[1] = 9

	if len(nums) <= 1 {
		return dp[len(nums)]
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}

	return dp[len(nums)-1]

}
