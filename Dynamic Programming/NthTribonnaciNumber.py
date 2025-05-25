# The Tribonacci sequence Tn is defined as follows: 

# T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

# Given n, return the value of Tn.

# Example 1:

# Input: n = 4
# Output: 4
# Explanation:
# T_3 = 0 + 1 + 1 = 2
# T_4 = 1 + 1 + 2 = 4
# Example 2:

# Input: n = 25
# Output: 1389537

class Solution:
        
    def tribonacciTabulation(self, n: int) -> int:
        # Tabulation (Bottom-up): Start with base cases, build up to T(n) 

        # base case
        dp = [0] * (n + 1)
        dp[0] = 0
        dp[1] = 1
        dp[2] = 1

        if n <= 2: return dp[n]

        for i in range(3, n + 1):
            dp[i] = dp[i - 3] + dp[i -2] + dp[i - 1]

        print(dp)
        return dp[n]
    
    def tribonacciMemoization(self, n: int) -> int:
        # Memoization (Top-down): 
        # You want tribonacci(n)
        # what do you need? tribonacci (n-1) + tribonacci(n-2) + tribonacci(n-3)
        # ... and so on, when do you stop?
        # when you hit the bases cases (0, 1, 2), which you have
        memo = {}
        return self.tribonacci_memo(n, memo)

    def tribonacci_memo(self, n: int, memo: dict = None) -> int:

        if n in memo:
            return memo[n]
        
        if n == 0: return 0
        if n == 1: return 1
        if n == 2: return 1

        memo[n] = (self.tribonacci_memo(n-1, memo) + self.tribonacci_memo(n-2, memo) + self.tribonacci_memo(n-3, memo))

        return memo[n]


        

    
n = 25 
solution = Solution()
print("Tabulation result:", solution.tribonacciTabulation(n))
print("Memoization result:", solution.tribonacciMemoization(n))