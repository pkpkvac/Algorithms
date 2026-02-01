package combination_sum

func combinationSum(nums []int, target int) [][]int {

	sol := [][]int{}
	currSet := []int{}

	backtrack(nums, 0, target, &currSet, &sol)

	return sol

}

func backtrack(n []int, index int, remainingTarget int, cs *[]int, solution *[][]int) {

	// base cases

	if remainingTarget == 0 {
		*solution = append(*solution, append([]int{}, *cs...))
		return
	}

	if remainingTarget < 0 {
		// exceeded target -> prune last element
		return
	}

	if index == len(n) {
		return
	}

	// choice 1: include nums[index]
	*cs = append(*cs, n[index])

	backtrack(n, index, remainingTarget-n[index], cs, solution)

	// choice 2: exclude, don't add to currSet
	*cs = (*cs)[:len(*cs)-1]

	backtrack(n, index+1, remainingTarget, cs, solution)

}
