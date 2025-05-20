

from typing import List


# Given an array of integers nums, return the length of the
#  longest consecutive sequence of elements that can be formed.

# A consecutive sequence is a sequence of elements
#  in which each element is exactly 1 greater than
#  the previous element. The elements do not have
#  to be consecutive in the original array.

# You must write an algorithm that runs in O(n) time.

# Example 1:
# Input: nums = [2,20,4,10,3,4,5]
# Output: 4
# Explanation: The longest consecutive sequence is [2, 3, 4, 5].

# Example 2:
# Input: nums = [0,3,2,5,4,6,1,1]
# Output: 7

class Solution:

    def longestConsecutiveBruteForce(self, nums: List[int]) -> int:
        return 0


    def longestConsecutive(self, nums: List[int]) -> int:
        # Edge case: empty array
        if not nums:
            return 0
        
        # Convert array to set for O(1) lookups
        num_set = set(nums)
        
        longest_streak = 0
        
        # Check each number that could be the start of a sequence
        for num in num_set:
            # Only consider numbers that are the start of a sequence
            # (i.e., num-1 is not in the set)
            if num - 1 not in num_set:
                current_num = num
                current_streak = 1
                
                # Count consecutive numbers
                while current_num + 1 in num_set:
                    current_num += 1
                    current_streak += 1
                    
                # Update longest streak
                longest_streak = max(longest_streak, current_streak)
    
        return longest_streak

solution =Solution() 

# nums = [0,3,2,5,4,6,1,1]
nums = [2,20,4,10,3,4,5]

print(solution.longestConsecutive(nums))