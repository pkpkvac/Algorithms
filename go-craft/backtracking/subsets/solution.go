package subsets

func subsets(nums []int) [][]int {

	sol := [][]int{}
	currSet := []int{}

	backtrack(nums, 0, &currSet, &sol)

	return sol
}

func backtrack(n []int, index int, cs *[]int, solution *[][]int) {

	// base case
	if index == len(n) {
		*solution = append(*solution, append([]int{}, *cs...))
		return
	}

	// choice 1: include nums[index]
	*cs = append(*cs, n[index])

	backtrack(n, index+1, cs, solution)

	// choice 2: exclude nums[index]
	*cs = (*cs)[:len(*cs)-1]

	backtrack(n, index+1, cs, solution)

}
