package laststoneweight

import (
	"container/heap"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MaxHeap) Peek() int {
	return h[0]
}

func lastStoneWeight(stones []int) int {

	h := &MaxHeap{}

	heap.Init(h)

	for _, stone := range stones {
		heap.Push(h, stone)
	}

	// fmt.Println(h)
	// fmt.Println(*h)

	for len(*h) > 1 {

		stone1 := heap.Pop(h).(int)
		stone2 := heap.Pop(h).(int)

		rem := stone1 - stone2

		if rem > 0 {
			heap.Push(h, rem)
		}

	}

	if len(*h) == 1 {
		return heap.Pop(h).(int)
	}

	return 0

}
