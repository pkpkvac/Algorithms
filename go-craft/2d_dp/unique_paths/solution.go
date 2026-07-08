package unique_paths

func uniquePathsTabulation(m int, n int) int {
	// easier method - tabulation
	// dp[i][j] = number of ways to reach cell (i, j) from (0, 0)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	// base case - first row and first column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	
	// fill the dp table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	
	return dp[m-1][n-1]
}


// this is the memoization approach, top down
func uniquePaths(m int, n int) int {
    memo := make(map[[2]int]int)
    return dfs(0, 0, m, n, memo)
}
 
func dfs(i, j, m, n int, memo map[[2]int]int) int {
    // Reached destination
    if i == m-1 && j == n-1 {
        return 1
    }
    
    // Out of bounds
    if i >= m || j >= n {
        return 0
    }
    
    // Check memo
    key := [2]int{i, j}
    if val, exists := memo[key]; exists {
        return val
    }
    
    // Recurse: go down + go right
    ways := dfs(i+1, j, m, n, memo) + dfs(i, j+1, m, n, memo)
    
    // Store in memo
    memo[key] = ways
    return ways
}
