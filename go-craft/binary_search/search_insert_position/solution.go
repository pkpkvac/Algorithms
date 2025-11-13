package searchinsertposition

func searchInsert(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	for l <= r {

		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			// should be in the left
			// move r to mid - 1
			r = mid - 1
		}
		if nums[mid] < target {
			// should be in the right
			// move l to mid + 1
			l = mid + 1
		}

	}

	return l

}
