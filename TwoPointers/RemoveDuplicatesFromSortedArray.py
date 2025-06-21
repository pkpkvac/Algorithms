from typing import List

# Remove Duplicates from Sorted Array

# You are given an integer array nums sorted in non-decreasing order. 
# Your task is to remove duplicates from nums in-place so that each element appears only once.

# After removing the duplicates, return the number of unique elements, denoted as k, 
# such that the first k elements of nums contain the unique elements.

# Note:
# - The order of the unique elements should remain the same as in the original array.
# - It is not necessary to consider elements beyond the first k positions of the array.
# - To be accepted, the first k elements of nums must contain all the unique elements.
# - Return k as the final result.

# Example 1:
# Input: nums = [1,1,2,3,4]
# Output: [1,2,3,4]
# Explanation: You should return k = 4 as we have four unique elements.

# Example 2:
# Input: nums = [2,10,10,30,30,30]
# Output: [2,10,30]
# Explanation: You should return k = 3 as we have three unique elements.

# ============================================================================
# APPROACH: TWO POINTERS
# ============================================================================

# KEY INSIGHT: Since the array is SORTED, all duplicates are adjacent!
# We can use two pointers:
# - `slow`: points to the position where the next unique element should be placed
# - `fast`: scans through the array to find the next unique element

# ALGORITHM:
# 1. Start with slow = 1 (first element is always unique)
# 2. Use fast to scan from index 1 to end
# 3. When nums[fast] != nums[fast-1], we found a new unique element
# 4. Place it at nums[slow] and increment slow
# 5. Return slow (which is the count of unique elements)

# ============================================================================
# STEP-BY-STEP TRACE
# ============================================================================

# Example: nums = [1,1,2,3,4]
#
# Initial: slow = 1, fast = 1
# nums = [1, 1, 2, 3, 4]
#         ^  ^
#      slow fast
#
# fast = 1: nums[1] = 1, nums[0] = 1 → same, skip
# fast = 2: nums[2] = 2, nums[1] = 1 → different!
#   nums[slow] = nums[2] → nums[1] = 2
#   slow++ → slow = 2
#   nums = [1, 2, 2, 3, 4]
#               ^
#            slow
#
# fast = 3: nums[3] = 3, nums[2] = 2 → different!
#   nums[slow] = nums[3] → nums[2] = 3  
#   slow++ → slow = 3
#   nums = [1, 2, 3, 3, 4]
#                  ^
#               slow
#
# fast = 4: nums[4] = 4, nums[3] = 3 → different!
#   nums[slow] = nums[4] → nums[3] = 4
#   slow++ → slow = 4
#   nums = [1, 2, 3, 4, 4]
#                     ^
#                  slow
#
# Return slow = 4
# First 4 elements: [1, 2, 3, 4] ✅

# ============================================================================
# EDGE CASES
# ============================================================================

# 1. Empty array: nums = [] → return 0
# 2. Single element: nums = [5] → return 1  
# 3. All same elements: nums = [2,2,2,2] → return 1
# 4. All unique elements: nums = [1,2,3,4] → return 4

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        # Handle edge case: empty array
        if not nums:
            return 0
        
        # slow pointer: position for next unique element
        slow = 1  # First element is always unique
        
        # fast pointer: scan through array
        for fast in range(1, len(nums)):
            # Found a new unique element
            if nums[fast] != nums[fast - 1]:
                nums[slow] = nums[fast]
                slow += 1
        
        return slow

# ============================================================================
# ALTERNATIVE APPROACH: Compare with Previous Unique
# ============================================================================

# Instead of comparing nums[fast] with nums[fast-1], 
# we can compare nums[fast] with nums[slow-1] (the last unique element we placed)

class SolutionAlternative:
    def removeDuplicates(self, nums: List[int]) -> int:
        if not nums:
            return 0
        
        slow = 1
        
        for fast in range(1, len(nums)):
            # Compare with the last unique element we placed
            if nums[fast] != nums[slow - 1]:
                nums[slow] = nums[fast]
                slow += 1
        
        return slow

# ============================================================================
# COMPLEXITY ANALYSIS
# ============================================================================

# Time Complexity: O(n) - single pass through the array
# Space Complexity: O(1) - only using two pointers, modifying in-place

# ============================================================================
# TEST THE SOLUTION
# ============================================================================

if __name__ == "__main__":
    solution = Solution()
    
    # Test case 1
    nums1 = [1, 1, 2, 3, 4]
    k1 = solution.removeDuplicates(nums1)
    print(f"Input: [1,1,2,3,4]")
    print(f"Output: k = {k1}, first {k1} elements: {nums1[:k1]}")
    print()
    
    # Test case 2  
    nums2 = [2, 10, 10, 30, 30, 30]
    k2 = solution.removeDuplicates(nums2)
    print(f"Input: [2,10,10,30,30,30]")
    print(f"Output: k = {k2}, first {k2} elements: {nums2[:k2]}")
    print()
    
    # Test case 3: All same elements
    nums3 = [5, 5, 5, 5]
    k3 = solution.removeDuplicates(nums3)
    print(f"Input: [5,5,5,5]")
    print(f"Output: k = {k3}, first {k3} elements: {nums3[:k3]}")
    print()
    
    # Test case 4: All unique elements
    nums4 = [1, 2, 3, 4, 5]
    k4 = solution.removeDuplicates(nums4)
    print(f"Input: [1,2,3,4,5]")
    print(f"Output: k = {k4}, first {k4} elements: {nums4[:k4]}") 