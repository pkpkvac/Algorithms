// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

function twoSumBruteForce(nums: number[], target: number): number[] {
	for (let i = 0; i < nums.length; i++) {
		let first = nums[i];
		let need = target - first;
		for (let j = i + 1; j < nums.length; j++) {
			if (nums[j] === need) {
				return [i, j];
			}
		}
	}

	// O(n^2)
	return [0, 0];
}

function twoSum(nums: number[], target: number): number[] {
	const map = new Map<number, number>();

	for (let i = 0; i < nums.length; i++) {
		// Calculate what number we need to find
		const complement = target - nums[i];

		// Check if we've seen the complement before
		if (map.has(complement)) {
			return [map.get(complement)!, i];
		}

		// Store current number and its index
		map.set(nums[i], i);
	}

	return [-1, -1]; // No solution found
}

// const nums = [2, 7, 11, 15];
// const target = 9;
const nums = [3, 2, 4];
const target = 6;

// console.log(twoSumBruteForce(nums, target));
console.log(twoSum(nums, target));
