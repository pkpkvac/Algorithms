import { MaxIntHeap } from "./MinAndMaxHeap";

// Given an integer array nums and an integer k, return the kth largest element in the array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
// Can you solve it without sorting?

// Example 1:

// Input: nums = [3,2,1,5,6,4], k = 2
// Output: 5
// Example 2:

// Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
// Output: 4

function findKthLargestBruteForce(nums: number[], k: number): number {
	// copy the array first

	let numCopy = [...nums];

	// remove k elements
	for (let i = 0; i < k - 1; i++) {
		let largestIndex = 0;
		let largestVal = -Infinity;
		for (let j = 0; j < nums.length; j++) {
			// find the largest element, set its value to -Infinity
			// to remove it
			if (numCopy[j] > largestVal) {
				largestVal = numCopy[j];
				largestIndex = j;
			}
		}
		numCopy[largestIndex] = -Infinity;
	}

	// now just scan and find the largest element since the k largest elements
	// have been removed
	let kthValue = -Infinity;
	for (let i = 0; i < numCopy.length; i++) {
		if (numCopy[i] > kthValue) {
			kthValue = numCopy[i];
		}
	}

	return kthValue;
}
// 	O(n) for copy
// + O(k * n) for removing k-1 elements
// + O(n) for final scan
// = O(kn) where k < n
// = O(n²) worst case

// utilizing your heap
function findKthLargestYourHeap(nums: number[], k: number): number {
	const heap = new MaxIntHeap();

	for (const num of nums) {
		heap.add(num);
	}

	let result = 0;
	for (let i = 0; i < k - 1; i++) {
		result = heap.poll() ?? 0;
	}

	return result;
}

const nums = [3, 2, 1, 5, 6, 4],
	k = 2;

// console.log(findKthLargestBruteForce(nums, k));
console.log(findKthLargestYourHeap(nums, k));
