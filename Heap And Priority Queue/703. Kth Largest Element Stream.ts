// You are part of a university admissions office and need to keep track of the kth highest test score from applicants in real-time. This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.

// You are tasked to implement a class which, for a given integer k, maintains a stream of test
// scores and continuously returns the kth highest test score after a new score has been submitted.
// More specifically, we are looking for the kth highest score in the sorted list of all scores.

// Implement the KthLargest class:

// KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of test scores nums.
// int add(int val) Adds a new test score val to the stream and returns the element representing
// the kth largest element in the pool of test scores so far.

// Example 1:

// Input:
// ["KthLargest", "add", "add", "add", "add", "add"]
// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

// Output: [null, 4, 5, 5, 8, 8]

// Explanation:

// KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
// kthLargest.add(3); // return 4
// kthLargest.add(5); // return 5
// kthLargest.add(10); // return 5
// kthLargest.add(9); // return 8
// kthLargest.add(4); // return 8

// Example 2:

// Input:
// ["KthLargest", "add", "add", "add", "add"]
// [[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

// Output: [null, 7, 7, 7, 8]

// Explanation:

// KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
// kthLargest.add(2); // return 7
// kthLargest.add(10); // return 7
// kthLargest.add(9); // return 7
// kthLargest.add(9); // return 8

import { MinIntHeap } from "./MinAndMaxHeap";

class KthLargest {
	// 	Key Insight: Use a Min Heap of Size K
	// Rather than using a Max Heap for all elements, the optimal approach is to use a Min Heap of size K. This way:
	// The smallest element in your heap (the root) will always be the Kth largest element overall
	// You only need to maintain K elements, not all elements

	private k: number = 0;
	private heap: MinIntHeap = new MinIntHeap();
	constructor(k: number, nums: number[]) {
		this.k = k;

		for (let i = 0; i < nums.length; i++) {
			if (this.heap.getSize() < k) {
				this.heap.add(nums[i]);
			} else {
				this.addConditionally(nums[i]);
			}
		}
	}

	addConditionally(val: number): number {
		if (this.heap.peek()! <= val) {
			this.heap.poll();
			this.heap.add(val);
		}
		return this.heap.peek()!;
	}

	add(val: number): number {
		return this.addConditionally(val);
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * var obj = new KthLargest(k, nums)
 * var param_1 = obj.add(val)
 */
