package nonoverlappingintervals

import (
	"sort"
)

type Interval [][]int

func (a Interval) Len() int {

	return len(a)
}

func (a Interval) Swap(i, j int) {

	a[i], a[j] = a[j], a[i]

}

func (a Interval) Less(i, j int) bool {

	return a[i][0] < a[j][0]

}

func eraseOverlapIntervals(intervals [][]int) int {
	// this is a greedy algorithm

	sort.Sort(Interval(intervals))

	if len(intervals) == 0 {
		return 0
	}

	// only need the last ending time
	// draw out why this is the case
	// (because starting times will not conflict)
	lastEnd := intervals[0][1]
	numErase := 0

	for i := 1; i < len(intervals); i++ {

		if intervals[i][0] < lastEnd {
			// case of an interval that is to be discarded
			// greedily select the earliest ending time
			// the maximize chance of next interval
			// fitting in the set
			if intervals[i][1] < lastEnd {
				lastEnd = intervals[i][1]
			}

			numErase++

		} else {
			// the set now has this new interval
			// just need to track its end
			lastEnd = intervals[i][1]
		}

	}

	return numErase

}
