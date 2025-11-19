package kclosestpointstoorigin

import (
	"container/heap"
)

type MaxHeap [][]int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosestPointsToOrigin(k int, points [][]int) [][]int {

	maxHeap := MaxHeap{}

	heap.Init(&maxHeap)

	for _, pointPair := range points {
		heap.Push(&maxHeap, pointPair)

		if maxHeap.Len() > k {
			heap.Pop(&maxHeap)
		}
	}

	return maxHeap

}
