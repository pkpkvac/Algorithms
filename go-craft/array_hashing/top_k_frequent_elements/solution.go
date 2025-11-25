package topkfrequentelements

import (
	"container/heap"
	"sort"
)

type PairHeap []Pair

type Pair struct {
	num       int
	frequency int
}

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequentHeap(nums []int, k int) []int {

	res := []int{}

	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	h := &PairHeap{}
	// will maintain heap properties
	heap.Init(h)

	for num, frequency := range m {

		heap.Push(h, Pair{num: num, frequency: frequency})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	for h.Len() > 0 {

		el := heap.Pop(h).(Pair)

		res = append(res, el.num)
	}

	return res

}

func topKFrequentBucketSort(nums []int, k int) []int {

	// use bucket sort, the trick is
	// the bucket is an array of items
	// that contain that frequency
	// we know an element cannot occur more
	// than len(nums) times

	// first use a map to store the frequencies
	// then group the entries based on frequencies
	res := []int{}
	m := map[int]int{}
	n := len(nums)

	for _, val := range nums {
		m[val]++
	}

	// create the buckets we're going to group into.
	// 1 indexed
	b := make([][]int, len(nums)+1)

	// populate buckets
	for key, value := range m {
		bucket := b[value]
		bucket = append(bucket, key)
		b[value] = bucket
	}

	for n > 0 && len(res) < k {

		bucket := b[n]

		for _, val := range bucket {

			if len(res) == k {
				break
			}
			res = append(res, val)

		}

		n--
	}

	sort.Slice(res,
		func(i, j int) bool {
			return res[i] < res[j]
		})

	return res

}
