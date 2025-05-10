// We are playing the Guess Game. The game is as follows:

// I pick a number from 1 to n. You have to guess which number I picked.

// Every time you guess wrong, I will tell you whether the number
//  I picked is higher or lower than your guess.

// You call a pre-defined API int guess(int num), which returns three possible results:

// -1: Your guess is higher than the number I picked (i.e. num > pick).
// 1: Your guess is lower than the number I picked (i.e. num < pick).
// 0: your guess is equal to the number I picked (i.e. num == pick).
// Return the number that I picked.

// Example 1:

// Input: n = 10, pick = 6
// Output: 6
// Example 2:

// Input: n = 1, pick = 1
// Output: 1
// Example 3:

// Input: n = 2, pick = 1
// Output: 1

/**
 * Forward declaration of guess API.
 * @param {number} num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * var guess = function(num) {}
 */

function guessNumber(n: number): number {
	let l = 1;
	let r = n;

	while (l <= r) {
		// Use Math.floor or integer division to avoid potential issues
		let mid = Math.floor((l + r) / 2);

		const result = guess(mid);

		if (result === -1) {
			// Our guess is too high, search in lower half
			r = mid - 1;
		} else if (result === 1) {
			// Our guess is too low, search in upper half
			l = mid + 1;
		} else {
			// We found the answer
			return mid;
		}
	}

	// The problem guarantees a valid answer
	return -1; // This should never be reached
}
