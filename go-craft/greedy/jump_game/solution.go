package jump_game

func canJump(nums []int) bool {

	maxReach := 0

	for i, num := range nums {

		if i > maxReach {
			return false
		}

		maxReach = max(maxReach, i+num)

		if maxReach >= len(nums)-1 {
			return true
		}

	}

	return true

}
