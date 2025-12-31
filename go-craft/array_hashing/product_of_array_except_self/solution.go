package productofarrayexceptself

func productExceptSelfBruteForce(nums []int) []int {
	sol := []int{}
	undef := -33
	for i := range nums {
		sum := undef
		for j, val := range nums {
			if i == j {
				continue
			}
			if val == 0 {
				sum = 0
				break
			}
			if sum == undef {
				sum = val
			} else {
				sum *= val
			}
		}
		sol = append(sol, sum)
	}

	return sol
}

func productExceptSelf(nums []int) []int {

	prefixProd := make([]int, len(nums))
	suffixProd := make([]int, len(nums))
	result := make([]int, len(nums))

	// product of prefix to left of i, not including self
	prod := 1
	for i := range nums {

		prefixProd[i] = prod

		prod *= nums[i]

	}

	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {

		suffixProd[i] = prod

		prod *= nums[i]

	}

	for i := range nums {
		result[i] = suffixProd[i] * prefixProd[i]
	}

	return result

}
