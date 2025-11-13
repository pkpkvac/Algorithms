package twosumiisorted

func TwoSumIISorted(numbers []int, target int) []int {

	l := 0
	r := len(numbers) - 1

	for l < r {

		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum > target {
			// move right down
			r--
		} else if sum < target {
			l++
		}

	}
	return []int{}
}
