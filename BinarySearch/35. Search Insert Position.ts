// Given a sorted array of distinct integers and a target value,
//  return the index if the target is found. If not, return the
//  index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:

// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4

function searchInsert(nums: number[], target: number): number {
	let l = 0;
	let r = nums.length - 1;
	while (l <= r) {
		let mid = Math.floor((l + r) / 2);
		if (nums[mid] < target) {
			// move left up
			l = mid + 1;
		} else if (nums[mid] > target) {
			r = mid - 1;
		} else {
			return mid;
		}
	}
	return l;
}

const nums = [1, 3, 5, 6];
const target = 7;

console.log(searchInsert(nums, target));
