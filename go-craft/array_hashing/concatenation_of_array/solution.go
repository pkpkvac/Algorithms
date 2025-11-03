package concatenationofarray

func getConcatenation(nums []int) []int {
	ans := make([]int, 0, len(nums)*2) // Pre-allocate capacity, empty non-nil slice

	for i := 0; i < 2; i++ {
		ans = append(ans, nums...)
	}

	return ans
}

func getConcatenation2(nums []int) []int {
	ans := make([]int, len(nums)*2) // Pre-allocate capacity, empty non-nil slice

	copy(ans[0:len(nums)], nums)

	copy(ans[:len(nums)], nums)

	return ans
}
