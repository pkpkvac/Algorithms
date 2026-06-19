package subsets_ii

import "sort"

func subsetsWithDup(nums []int) [][]int {

	// critical to sort first to group the duplicates
	sort.Ints(nums)
	
	sol := [][]int{}
	currSet := []int{}

	backtrack(nums, 0, &currSet, &sol)

	return sol
}

func backtrack(n []int, start int, cs *[]int, solution *[][]int) {
	// Add current subset to solution
	*solution = append(*solution, append([]int{}, *cs...))

	// Try adding each remaining element
	for i := start; i < len(n); i++ {
		// Skip duplicates: if current element == previous element
		// AND we're not at the start position, skip it
		if i > start && n[i] == n[i-1] {
			continue
		}

		// Choose: add n[i] to current subset
		*cs = append(*cs, n[i])
		
		// Explore: recurse with next index
		backtrack(n, i+1, cs, solution)
		
		// Unchoose: remove n[i] (backtrack)
		*cs = (*cs)[:len(*cs)-1]
	}
}