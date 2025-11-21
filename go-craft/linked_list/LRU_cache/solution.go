package lrucache

import (
	"algorithms/linked_list/common"
)

// this is more performant than the hash map implementation
// this has O(1) for put and get

type LRUCache struct {
	capacity int
	cache    map[int]int
	nodes    map[int]*common.DoublyLinkedNode
	head     *common.DoublyLinkedNode // dummy head (MRU side)
	tail     *common.DoublyLinkedNode // dummy tail (LRU side)
}

func Constructor(capacity int) LRUCache {
	// Create dummy head and tail nodes, connected together
	head := &common.DoublyLinkedNode{}
	tail := &common.DoublyLinkedNode{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		capacity: capacity,
		cache:    map[int]int{},
		nodes:    map[int]*common.DoublyLinkedNode{},
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	value, ok := l.cache[key]
	if !ok {
		return -1
	}

	// Move node to head (MRU position)
	l.moveToHead(key)
	return value
}

func (l *LRUCache) Put(key int, value int) {
	_, exists := l.cache[key]

	if exists {
		// Update existing: update value and move to head
		l.cache[key] = value
		l.moveToHead(key)
		return
	}

	// New node case
	// Check if we need to evict BEFORE adding
	if len(l.nodes) >= l.capacity {
		l.evictLRU()
	}

	// Create and add new node
	newNode := common.NewDoublyLinkedNode(key)
	l.addToHead(newNode)
	l.nodes[key] = newNode
	l.cache[key] = value
}

// moveToHead moves an existing node to the head (MRU position)
func (l *LRUCache) moveToHead(key int) {
	node := l.nodes[key]
	if node == nil {
		return
	}

	// Remove node from its current position
	l.removeNode(node)

	// Add to head
	l.addToHead(node)
}

// removeNode removes a node from the list
func (l *LRUCache) removeNode(node *common.DoublyLinkedNode) {
	if node == nil {
		return
	}
	// Update neighbors
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
}

// addToHead adds a node right after the dummy head (MRU position)
func (l *LRUCache) addToHead(node *common.DoublyLinkedNode) {
	if node == nil {
		return
	}
	// Insert after head
	node.Prev = l.head
	node.Next = l.head.Next
	l.head.Next.Prev = node
	l.head.Next = node
}

// evictLRU removes the least recently used node (before tail)
func (l *LRUCache) evictLRU() {
	// LRU is the node right before tail
	lruNode := l.tail.Prev
	if lruNode == l.head {
		// List is empty (only dummy nodes)
		return
	}

	// Remove from list
	l.removeNode(lruNode)

	// Remove from maps
	lruKey := lruNode.Val
	delete(l.cache, lruKey)
	delete(l.nodes, lruKey)
}
