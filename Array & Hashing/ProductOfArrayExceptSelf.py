# Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].

# Each product is guaranteed to fit in a 32-bit integer.

# Follow-up: Could you solve it in
# O(n) time without using the division operation?

# Example 1:

# Input: nums = [1,2,4,6]

# Output: [48,24,12,8]
# Example 2:

# Input: nums = [-1,0,1,2,3]

# Output: [0,-6,0,0,0]

from typing import List


class Solution:
    def productExceptSelfBruteForce(self, nums: List[int]) -> List[int]:

        result = [0] * len(nums)

        for i in range(len(nums)):

            amount = None

            for j in range(len(nums)):

                if i != j:
                    if amount == None:
                        amount = nums[j]
                    else:
                        amount *= nums[j]

            result[i] = amount

        print(result)
        return result
    
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        
        # Step 1: Build prefix array (product of elements BEFORE each index)
        prefix = [1] * n
        for i in range(1, n):
            prefix[i] = prefix[i-1] * nums[i-1]
        
        # Step 2: Build suffix array (product of elements AFTER each index)
        suffix = [1] * n
        for i in range(n-2, -1, -1):
            suffix[i] = suffix[i+1] * nums[i+1]
        
        # Step 3: Multiply prefix[i] * suffix[i] for final result
        result = [0] * n
        for i in range(n):
            result[i] = prefix[i] * suffix[i]
        
        return result



solution = Solution()

nums = [1,2,4,6]

solution.productExceptSelfBruteForce(nums)
print("Optimized solution:", solution.productExceptSelf(nums))

