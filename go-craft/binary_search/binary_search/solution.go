package binarysearch

func BinarySearch(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	// fmt.Println("target ", target)

	for l <= r {

		mid := l + (r-l)/2

		// fmt.Println("mid ", mid, " nums[mid] ", nums[mid])

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			// in the left
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return -1
}
