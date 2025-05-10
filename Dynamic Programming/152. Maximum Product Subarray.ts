// Note:
// Use Kadane's Algorithm

// Given an integer array nums, find a subarray that has the largest product within the array and return it.
// A subarray is a contiguous non-empty sequence of elements within an array.
// You can assume the output will fit into a 32-bit integer.

// Example 1:
// Input: nums = [1,2,-3,4]
// Output: 4

// Example 2:
// Input: nums = [-2,-1]
// Output: 2

function maxProductBruteForce(nums: number[]): number {
	let maxProduct = nums[0];

	// Try every possible subarray
	for (let i = 0; i < nums.length; i++) {
		let currProduct = nums[i];
		// Update max for single number
		maxProduct = Math.max(maxProduct, currProduct);

		// Try multiplying with subsequent numbers
		for (let j = i + 1; j < nums.length; j++) {
			currProduct *= nums[j];
			maxProduct = Math.max(maxProduct, currProduct);
		}
	}

	return maxProduct;
}

function maxProductKadanes(nums: number[]): number {}

const nums = [1, 2, -3, 4];
// const nums = [-2, -1];

// [-1,-2,-3,0]
// output 6

console.log(maxProductBruteForce(nums));
