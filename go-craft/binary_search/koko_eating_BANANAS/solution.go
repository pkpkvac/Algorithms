package kokoeatingbananas

func minEatingSpeed(piles []int, h int) int {

	maxSpeed := 0
	minSpeed := 1
	for i := range piles {
		maxSpeed = max(maxSpeed, piles[i])
	}

	for maxSpeed >= minSpeed {

		// this loop should eventually converge to the solution
		speed := (maxSpeed + minSpeed) / 2

		totalH := 0

		for _, pile := range piles {
			// ceil(a / b) == (a + b - 1) / b
			// This exact pattern shows up in:
			// batching requests
			// chunking files
			// rate limiting
			// pagination
			// worker pool sizing
			// infra capacity planning
			// If you remember one rule, remember this:
			// If you need “how many chunks of size k to cover n items” → use (n + k - 1) / k.
			totalH += (pile + speed - 1) / speed
		}

		if totalH <= h {
			// this is a valid speed, but we could go... SLOWER
			// set the maxSpeed = speed
			// res = min(speed, maxSpeed)
			maxSpeed = speed - 1
		} else {
			// totalH > h, which means we ate too slow,
			// and need to boost minSpeed, eat faster
			minSpeed = speed + 1
		}

	}
	return minSpeed
}
