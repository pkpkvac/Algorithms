package containermostwater

func maxArea(heights []int) int {

	l := 0
	r := len(heights) - 1
	mA := 0

	for l < r {

		length := r - l
		h := min(heights[l], heights[r])
		a := length * h

		if heights[l] > heights[r] {
			// no need to forward l
			r--
		} else {
			l++
		}

		mA = max(mA, a)

	}

	return mA
}
