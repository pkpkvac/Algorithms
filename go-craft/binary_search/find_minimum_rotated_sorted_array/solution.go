package find_minimum_rotated_sorted_array

func findMin(nums []int) int {

	l := 0
	r := len(nums) - 1
	minVal := nums[0]

	for l < r {
		mid := (l + r) / 2
		minVal = min(minVal, nums[mid])

		if nums[l] > nums[r] {
			if nums[l] < nums[mid] {
				// inflection point is at or beyond mid
				l = mid + 1
			} else {
				// inflection point is between l and m
				r = mid
			}

		} else if nums[l] < nums[r] {
			// seems like usual scenario, move r = mid
			r = mid - 1
		}

	}

	return minVal

}

func findMinProper(nums []int) int {

	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > nums[r] {
			// right half unsorted, search right
			l = mid + 1
		} else {
			r = mid
		}

	}

	return nums[r]
}

// func BinarySearch(nums []int, target int) int {

// 	l := 0
// 	r := len(nums) - 1

// 	// fmt.Println("target ", target)

// 	for l <= r {

// 		mid := l + (r-l)/2

// 		// fmt.Println("mid ", mid, " nums[mid] ", nums[mid])

// 		if nums[mid] == target {
// 			return mid
// 		}
// 		if nums[mid] > target {
// 			// in the left
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}

// 	}

// 	return -1
// }
