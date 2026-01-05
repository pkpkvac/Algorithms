package trappingrainwater

func trapBruteForce(height []int) int {

	sumWater := 0
	for i, h := range height {
		// h is height at i

		leftMax := 0
		rightMax := 0

		// calculate the water at height at i (h)

		// we need to know the max height to the left of i, to the right of i
		// the minimum of these two, and the height at i tell us the water units
		// present at i
		for j := i; j > 0; j-- {

			leftMax = max(leftMax, height[j])

		}

		for j := i; j < len(height); j++ {

			rightMax = max(rightMax, height[j])

		}

		// take the minimum to get the highest water level
		waterLevel := min(rightMax, leftMax)

		// this is the waterVolume at the index, add it to the sum
		waterVolume := waterLevel - h

		sumWater += waterVolume

	}

	return sumWater

}

func trap(height []int) int {
	// to prevent recalculating/finding the
	// maximum height at a given index
	// use a prefix maximum leftMax[i] - max from 0 to i-1
	// and a suffix maximum rightMax[i] - max from i+1 to n-1

	waterSum := 0

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	l_h := 0
	for i := range height {
		// prefix maximum left
		leftMax[i] = l_h
		l_h = max(l_h, height[i])
	}

	r_h := 0
	for i := len(height) - 1; i >= 0; i-- {
		rightMax[i] = r_h
		r_h = max(r_h, height[i])
	}

	for i, h := range height {

		waterVolume := min(leftMax[i], rightMax[i]) - h

		if waterVolume > 0 {
			waterSum += waterVolume
		}

	}

	return waterSum
}
