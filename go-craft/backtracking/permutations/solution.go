package permutations

func permute(nums []int) [][]int {
	result := [][]int{}
	current := []int{}
	used := make([]bool, len(nums))

	backtrack(nums, current, used, &result)
	return result
}

func backtrack(nums []int, current []int, used []bool, result *[][]int) {

	// base case (complete perm)
	if len(current) == len(nums) {
		perm := make([]int, len(current))
		copy(perm, current)
		*result = append(*result, perm)
		return
	}

	// recursive case (try each unused number)
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		// choose -> add nums[i] to current path
		current = append(current, nums[i])
		used[i] = true

		// explore -> recurse with the choice
		backtrack(nums, current, used, result)

		// unchoose -> remove nums[i] (backtrack)
		current = current[:len(current)-1]
		used[i] = false
	}

}

// The Pattern (Choose → Explore → Unchoose):
//   1. Choose: Add element to path, mark as used
//   2. Explore: Recurse deeper
//   3. Unchoose: Remove element, mark as unused (backtrack)

//   Example trace with [1,2,3]:
//   Start: current=[], used=[F,F,F]

//   Choose 1: current=[1], used=[T,F,F]
//     Choose 2: current=[1,2], used=[T,T,F]
//       Choose 3: current=[1,2,3], used=[T,T,T] → Add [1,2,3]
//       Unchoose 3: current=[1,2], used=[T,T,F]
//     Unchoose 2: current=[1], used=[T,F,F]

//     Choose 3: current=[1,3], used=[T,F,T]
//       Choose 2: current=[1,3,2], used=[T,T,T] → Add [1,3,2]
//       ...

//   The used array prevents re-using the same element in a single permutation
