package kthlargestelementinastream

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MinHeap) Peek() int {
	return h[0]
}

type KthLargest struct {
	minHeap MinHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		minHeap: MinHeap{},
		k:       k,
	}
	heap.Init(&kl.minHeap)

	for _, num := range nums {
		heap.Push(&kl.minHeap, num)
		if kl.minHeap.Len() > k {
			heap.Pop(&kl.minHeap)
		}
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.minHeap, val)
	if kl.minHeap.Len() > kl.k {
		heap.Pop(&kl.minHeap)
	}
	return kl.minHeap.Peek()
}
