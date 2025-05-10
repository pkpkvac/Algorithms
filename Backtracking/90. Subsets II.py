from typing import List
# Given an array nums of unique integers, return all possible subsets of nums.
# The solution set must not contain duplicate subsets. You may return the solution in any order.

# Example 1:

# Input: nums = [1,2,3]
# Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

# Example 2:

# Input: nums = [7]
# Output: [[],[7]]

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:

        solutions = []
        curr = []

        def backtrack(i):

            if i >= len(nums):
                # COPY because lists are mutable
                solutions.append(curr.copy())
                return
        
            # add choice
            # recurse with new state
            curr.append(nums[i])
            backtrack(i + 1)

            # undo choice
            # recurse with new state
            curr.pop()
            backtrack(i + 1)

        backtrack(0)
        return solutions


nums = [1,2,3]
solution = Solution()
print(solution.subsets(nums))


# Time complexity
# Why 2ⁿ?
# For each element, we have 2 choices:
# - Include it
# - Don't include it


# All possible subsets:
# [], [1], [2], [3], [1,2], [1,3], [2,3], [1,2,3]
# Total: 2³ = 8 subsets