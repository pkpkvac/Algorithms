package sumofallsubsetsxortotals

func subsetXORTotal(nums []int) int {

	return backtrack(nums, 0, 0)

}

func backtrack(n []int, index int, currentXOR int) int {

	if index == len(n) {
		return currentXOR
	}

	total := 0

	total += backtrack(n, index+1, currentXOR^(n[index]))
	total += backtrack(n, index+1, currentXOR)

	return total

}
