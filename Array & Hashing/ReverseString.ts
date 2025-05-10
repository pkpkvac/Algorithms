// Write a function that reverses a string. The input string is given as an array of characters s.

// You must do this by modifying the input array in-place with O(1) extra memory.

// Example 1:

// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]
// Example 2:

// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]

/**
 Do not return anything, modify s in-place instead.
 */
function reverseString(s: string[]): void {
	let l = 0;
	let r = s.length - 1;

	while (l < r) {
		let tmp = s[l];
		s[l] = s[r];
		s[r] = tmp;
		l++;
		r--;
	}
	console.log(s);
}

let s = ["H", "a", "n", "n", "a", "h"];
reverseString(s);
