from collections import deque
import heapq
import math
from typing import *

# You are given an array of integers cost where cost[i] is the cost 
# of taking a step from the ith floor of a staircase. After paying
#  the cost, you can step to either the (i + 1)th floor or the (i + 2)th floor.

# You may choose to start at the index 0 or the index 1 floor.

# Return the minimum cost to reach the top of the staircase, i.e. just past the last index in cost.

# Example 1:

# Input: cost = [1,2,3]

# Output: 2
# Explanation: We can start at index = 1 and pay the cost of cost[1] = 2 and take two steps to reach the top. The total cost is 2.

# Example 2:

# Input: cost = [1,2,1,2,1,1,1]

# Output: 4
# Explanation: Start at index = 0.

# Pay the cost of cost[0] = 1 and take two steps to reach index = 2.
# Pay the cost of cost[2] = 1 and take two steps to reach index = 4.
# Pay the cost of cost[4] = 1 and take two steps to reach index = 6.
# Pay the cost of cost[6] = 1 and take one step to reach the top.
# The total cost is 4.

class Solution:

    # Bottom Up (Tabulation) - iterative, table, for while, calculates all solutions
    def minCostClimbingStairsTabulation(self, cost: List[int]) -> int:
        n = len(cost)
        # dp[i] -> min cost to reach top from position i
        dp = [0] * (n + 1)
        # base case: cost to reach the top from the top is 0
        dp[n] = 0

        for i in range(n -1, -1, -1):
            dp[i] = cost[i] + min(dp[i + 1], dp[i + 2] if i + 2 <= n else 0)

        return min(dp[0], dp[1])
    
    # Top Down (Memoization) - recursive, avoids redundant calcs, uses only dp memo table
    def minCostClimbingStairsMemoization(self, cost: List[int]) -> int:
        n = len(cost)
        memo = {}
        def dp(i):
            if i >= n: return 0
            if i in memo: return memo[i]

            # recursive call, take 1 or 2 steps
            memo[i] = cost[i] + min(dp(i + 1), dp(i + 2))
            return memo[i]
        return min(dp(0), dp(1))




cost = [1,2,1,2,1,1,1]
solution = Solution()

print(solution.minCostClimbingStairsMemoization(cost))
print(solution.minCostClimbingStairsTabulation(cost))