// Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring
// without duplicate characters.

// A substring is a contiguous sequence of characters within a string.

// Example 1:

// Input: s = "zxyzxyz"

// Output: 3
// Explanation: The string "xyz" is the longest without duplicate characters.

// Example 2:

// Input: s = "xxxx"

// Output: 1

function lengthOfLongestSubstring(s: string): number {
	if (!s) return 0;

	let l = 0;
	let r = 0;
	let currSeq = 0;
	let maxSeq = 0;
	const map = new Map<string, number>();

	while (r < s.length) {
		const char = s[r];
		if (map.has(char) && map.get(char)! >= l) {
			const prevIndex = map.get(char)!;
			l = Math.max(l, prevIndex + 1);
		}
		currSeq = r - l + 1;
		map.set(char, r);
		maxSeq = Math.max(currSeq, maxSeq);
		r++;
	}
	return maxSeq;
}
