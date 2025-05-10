// Given an array nums with n objects colored red, white, or blue, sort them in-place
// so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

// You must solve this problem without using the library's sort function.

// Example 1:

// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
// Example 2:

// Input: nums = [2,0,1]
// Output: [0,1,2]

function sortColors(nums: number[]): void {
	// find max element (we know it's 2)

	const countArray: number[] = Array(3).fill(0);

	for (let i = 0; i < nums.length; i++) {
		countArray[nums[i]]++;
	}

	// store cumulative, or prefix sum to help placing the elements
	for (let i = 1; i < countArray.length; i++) {
		countArray[i] = countArray[i - 1] + countArray[i];
	}
	console.log(countArray);
	// [2, 4, 6]

	// Iterate from end of the input array and because traversing input array from end preserves
	//  the order of equal elements, which eventually makes this sorting algorithm stable.

	// 	Update outputArray[ countArray[ inputArray[i] ] – 1] = inputArray[i].
	// Also, update countArray[ inputArray[i] ]  = countArray[ inputArray[i] ]– -.
	let outputArray: number[] = Array(nums.length).fill(0);

	for (let i = nums.length - 1; i >= 0; i--) {
		outputArray[countArray[nums[i]] - 1] = nums[i];
		countArray[nums[i]]--;
	}

	// Copy back to original array (in place)
	for (let i = 0; i < nums.length; i++) {
		nums[i] = outputArray[i];
	}

	console.log(nums);
}

const numbers = [2, 0, 2, 1, 1, 0];

sortColors(numbers);

// COMPLEXITY: O(n + k), which is better than quicksort (O(nlogn)), for small k
// CONS: mostly useful for when the range of values is not too big compared to the number of elements

// For example, if you're sorting 100 integers in the range 0-1,000,000,
// counting sort would be inefficient. But for sorting 1,000,000 integers in the range 0-100, it would be extremely fast.
