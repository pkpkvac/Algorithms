# 2D Dynamic Programming

## What Makes a Problem "2D DP"?

**1D DP:** State depends on a single variable (index, amount, position)
- `dp[i]` = "answer for first i elements"

**2D DP:** State depends on two variables (two indices, two strings, grid position)
- `dp[i][j]` = "answer for first i elements of X and first j elements of Y"

## Common 2D DP Patterns

### Pattern 1: Grid Path Counting
**Example:** Unique Paths
- `dp[i][j]` = number of ways to reach cell (i, j)
- Recurrence: `dp[i][j] = dp[i-1][j] + dp[i][j-1]`

### Pattern 2: String Matching
**Example:** Longest Common Subsequence, Edit Distance
- `dp[i][j]` = answer for `s1[0:i]` and `s2[0:j]`
- Recurrence depends on whether `s1[i-1] == s2[j-1]`

### Pattern 3: Knapsack Variants
**Example:** 0/1 Knapsack
- `dp[i][w]` = max value using first i items with weight limit w

## How to Approach 2D DP

1. **Define the state:** What does `dp[i][j]` represent?
2. **Base cases:** What are `dp[0][j]` and `dp[i][0]`?
3. **Recurrence:** How does `dp[i][j]` relate to previous states?
4. **Fill order:** Usually row-by-row or column-by-column
5. **Answer location:** Often `dp[m][n]` (bottom-right)

## Key Difference from 1D DP

- **1D DP:** Linear dependencies (previous element)
- **2D DP:** Grid dependencies (top, left, or diagonal cells)

You already know 1D DP. 2D DP is just extending that to two dimensions.
