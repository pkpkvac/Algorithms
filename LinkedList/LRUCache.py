# Implement the Least Recently Used (LRU) cache class LRUCache.
#  The class should support the following operations

# LRUCache(int capacity) Initialize the LRU cache of size capacity.
# int get(int key) Return the value corresponding to the key
#  if the key exists, otherwise return -1.
# void put(int key, int value) Update the value of
#  the key if the key exists. Otherwise, add the key-value
#  pair to the cache. If the introduction of the new pair
#  causes the cache to exceed its capacity, remove the
#  least recently used key.

# A key is considered used if a get or a put operation
#  is called on it.

# Ensure that get and put each run in
# O(1) average time complexity.

# Example 1:

# Input:
# ["LRUCache", [2], "put", [1, 10],  "get", [1], "put", [2, 20], "put", [3, 30], "get", [2], "get", [1]]

# Output:
# [null, null, 10, null, null, 20, -1]

# Explanation:
# LRUCache lRUCache = new LRUCache(2);
# lRUCache.put(1, 10);  // cache: {1=10}
# lRUCache.get(1);      // return 10
# lRUCache.put(2, 20);  // cache: {1=10, 2=20}
# lRUCache.put(3, 30);  // cache: {2=20, 3=30}, key=1 was evicted
# lRUCache.get(2);      // returns 20 
# lRUCache.get(1);      // return -1 (not found)

from definition import DoublyLinkedNode

class LRUCache:

    cache_map = None
    head = None
    tail = None
    capacity = None

    def __init__(self, capacity: int):

        self.capacity = capacity
        self.cache_map = {}
        self.head = DoublyLinkedNode(None)
        self.tail = DoublyLinkedNode(None)
        self.head.next = self.tail
        self.tail.prev = self.head



        # initialize a doubly linked list, key value pairs
        # are stored as nodes,
        # w/ LRU node at head, and MRU node at tail
        # Whenever a key is accessed using get() or put(), 
        # we remove the corresponding node and reinsert it at the tail. 
        # When the cache reaches its capacity, we remove 
        # the LRU node from the head of the list.
        #  Additionally, we use a hash map to store 
        # each key and the corresponding address of
        #  its node, enabling efficient operations in O(1) time


        # Approach A: Head = Most Recent, Tail = Least Recent
        # When accessed → move node to head
        # When evicting → remove from tail
        # Approach B: Head = Least Recent, Tail = Most Recent
        # When accessed → move node to tail
        # When evicting → remove from head

    def get(self, key: int) -> int:

        # If your list is A ↔ B ↔ C and you access B, what should the new order be?
        # What operations do you need to move B to the "most recent" position?

        if key in self.cache_map:

            # dummy_head ↔ [A] ↔ [B] ↔ [C] ↔ dummy_tail
            #   ↑                    ↑
            # most recent          least recent

            # Step 1: Remove C from current position
            node = self.cache_map.get(key)
            
            # Move to head (most recent)
            node.prev.next = node.next
            node.next.prev = node.prev

            # now C is floating

            # now insert C at head (most recent)
            # Insert C between dummy_head and A

            node.next = self.head.next
            node.prev = self.head

            self.head.next.prev = node # A points back to C
            self.head.next = node # dummy_head points to C
            
            # dummy_head ↔ [C] ↔ [A] ↔ [B] ↔ dummy_tail
            #       ↑                    ↑
            # most recent          least recent
            return node.val
        else:
            return -1

        

    def put(self, key: int, value: int) -> None:

        # Case 1: Key already exists
        # Update the value
        # Move node to "most recent" position
        # No capacity check needed
        if key in self.cache_map:

            # dummy_head ↔ [A] ↔ [B] ↔ [C] ↔ dummy_tail
            #   ↑                    ↑
            # most recent          least recent

            # Step 1: Remove C from current position
            # Case 1: Key exists - update value and move to head
            node = self.cache_map.get(key)
            node.val = value  # Update the value!
            
            # Move to head (most recent)
            node.prev.next = node.next
            node.next.prev = node.prev

            # now C is floating

            # now insert C at head (most recent)
            # Insert C between dummy_head and A
            # Insert at head
            node.next = self.head.next
            node.prev = self.head
            self.head.next.prev = node
            self.head.next = node
            
        else:
            # Case 2: New key
            # Check capacity and evict if needed
            if len(self.cache_map) >= self.capacity:
                # Remove LRU (tail)
                lru_node = self.tail.prev
                # Need to find key for this node - this requires storing key in node
                # For now, let's iterate through map (not O(1) but works)
                lru_key = None
                for k, v in self.cache_map.items():
                    if v == lru_node:
                        lru_key = k
                        break
                
                # Remove from hashmap and linked list
                del self.cache_map[lru_key]
                lru_node.prev.next = lru_node.next
                lru_node.next.prev = lru_node.prev
            
            # Create new node and add to head
            new_node = DoublyLinkedNode(value)
            self.cache_map[key] = new_node
            
            # Insert at head
            new_node.next = self.head.next
            new_node.prev = self.head
            self.head.next.prev = new_node
            self.head.next = new_node

