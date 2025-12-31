package majorityelement

// func majorityElement(nums []int) int {

// 	// linear time, constant space

// }

func majorityElementHashMap(nums []int) int {

	majority := len(nums) / 2
	m := make(map[int]int)

	for _, val := range nums {
		m[val]++
	}

	for key, value := range m {

		if value > majority {
			return key
		}
	}

	return -1

}
