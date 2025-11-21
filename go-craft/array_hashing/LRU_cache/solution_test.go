package lrucache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {

	t.Run("leetcode example 1", func(t *testing.T) {
		// LRUCache lRUCache = new LRUCache(2);
		lRUCache := Constructor(2)

		// lRUCache.put(1, 10); // cache: {1=10}
		lRUCache.Put(1, 10)

		// lRUCache.get(1); // return 10
		result1 := lRUCache.Get(1)
		if result1 != 10 {
			t.Errorf("Get(1) = %v, want %v", result1, 10)
		}

		// lRUCache.put(2, 20); // cache: {1=10, 2=20}
		lRUCache.Put(2, 20)

		// lRUCache.put(3, 30); // cache: {2=20, 3=30}, key=1 was evicted
		lRUCache.Put(3, 30)

		// lRUCache.get(2); // returns 20
		result2 := lRUCache.Get(2)
		if result2 != 20 {
			t.Errorf("Get(2) = %v, want %v", result2, 20)
		}

		// lRUCache.get(1); // return -1 (not found)
		result3 := lRUCache.Get(1)
		if result3 != -1 {
			t.Errorf("Get(1) = %v, want %v", result3, -1)
		}
	})
}
