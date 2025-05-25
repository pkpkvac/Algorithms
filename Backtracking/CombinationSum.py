from typing import List

# Combination Sum
# You are given an array of distinct integers nums and a target
#  integer target. Your task is to return a list of all unique
#  combinations of nums where the chosen numbers sum to target.

# The same number may be chosen from nums an unlimited number of times.
#  Two combinations are the same if the frequency of each of
#  the chosen numbers is the same, otherwise they are different.

# You may return the combinations in any order and the order 
# of the numbers in each combination can be in any order.

# Example 1:

# Input: 
# nums = [2,5,6,9] 
# target = 9

# Output: [[2,2,5],[9]]
# Explanation:
# 2 + 2 + 5 = 9. We use 2 twice, and 5 once.
# 9 = 9. We use 9 once.

# Example 2:

# Input: 
# nums = [3,4,5]
# target = 16

# Output: [[3,3,3,3,4],[3,3,5,5],[4,4,4,4],[3,4,4,5]]
# Example 3:

# Input: 
# nums = [3]
# target = 5

# Output: []

class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        result = []  # Store all valid combinations
        current_path = []  # Current combination we're building
        
        def backtrack(remaining_target, start_index):
            # BASE CASE 1: Found a valid combination!
            if remaining_target == 0:
                result.append(current_path.copy()) 
                return 
            
            # BASE CASE 2: Went too far, impossible to continue
            if remaining_target < 0:
                return

            # TRY each number from start_index onwards
            for i in range(start_index, len(nums)):
                # Make choice
                current_path.append(nums[i])

                # Explore choice
                backtrack(remaining_target - nums[i], i)

                # Undo choice
                current_path.pop()
            
        
        # Start the exploration
        backtrack(target, 0)
        return result

# Test it out
solution = Solution()
nums = [2, 5, 6, 9]
target = 9
result = solution.combinationSum(nums, target)
print(f"Input: nums = {nums}, target = {target}")
print(f"Output: {result}")
        