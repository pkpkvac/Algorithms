package coin_change

import "math"

//  Rough complexity (naive, no memo):
//  Time: At each call you branch for every coin,
//  so it’s like a tree of degree len(coins) and
//  depth on the order of amount (if the smallest coin is 1).
//  So something like O(coins^amount) in the worst case — exponential.

// Space: The recursion stack is as deep as the longest path (e.g. all 1s), so O(amount).
func coinChange(coins []int, amount int) int {

	// what we are tracking is the number of coins to make
	// the amount, so the return value actually IS the minimum

	if amount == 0 {
		// you need 0 coints to make the amount
		return 0
	}
	if amount < 0 {
		// a very large number, because we'll take a minimum
		return math.MaxInt32
	}

	best := math.MaxInt32
	for _, coin := range coins {
		best = min(best, 1+coinChange(coins, amount-coin))
	}

	if best == math.MaxInt32 {
		return -1
	}
	return best
}

// The important point: the minimum number of coins to make a
// given amount doesn’t depend on how you got to that amount.

// Example: amount = 7, coins = [1, 2, 5].
// You can get to “make 7” in different ways:
// From 12 by using a 5: “make 7”
// From 8 by using a 1: “make 7”
// From 9 by using a 2: “make 7”
// But the question is always the same: “What is the minimum number of coins to make
// exactly 7?” The answer is always 2 (e.g. 5+2). It doesn’t matter
// whether we’re solving it from 12→7, 8→7, or 9→7.
// So the subproblem is fully described by one number: the remaining amount.
// The denominations don’t change (they’re fixed in coins); we only care
// “for this amount, what’s the min coins?” So we can memoize by amount only.
// Memo: memo[amount] = “minimum number of coins to make exactly amount (with the fixed coins).”

func coinChangeMemoized(coins []int, amount int) int {

	m := make(map[int]int)

	res := coinChangeMemoizedHelper(coins, m, amount)

	if res == math.MaxInt32 {
		return -1
	}

	return res

}

func coinChangeMemoizedHelper(coins []int, m map[int]int, amount int) int {

	// what we are tracking is the number of coins to make
	// the amount, so the return value actually IS the minimum

	if value, ok := m[amount]; ok {
		return value
	}

	if amount == 0 {
		// you need 0 coints to make the amount
		return 0
	}
	if amount < 0 {
		// a very large number, because we'll take a minimum
		return math.MaxInt32
	}

	best := math.MaxInt32
	for _, coin := range coins {
		best = min(best, 1+coinChangeMemoizedHelper(coins, m, amount-coin))
		m[amount] = best
	}

	return best

}

// tabulation can be done here.
// base case:
// dp[i] = minimum number of coins to make amount i (but we need to consider the coin)
// order:
// fill from i = 1, up to i = amount  i = [1...amount]
// recurrence: for each i, do every coin in coins
// for coin <= i: dp[i] = min(dp[i], 1 + dp[i-coin])
// for coin > i: skip, or treat as impossible
func coinChangeTabulation(coins []int, amount int) int {

	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			// If I’m trying to make amount i, what happens if I use this coin?

			// If we use coin, we need 1 coin plus the best way to make i - coin.
			// So one candidate for dp[i] is 1 + dp[i - coin]
			if coin <= i {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
