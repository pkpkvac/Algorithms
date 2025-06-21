# Design and implement circular queue. The circular queue
#  is a linear data structure in which the operations are
#  performed based on FIFO (First In First Out) principle,
#  and the last position is connected back to the first
#  position to make a circle. It is also called "Ring Buffer".

# One of the benefits of the circular queue is that we can
#  make use of the spaces in front of the queue. In a normal
#  queue, once the queue becomes full, we cannot insert the
#  next element even if there is a space in front of the
#  queue. But using the circular queue, we can use
#  the space to store new values.

# ============================================================================
# ARE CIRCULAR QUEUES USED OFTEN? YES! EVERYWHERE!
# ============================================================================

# Circular queues are EXTREMELY common in real-world systems:

# 🖥️  OPERATING SYSTEMS:
# - Process scheduling (CPU time slices)
# - Keyboard/mouse input buffers
# - Print job queues
# - System call buffers

# 🌐 NETWORKING:
# - Network packet buffers (routers, switches)
# - TCP/UDP receive buffers
# - Flow control in protocols
# - Load balancer request queues

# 🎵 AUDIO/VIDEO:
# - Audio streaming buffers (Spotify, YouTube)
# - Video frame buffers
# - Real-time audio processing (music production)
# - Gaming audio engines

# 🔧 EMBEDDED SYSTEMS:
# - Sensor data collection
# - UART/SPI communication buffers
# - Microcontroller I/O
# - IoT device message queues

# 🏭 PRODUCER-CONSUMER PROBLEMS:
# - Multi-threaded applications
# - Message queues (Redis, RabbitMQ)
# - Event-driven architectures
# - Stream processing (Kafka)

# 💾 HARDWARE:
# - CPU instruction pipelines
# - Memory controllers
# - Graphics card frame buffers
# - Network interface cards

# ============================================================================
# WHY ARE THEY SO POPULAR?
# ============================================================================

# 1. FIXED MEMORY USAGE:
#    - Pre-allocated array (no dynamic allocation)
#    - Perfect for embedded systems with limited RAM
#    - No memory fragmentation

# 2. CONSTANT TIME OPERATIONS:
#    - O(1) enqueue, dequeue, front, rear
#    - No shifting elements like in regular arrays
#    - Predictable performance

# 3. EFFICIENT SPACE UTILIZATION:
#    - Reuses freed space automatically
#    - No wasted memory from "holes"
#    - Maximum capacity always available

# 4. CACHE-FRIENDLY:
#    - Data stored contiguously in memory
#    - Better CPU cache performance
#    - Important for high-frequency operations

# 5. THREAD-SAFE VARIANTS:
#    - Lock-free implementations possible
#    - Single producer, single consumer patterns
#    - High-performance concurrent systems

# ============================================================================
# REAL-WORLD EXAMPLE: AUDIO STREAMING
# ============================================================================

# Imagine Spotify playing music:
# 1. Network thread downloads audio chunks → enQueue()
# 2. Audio thread plays chunks → deQueue()
# 3. If network is slow: buffer runs low but doesn't crash
# 4. If network is fast: buffer fills up, prevents memory bloat
# 5. Circular buffer automatically reuses space as audio plays

# Without circular queue: 
# - Regular queue would grow indefinitely (memory leak)
# - Array would need expensive shifting operations
# - Dynamic allocation would cause audio stuttering

# ============================================================================
# INTERVIEW IMPORTANCE
# ============================================================================

# This is NOT just an academic exercise!
# Interviewers ask about circular queues because:
# - Tests understanding of memory management
# - Shows knowledge of performance optimization
# - Demonstrates systems thinking
# - Common in system design questions

# Companies that use circular queues heavily:
# - Netflix (video streaming)
# - Uber (ride request buffering) 
# - Google (search result caching)
# - Amazon (order processing)
# - Any company with real-time systems!

# ============================================================================
# ARRAY VS LINKED LIST IMPLEMENTATION
# ============================================================================

# CIRCULAR QUEUES USE ARRAYS, NOT LINKED LISTS!
#
# WHY ARRAYS ARE PREFERRED:
# 
# 1. FIXED SIZE FITS PERFECTLY:
#    - Circular queues have fixed capacity
#    - Arrays naturally have fixed size
#    - Perfect match!
#
# 2. O(1) INDEX ACCESS:
#    - Arrays: queue[index] is O(1)
#    - Linked lists: must traverse from head - O(n)
#    - Critical for performance
#
# 3. MEMORY EFFICIENCY:
#    - Arrays: Just store the data
#    - Linked lists: Data + pointer overhead (8+ bytes per node)
#    - Significant savings for large buffers
#
# 4. CACHE PERFORMANCE:
#    - Arrays: Data stored contiguously → great cache locality
#    - Linked lists: Nodes scattered in memory → cache misses
#    - Huge performance difference in practice
#
# 5. NO DYNAMIC ALLOCATION:
#    - Arrays: Pre-allocated once
#    - Linked lists: malloc/free on every enqueue/dequeue
#    - Eliminates allocation overhead and fragmentation

# ============================================================================
# KEY IMPLEMENTATION CONCEPTS
# ============================================================================

# The "CIRCULAR" part is achieved with MODULO ARITHMETIC:
# 
# Instead of:
#   if (rear == capacity - 1):
#       rear = 0  # wrap around manually
#   else:
#       rear = rear + 1
#
# We use:
#   rear = (rear + 1) % capacity  # automatic wrap-around!
#
# CRITICAL CHALLENGE: How to distinguish EMPTY vs FULL?
# Both states can have front == rear!
#
# Common solutions:
# 1. Keep a separate 'size' counter
# 2. Keep a separate 'is_full' flag  
# 3. Sacrifice one slot (capacity - 1 items max)
# 4. Use sentinel values

# HINTS (without spoilers):
# 🎯 You'll need: array, front index, rear index, size counter
# 🎯 All operations should be O(1)
# 🎯 Use modulo (%) for wrap-around
# 🎯 Think about edge cases: empty queue, full queue

# Implement the MyCircularQueue class:

# MyCircularQueue(k) Initializes the object with the size of the queue to be k.
# int Front() Gets the front item from the queue. If the queue is empty, return -1.
# int Rear() Gets the last item from the queue. If the queue is empty, return -1.
# boolean enQueue(int value) Inserts an element into the circular queue. Return 
# true if the operation is successful.
# boolean deQueue() Deletes an element from the circular queue. Return true
#  if the operation is successful.
# boolean isEmpty() Checks whether the circular queue is empty or not.
# boolean isFull() Checks whether the circular queue is full or not.
# You must solve the problem without using the built-in queue data 
# structure in your programming language.
# Example 1:

# Input: ["MyCircularQueue", "enQueue", "enQueue", "enQueue",
#  "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
# [[3], [1], [2], [3], [4], [], [], [], [4], []]

# Output: [null, true, true, true, false, 3, true, true, true, 4]
# Explanation:
# MyCircularQueue myCircularQueue = new MyCircularQueue(3);
# myCircularQueue.enQueue(1); // return True
# myCircularQueue.enQueue(2); // return True
# myCircularQueue.enQueue(3); // return True
# myCircularQueue.enQueue(4); // return False
# myCircularQueue.Rear(); // return 3
# myCircularQueue.isFull(); // return True
# myCircularQueue.deQueue(); // return True
# myCircularQueue.enQueue(4); // return True
# myCircularQueue.Rear(); // return 4

# Constraints:

# 1 <= k <= 1000.
# 0 <= value <= 1000
# At most 3000 calls will be made to enQueue, deQueue,
#  Front, Rear, isEmpty, and isFull.


class MyCircularQueue:

    def __init__(self, k: int):
        # FIXED: All variables are now instance variables (with self.)
        self.capacity = k
        self.buffer = [0] * k  # Pre-allocate array
        self.front = 0
        self.rear = 0
        self.size = 0  # This is all we need to track empty/full

    def enQueue(self, value: int) -> bool:
        if self.isFull():
            return False
        
        self.buffer[self.rear] = value
        self.rear = (self.rear + 1) % self.capacity  # FIXED: Use modulo
        self.size += 1
        return True

    def deQueue(self) -> bool:  # FIXED: Return boolean, not value
        if self.isEmpty():
            return False
        
        # Don't need to actually "remove" - just move front pointer
        self.front = (self.front + 1) % self.capacity  # FIXED: Use modulo
        self.size -= 1
        return True

    def Front(self) -> int:
        if self.isEmpty():
            return -1
        return self.buffer[self.front]  # FIXED: front points to actual element

    def Rear(self) -> int:
        if self.isEmpty():
            return -1
        # FIXED: rear points to NEXT slot, so last element is rear-1
        return self.buffer[(self.rear - 1) % self.capacity]

    def isEmpty(self) -> bool:
        return self.size == 0  # FIXED: Simple size check

    def isFull(self) -> bool:
        return self.size == self.capacity  # FIXED: Simple size check

# Test the implementation
queue = MyCircularQueue(3)
print(f"enQueue(1): {queue.enQueue(1)}")  # True
print(f"enQueue(2): {queue.enQueue(2)}")  # True  
print(f"enQueue(3): {queue.enQueue(3)}")  # True
print(f"enQueue(4): {queue.enQueue(4)}")  # False (full)
print(f"Rear(): {queue.Rear()}")          # 3
print(f"isFull(): {queue.isFull()}")      # True
print(f"deQueue(): {queue.deQueue()}")    # True
print(f"enQueue(4): {queue.enQueue(4)}")  # True
print(f"Rear(): {queue.Rear()}")          # 4
