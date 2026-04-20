package max_subarray

import "math"

func maxSubArray(nums []int) int {

	// At each step you only decide:
	// extend the current run or start
	// at this element. One pass, two variables.

	// keep the best so far as a final return value

	// Kadane's algortihm
	bestHere := math.MinInt32
	bestSoFar := math.MinInt32

	for _, num := range nums {

		bestHere = max(num, bestHere+num)

		bestSoFar = max(bestSoFar, bestHere)

	}
	return bestSoFar
}
