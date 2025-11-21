package lrucache

import "slices"

// start first with hash map implementation
// the issue is this does not have O(1) for put and get,

type LRUCache struct {
	capacity    int
	m           map[int]int
	accessOrder []int
}

func Constructor(capacity int) LRUCache {

	lru := LRUCache{capacity: capacity, m: map[int]int{}, accessOrder: []int{}}

	return lru

}

func (l *LRUCache) Get(key int) int {

	// if you call get,
	// this needs to bump
	// the element forward in the LRU Cache

	value, ok := l.m[key]

	if !ok {
		return -1
	}

	// if key exists in the accessOrder array
	// slice it out at that index
	// then append it to the accessOrder array
	// operation costs

	l.getAndUpdateAccessOrder(key)

	return value
}

func (l *LRUCache) Put(key int, value int) {

	// determine if exists already
	_, exists := l.m[key]

	// assign the value
	l.m[key] = value

	if exists {
		// we do not need to remove LRU
		l.getAndUpdateAccessOrder(key)
		return
	}

	// put key as MRU in accessOrder at the end if DNE
	l.accessOrder = append(l.accessOrder, key)

	if len(l.accessOrder) > l.capacity {
		// have to remove the LRU element if at capacity
		// also need to remove it from the map
		// and remove LRU element if at capacity
		lruKey := l.accessOrder[0]
		l.accessOrder = l.accessOrder[1:]

		// O(1)
		delete(l.m, lruKey)
	}

}

func (l *LRUCache) getAndUpdateAccessOrder(key int) {

	// O(n)
	idxKey := slices.Index(l.accessOrder, key)

	// O(n)
	l.accessOrder = slices.Delete(l.accessOrder, idxKey, idxKey+1)

	// add at end to MRU, O(1)
	l.accessOrder = append(l.accessOrder, key)

}

// func ReverseList(head *common.ListNode) *common.ListNode {

// 	if head == nil {
// 		return nil
// 	}

// 	prev := (*common.ListNode)(nil)
// 	current := head

// 	for current != nil {

// 		next := current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next

// 	}

// 	return prev
// }
