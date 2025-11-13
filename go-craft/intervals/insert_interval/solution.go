package insertinterval

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i := 0
	n := len(intervals)

	// Add all intervals that end before newInterval starts
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// Merge all overlapping intervals
	merged := []int{newInterval[0], newInterval[1]}
	for i < n && intervals[i][0] <= merged[1] {
		merged[0] = min(merged[0], intervals[i][0])
		merged[1] = max(merged[1], intervals[i][1])
		i++
	}
	res = append(res, merged)

	// Add all remaining intervals
	for i < n {
		res = append(res, intervals[i])
		i++
	}

	return res
}
