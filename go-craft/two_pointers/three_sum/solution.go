package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {

	res := [][]int{}
	sort.Ints(nums)
	// for now, just gather the indices
	// of the solution, can make it
	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			// skip duplicate values
			// this result would already exist in the solution space
			continue
		}
		target := -nums[i]

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[j] + nums[k]

			if sum == target {
				// for these indexes, store the result
				var triplet = []int{nums[i], nums[j], nums[k]}
				res = append(res, triplet)

				// found a solution, let's up front
				// ensure that we do not
				// add duplicates
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum < target {
				// boost j
				j++
			} else {
				// reduce k
				k--
			}

		}

	}

	return res
}

// func TwoSumIISorted(numbers []int, target int) []int {

// 	l := 0
// 	r := len(numbers) - 1

// 	for l < r {

// 		sum := numbers[l] + numbers[r]
// 		if sum == target {
// 			return []int{l + 1, r + 1}
// 		} else if sum > target {
// 			// move right down
// 			r--
// 		} else if sum < target {
// 			l++
// 		}

// 	}
// 	return []int{}
// }
