package meetingroom

import (
	"sort"
)

type Interval struct {
	start int
	end   int
}

type Intervals []Interval

func (a Intervals) Len() int {
	return len(a)
}

func (a Intervals) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Intervals) Less(i, j int) bool {
	return a[i].start < a[j].start
}

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

func canAttendMeetings(intervals []Interval) bool {

	sort.Sort(Intervals(intervals))

	if len(intervals) == 0 {
		return true
	}

	prevMeeting := intervals[0].end

	for i := 1; i < len(intervals); i++ {
		if prevMeeting > intervals[i].start {
			return false
		}
		prevMeeting = intervals[i].end
	}

	return true

}
