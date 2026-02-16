package climbing_stairs

// naive solution, very slow, recomputes sub problems over and over again
// this is the 'recursive solution'

// interesting thing about this problem is that it's a fibonacci sequence

func climbStairs(n int) int {

	if n < 0 {
		// not a way
		return 0
	}
	if n == 0 {
		// this is 1 way
		return 1
	}

	numWays := climbStairs(n-1) + climbStairs(n-2)

	return numWays
}

// this is the recursive solution with cache
// memoization - use helper or closure, so memo map is used for every recursive call
// this is a "top down" style of problem solving
func climbStairsMemo(n int) int {

	m := make(map[int]int)

	m[0] = 1
	m[1] = 1
	m[2] = 2

	return climbStairsWithMemoHelper(n, m)

}

func climbStairsWithMemoHelper(n int, m map[int]int) int {

	if value, ok := m[n]; ok {
		// exists in cache, return it
		return value
	}

	res := climbStairsWithMemoHelper(n-1, m) + climbStairsWithMemoHelper(n-2, m)
	m[n] = res

	return res

}

// tabulation - no helper needed, one function, one loop,
// build up dp[0..n], with dp[i] = dp[i-1] + dp[i-2], then return dp[n]
// this is a 'bottom up' solution
// this is the iterative + table type of solution
func climbStairsTabulation(n int) int {
	dp := make([]int, n+1)

	// base case
	dp[0] = 1
	// 1 way to climb 1 step
	dp[1] = 1

	if n <= 1 {
		return dp[n]
	}

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
