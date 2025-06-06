from typing import List

# You are given an integer array coins representing coins of 
# different denominations (e.g. 1 dollar, 5 dollars, etc) and
#  an integer amount representing a target amount of money.

# Return the fewest number of coins that you need to make 
# up the exact target amount. If it is impossible to make 
# up the amount, return -1.

# You may assume that you have an unlimited number of each coin.

# Example 1:

# Input: coins = [1,5,10], amount = 12

# Output: 3
# Explanation: 12 = 10 + 1 + 1. Note that we do not have to 
# use every kind coin available.

# Example 2:

# Input: coins = [2], amount = 3

# Output: -1
# Explanation: The amount of 3 cannot be made up with coins of 2.

# Example 3:

# Input: coins = [1], amount = 0

# Output: 0
# Explanation: Choosing 0 coins is a valid way to make up 0.


# Why Backtracking is Inefficient Here
# Backtracking approach: Generate ALL possible combinations that sum to amount, then pick the shortest.
# For coins = [1,5,10], amount = 12, backtracking would explore:
# 1+1+1+1+1+1+1+1+1+1+1+1 (12 coins)
# 1+1+1+1+1+1+1+5 (8 coins)
# 1+1+5+5 (4 coins)
# 10+1+1 (3 coins) ← optimal
# 5+5+1+1 (4 coins)
# ... and many more!
# Problem: Exponential explosion! With unlimited coins, 
# there are potentially millions of combinations.

# Dynamic Programming Key Insight
# Instead of finding ALL solutions, 
# DP asks: "What's the optimal way to build up to each amount?"


class Solution:
    def coinChangeTabulation(self, coins: List[int], amount: int) -> int:
        # dp[i] = min coins to make amount i
        dp = [float('inf')] * (amount + 1)
        dp[0] = 0  # Base case: 0 coins needed for amount 0
        
        # Fill table from amount 1 to target amount
        for current_amount in range(1, amount + 1):
            for coin in coins:
                if coin <= current_amount:  # Can we use this coin?
                    dp[current_amount] = min(dp[current_amount],  1 + dp[current_amount - coin])
        
        return dp[amount] if dp[amount] != float('inf') else -1

    def coinChangeMemoization(self, coins: List[int], amount: int) -> int:
        memo = {}

        def dp(remaining):
            if remaining in memo:
                return memo[remaining]
            
            # base case
            if remaining == 0: return 0
            if remaining < 0: return float('inf') #not possible

            # try each coin
            min_coins = float('inf')
            for coin in coins:
                result = dp(remaining - coin)
                min_coins = min(min_coins, 1 + result)
            
            memo[remaining] = min_coins

            return min_coins
        
        return dp(amount)


coins = [1,5,10]
amount = 12

solution = Solution()

print(solution.coinChangeTabulation(coins, amount))