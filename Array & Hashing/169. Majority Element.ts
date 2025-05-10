// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

function majorityElement(nums: number[]): number {
	const count = Math.ceil(nums.length / 2);

	const map = new Map();

	for (let i = 0; i < nums.length; i++) {
		if (map.has(nums[i])) {
			const occurences = map.get(nums[i]) + 1;
			map.set(nums[i], occurences);

			if (occurences >= count) return nums[i];
		} else {
			map.set(nums[i], 1);
			if (count === 1) return nums[i];
		}
	}

	return -1;
}

const numsArray = [2, 2, 1, 1, 1, 2, 2];

console.log(majorityElement(numsArray));
