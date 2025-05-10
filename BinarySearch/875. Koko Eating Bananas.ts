// You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.

// You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.

// Return the minimum integer k such that you can eat all the bananas within h hours.

// Example 1:
// Input: piles = [1,4,3,2], h = 9
// Output: 2
// Explanation: With an eating rate of 2, you can eat the bananas in 6 hours. With an eating rate of 1, you would need 10 hours to eat all the bananas (which exceeds h=9), thus the minimum eating rate is 2.

// Example 2:
// Input: piles = [25,10,23,4], h = 4
// Output: 25
// Constraints:

// 1 <= piles.length <= 1,000
// piles.length <= h <= 1,000,000
// 1 <= piles[i] <= 1,000,000,000

function minEatingSpeedBruteForce(piles: number[], h: number): number {
	// Find maximum pile
	let maxPile = Math.max(...piles);

	// Try every eating speed from 1 to maxPile
	for (let k = 1; k <= maxPile; k++) {
		let totalHours = 0;

		// Calculate hours needed at current speed k
		for (const pile of piles) {
			totalHours += Math.ceil(pile / k);
		}

		// If we found a valid k, return it
		// (it will be minimum because we're checking in order)
		if (totalHours <= h) {
			return k;
		}
	}

	return maxPile; // Problem guarantees a solution exists
}

function minEatingSpeedBinarySearch(piles: number[], h: number): number {
	let maxPile = Math.max(...piles);

	let L = 1,
		R = maxPile;
	while (L < R) {
		let mid = Math.floor((L + R) / 2);
		let totalT = 0;

		for (let i = 0; i < piles.length; i++) {
			totalT += Math.ceil(piles[i] / mid);
		}

		// we want to eat bananas slowly, but eat them all within h hours
		// if time is less than h, maybe we can eat them slower, setting R = mid
		if (totalT <= h) {
			// take the minimum
			R = mid;
		} else {
			// we're eating too slow, boost L
			L = mid + 1;
		}
	}
	return L;
}

// const piles = [1, 4, 3, 2];
// const h = 9;
const piles = [25, 10, 23, 4];
const h = 4;

// console.log(minEatingSpeedBruteForce(piles, h));
console.log(minEatingSpeedBinarySearch(piles, h));
