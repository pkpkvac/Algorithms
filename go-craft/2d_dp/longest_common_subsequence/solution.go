package longest_common_subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] = length of LCS for text1[0:i] and text2[0:j]
	m, n := len(text1), len(text2)
	
	// Create 2D dp table with (m+1) x (n+1) dimensions
	// Extra row/col for empty string base case
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Base case: dp[0][j] = 0 and dp[i][0] = 0 (already initialized to 0)
	// Empty string has LCS of 0 with any string
	
	// Example: text1 = "ace", text2 = "abcde"
	//       ""  a  b  c  d  e
	//   ""   0  0  0  0  0  0
	//   a    0  1  1  1  1  1  ← 'a'=='a' at (1,1): dp[0][0]+1 = 1
	//   c    0  1  1  2  2  2  ← 'c'=='c' at (2,3): dp[1][2]+1 = 2
	//   e    0  1  1  2  2  3  ← 'e'=='e' at (3,5): dp[2][4]+1 = 3
	//   Answer: dp[3][5] = 3 (LCS = "ace")
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Compare characters at text1[i-1] and text2[j-1]
			// (i-1 and j-1 because dp is 1-indexed but strings are 0-indexed)
			
			if text1[i-1] == text2[j-1] {
				// Characters match! Extend the LCS from diagonal
				// dp[i-1][j-1] represents LCS without these matching chars
				// Add 1 to include this matching character
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// Characters don't match
				// Take the best result from either:
				// - Skipping current char in text1: dp[i-1][j]
				// - Skipping current char in text2: dp[i][j-1]
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	// Answer is at bottom-right: LCS of full text1 and full text2
	return dp[m][n]
}
