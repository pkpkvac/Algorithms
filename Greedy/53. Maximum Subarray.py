from typing import List

# 7 ways to solve (neetcode), will do 3:
# *** 1. Brute force
# *** 2. Kadane's algorithm (for returning JUST THE VALUE of the max sub array)
    # - is sophisticated
    # - deceptively 'easy' looking
    # - implicitly uses a sliding window
# (later) 3. Recursion
# 4. Dynamic programming
# *** 5. Two pointer solution (in the case you're required to return the indices
# or the array that is the max sub array)

class Solution:
    def maxSubArrayNaive(self, nums: List[int]) -> int:

        left = 0
        right = 0        
        total_max = nums[0]        

        for i in range(len(nums)):

            curr_sub_max = 0

            for  j in range(i, len(nums)):
                curr_sub_max += nums[j]

                if curr_sub_max >= total_max:
                    total_max = max(curr_sub_max, total_max)
                    left = i
                    right = j
            
        print(nums[left:right + 1])
        return nums[left:right + 1]                

    def maxSubArrayKadanes(self, nums: List[int]) -> int:

        # implicity uses a sliding window
        # only returns the value, not the array

        maxSum = nums[0]
        currSum = 0

        for i in range(len(nums)):
            
            currSum += nums[i]

            maxSum = max(currSum, maxSum)

            # ensure that currSum is > 0, otherwise drop and restart
            # if it's ever negative, may as well restart
            currSum = max(0, currSum)

        print(maxSum)

        return maxSum

    def maxSubArraySlidingWindow(self, nums: List[int]) -> int:

        maxSum = nums[0]
        currSum = 0
        left = 0
        right = 0

        for i in range(len(nums)):
            
            currSum += nums[i]

            if currSum > maxSum:            
                right = i
                maxSum = max(currSum, maxSum)

            # ensure that currSum is > 0, otherwise drop and restart
            # if it's ever negative, may as well restart
            if currSum <= 0:
                left = i + 1
                currSum = max(0, currSum)

        print(nums[left:right + 1])
        return nums[left:right + 1]  


solution = Solution()

nums = [2,-3,4,-2,2,1,-1,4]

# solution.maxSubArrayNaive(nums)
# solution.maxSubArrayKadanes(nums)
solution.maxSubArraySlidingWindow(nums)