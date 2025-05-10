# The XOR total of an array is defined as the bitwise XOR of all its elements, or 0 if the array is empty.

# For example, the XOR total of the array [2,5,6] is 2 XOR 5 XOR 6 = 1.
# Given an array nums, return the sum of all XOR totals for every subset of nums.

# Note: Subsets with the same elements should be counted multiple times.

# An array a is a subset of an array b if a can be obtained from b by deleting some (possibly zero) elements of b.

# Example 1:

# Input: nums = [1,3]
# Output: 6
# Explanation: The 4 subsets of [1,3] are:
# - The empty subset has an XOR total of 0.
# - [1] has an XOR total of 1.
# - [3] has an XOR total of 3.
# - [1,3] has an XOR total of 1 XOR 3 = 2.
# 0 + 1 + 3 + 2 = 6
# Example 2:

# Input: nums = [5,1,6]
# Output: 28
# Explanation: The 8 subsets of [5,1,6] are:
# - The empty subset has an XOR total of 0.
# - [5] has an XOR total of 5.
# - [1] has an XOR total of 1.
# - [6] has an XOR total of 6.
# - [5,1] has an XOR total of 5 XOR 1 = 4.
# - [5,6] has an XOR total of 5 XOR 6 = 3.
# - [1,6] has an XOR total of 1 XOR 6 = 7.
# - [5,1,6] has an XOR total of 5 XOR 1 XOR 6 = 2.
# 0 + 5 + 1 + 6 + 4 + 3 + 7 + 2 = 28
# Example 3:

# Input: nums = [3,4,5,6,7,8]
# Output: 480
# Explanation: The sum of all XOR totals for every subset is 480.

from collections import deque
import heapq
import math
from typing import *


class Solution:
    def calculateXOR(self, subset: List[int]) -> int:
        xorResult = 0
        for num in subset:
            xorResult ^= num
        return xorResult

    def subsetXORSum(self, nums: List[int]) -> int:

        result = []
        xorResult = []

        def backtrack(index, current_subset):
            # base case: considered all elements
            if index == len(nums):
                result.append(current_subset.copy())
                return

            # Choice 1: include the current element
            current_subset.append(nums[index])
            backtrack(index + 1, current_subset)

            # Choice 2: exclude current element (backtrack)
            current_subset.pop()
            backtrack(index + 1, current_subset)
        
        backtrack(0, [])

        # result has all subsets now, go through each of these entries, and create an xor
        print(result)

        # for each subset, calculate the xor
        for subset in result:
            xorResult.append(self.calculateXOR(subset))

        sum = 0
        for xor in xorResult:
            sum += xor

        return sum


nums = [3,4,5,6,7,8]
solution = Solution()

print(solution.subsetXORSum(nums))