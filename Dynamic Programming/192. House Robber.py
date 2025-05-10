from typing import List

# dynamic programming - recurrence relation

class Solution:
    def rob(self, nums: List[int]) -> int:
       # 1. Using DP array
        dp = [0] * len(nums)  # dp[i] = max money at position i
        dp[0] = nums[0]

        if len(nums) < 2: return dp[0]

        dp[1] = max(nums[0], nums[1])

        for i in range(2, len(nums)):
            dp[i] = max(
                nums[i] + dp[i-2], # rob current, rob before prior
                dp[i-1] # or rob prior
                )
            
        print(dp)

        return dp[-1]

# nums = [1,2,3,1]
nums = [2,7,9,3,1]
solution = Solution()
solution.rob(nums)
