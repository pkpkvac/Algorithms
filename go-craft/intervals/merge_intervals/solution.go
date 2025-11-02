package mergeintervals

import "sort"

type Interval [][]int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a Interval) Len() int {

	return len(a)
}

func (a Interval) Swap(i, j int) {

	a[i], a[j] = a[j], a[i]

}

func (a Interval) Less(i, j int) bool {

	return a[i][0] < a[j][0]

}

func merge(intervals [][]int) [][]int {

	// sort intervals
	sort.Sort(Interval(intervals))

	// sort.Slice(intervals, func(i, j int) bool {
	// 	return intervals[i][0] < intervals[j][0]
	// })
	if len(intervals) == 0 {
		return intervals
	}

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := result[len(result)-1]

		if lastMerged[1] >= current[0] {
			// Overlap! Merge by updating the last interval's end time
			result[len(result)-1][1] = max(lastMerged[1], current[1])
		} else {
			// No overlap! Add new interval
			result = append(result, current)
		}
	}

	return result

}
