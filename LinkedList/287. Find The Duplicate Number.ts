import { ListNode, printLinkedList, createLinkedList } from "./definition";

// Given an array of integers nums containing n + 1 integers where each integer
// is in the range [1, n] inclusive.

// There is only one repeated number in nums, return this repeated number.

// You must solve the problem without modifying the array nums and using only constant extra space.

// Example 1:

// Input: nums = [1,3,4,2,2]
// Output: 2
// Example 2:

// Input: nums = [3,1,3,4,2]
// Output: 3
// Example 3:

// Input: nums = [3,3,3,3,3]
// Output: 3

// Why This Works
// The duplicate number creates a cycle because two different indices point to the same next location
// The mathematics of Floyd's algorithm guarantees that:
// 1. If there's a cycle, the fast and slow pointers will meet
// When they meet, if we reset one pointer to the start and move both at the same speed, they'll meet at the cycle entrance
// The cycle entrance corresponds to our duplicate number

function findDuplicate(nums: number[]): number {
	// Phase 1: Find meeting point inside the cycle
	let slow = nums[0];
	let fast = nums[0];

	do {
		slow = nums[slow];
		fast = nums[nums[fast]];
	} while (slow !== fast);

	// Phase 2: Find the entrance to the cycle
	slow = nums[0];
	while (slow !== fast) {
		slow = nums[slow];
		fast = nums[fast];
	}

	return slow;
}
