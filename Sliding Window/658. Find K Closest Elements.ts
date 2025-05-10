// Given a sorted integer array arr, two integers k and x,
// return the k closest integers to x in the array. The result should also be sorted in ascending order.

// An integer a is closer to x than an integer b if:

// |a - x| < |b - x|, or
// |a - x| == |b - x| and a < b

// Example 1:
// Input: arr = [1,2,3,4,5], k = 4, x = 3
// Output: [1,2,3,4]

// Example 2:
// Input: arr = [1,1,2,3,4,5], k = 4, x = -1
// Output: [1,1,2,3]

// Constraints:
// 1 <= k <= arr.length
// 1 <= arr.length <= 104
// arr is sorted in ascending order.
// -104 <= arr[i], x <= 104

function findClosestElements2(arr: number[], k: number, x: number): number[] {
	let left = 0;
	let right = arr.length - 1;

	// Eliminate elements until we have k elements left
	while (right - left + 1 > k) {
		const diffLeft = Math.abs(arr[left] - x);
		const diffRight = Math.abs(arr[right] - x);

		if (diffLeft <= diffRight) {
			right--;
		} else {
			left++;
		}
	}

	return arr.slice(left, right + 1);
}

const arr = [1, 2, 3, 4, 5];
const k_val = 4;
const x = 3;

console.log(findClosestElements2(arr, k_val, x));
