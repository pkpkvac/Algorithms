// You are given an integer array height of length n. There are
//  n vertical lines drawn such that the two endpoints of the ith
// line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container,
//  such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49

function maxAreaBruteForce(height: number[]): number {
	let maxArea = 0;
	for (let i = 0; i < height.length; i++) {
		for (let j = i + 1; j < height.length; j++) {
			const b = j - i;
			const h = Math.min(height[i], height[j]);
			maxArea = Math.max(maxArea, b * h);
		}
	}

	return maxArea;
}

function maxArea(height: number[]): number {
	let L = 0;
	let R = height.length - 1;
	let maxArea = 0;
	while (L < R) {
		const h = Math.min(height[L], height[R]);
		const b = R - L;

		maxArea = Math.max(maxArea, h * b);

		if (height[L] < height[R]) {
			L++;
		} else {
			R--;
		}
	}
	return maxArea;
}

const height = [1, 8, 6, 2, 5, 4, 8, 3, 7];

// maxAreaBruteForce(height);
console.log(maxArea(height));
