package meetingroomii

type Interval struct {
	start int
	end   int
}

// this needs to be implemented
func minMeetingRoomsSweepLine(intervals []Interval) int {

}

// type MinHeapByEnd []Interval

// // This heap orders the end times to give the earliest ending meeting.
// // This heap orders the end times to give the earliest ending meeting.
// // This heap orders the end times to give the earliest ending meeting.

// func (h MinHeapByEnd) Len() int           { return len(h) }
// func (h MinHeapByEnd) Less(i, j int) bool { return h[i].end < h[j].end }
// func (h MinHeapByEnd) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// type ByStart []Interval

// func (a ByStart) Len() int           { return len(a) }
// func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByStart) Less(i, j int) bool { return a[i].start < a[j].start }

// func (h *MinHeapByEnd) Push(x any) {
// 	*h = append(*h, x.(Interval))
// }

// func (h MinHeapByEnd) Peek() Interval {
// 	return h[0]
// }

// func (h *MinHeapByEnd) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func minMeetingRoomsWithHeap(intervals []Interval) int {

// 	daysNeeded := 0

// 	if len(intervals) == 0 {
// 		return daysNeeded
// 	}

// 	sort.Sort(ByStart(intervals))

// 	h := &MinHeapByEnd{}
// 	heap.Init(h)
// 	heap.Push(h, intervals[0])
// 	daysNeeded++

// 	for i := 1; i < len(intervals); i++ {

// 		if h.Len() > 0 && intervals[i].start >= h.Peek().end {
// 			// this means the room is free, the last meeting has ended
// 			// pop the last meeting off the heap until this is no longer true

// 			for h.Len() > 0 && intervals[i].start >= h.Peek().end {

// 				heap.Pop(h)

// 			}
// 			// push the current meeting onto the heap
// 			heap.Push(h, intervals[i])

// 		} else {
// 			// this means the start of the new meeting < end of the last, which means
// 			// need to get a new room
// 			heap.Push(h, intervals[i])
// 		}

// 		// the number of days is the maximum of the number of rooms needed after this processing step
// 		daysNeeded = max(h.Len(), daysNeeded)

// 	}

// 	return daysNeeded
// }
