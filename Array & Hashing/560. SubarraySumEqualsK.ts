// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:

// Input: nums = [1,1,1], k = 2
// Output: 2
// Example 2:

// Input: nums = [1,2,3], k = 3
// Output: 2

function subarraySumBruteForce(nums: number[], k: number): number {
	let count = 0;
	let res: number[][] = [];
	for (let i = 0; i < nums.length; i++) {
		let sum = nums[i];
		if (sum === k) {
			count++;
			res.push([i]);
		}

		for (let j = i + 1; j < nums.length; j++) {
			sum += nums[j];
			if (sum === k) {
				count++;
				res.push([i, j]);
			}
		}
	}
	console.log(res);
	return count;
}

function subarraySumMap(nums: number[], k: number): number {
	let count = 0;
	let sum = 0;

	const map = new Map<number, number>();

	map.set(0, 1);

	for (let i = 0; i < nums.length; i++) {
		sum += nums[i];
		if (map.has(sum - k)) {
			count += map.get(sum - k) ?? 0;
		}
		map.set(sum, (map.get(sum) ?? 0) + 1);
	}

	return count;
}

const numeros = [1, 1, 1];
// const k = 3;
const k = 2;

// subarraySumBruteForce(numeros, k);
subarraySumMap(numeros, k);
