// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order,
// starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

// Return the merged string.

// Example 1:

// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r
// Example 2:

// Input: word1 = "ab", word2 = "pqrs"
// Output: "apbqrs"
// Explanation: Notice that as word2 is longer, "rs" is appended to the end.
// word1:  a   b
// word2:    p   q   r   s
// merged: a p b q   r   s
// Example 3:

// Input: word1 = "abcd", word2 = "pq"
// Output: "apbqcd"
// Explanation: Notice that as word1 is longer, "cd" is appended to the end.
// word1:  a   b   c   d
// word2:    p   q
// merged: a p b q c   d

function mergeAlternately(word1: string, word2: string): string {
	let m = word1.length;
	let n = word2.length;
	let i = 0;
	let j = 0;
	let newStr = "";
	while (i < m || j < n) {
		if (i < m) {
			newStr += word1[i];
		}
		if (j < n) {
			newStr += word2[j];
		}
		i++;
		j++;
	}
	return newStr;
}

function mergeAlternatelyOnePass(word1: string, word2: string): string {
	let res = "";
	for (let i = 0; i < Math.max(word1.length, word2.length); i++) {
		if (i < word1.length) {
			res += word1[i];
		}
		if (i < word2.length) {
			res += word2[i];
		}
	}
	return res;
}

const word1 = "ab";
const word2 = "pqrs";

console.log(mergeAlternatelyOnePass(word1, word2));
