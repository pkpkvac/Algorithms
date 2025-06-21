from typing import List
# Subarray Sum Equals K

# You are given an array of integers nums and an integer k,
#  return the total number of subarrays whose sum equals to k.

# A subarray is a contiguous non-empty sequence of elements within an array.

# Example 1:

# Input: nums = [2,-1,1,2], k = 2

# Output: 4
# Explanation: [2], [2,-1,1], [-1,1,2], [2] are the subarrays whose sum is equals to k.

# Example 2:

# Input: nums = [4,4,4,4,4,4], k = 4

# Output: 6

# ============================================================================
# UNDERSTANDING THE CONFUSING i-1 TERM
# ============================================================================

# YOUR CONFUSION IS TOTALLY VALID! Let me explain step by step:

# Let's use: nums = [2, -1, 1, 2]
# 
# First, let's define prefix sums clearly:
# prefix_sum[0] = nums[0] = 2                    (sum from index 0 to 0)
# prefix_sum[1] = nums[0] + nums[1] = 2 + (-1) = 1    (sum from index 0 to 1) 
# prefix_sum[2] = nums[0] + nums[1] + nums[2] = 2 + (-1) + 1 = 2    (sum from index 0 to 2)
# prefix_sum[3] = nums[0] + nums[1] + nums[2] + nums[3] = 2 + (-1) + 1 + 2 = 4    (sum from index 0 to 3)

# Now, HOW DO WE GET SUBARRAY SUMS?
# 
# Subarray from index 1 to 3: nums[1] + nums[2] + nums[3] = (-1) + 1 + 2 = 2
# 
# Using prefix sums:
# subarray_sum(1 to 3) = prefix_sum[3] - prefix_sum[0] = 4 - 2 = 2 ✅
#                                        ^^^^^^^^^^^^
#                                        This is prefix_sum[i-1] where i=1

# WHY prefix_sum[i-1]?
# Because prefix_sum[0] contains nums[0], which we DON'T want in our subarray (1 to 3)
# So we subtract it out!

# ============================================================================
# THE EDGE CASE: SUBARRAYS STARTING FROM INDEX 0
# ============================================================================

# What about subarray from index 0 to 2?
# subarray_sum(0 to 2) = nums[0] + nums[1] + nums[2] = 2 + (-1) + 1 = 2
#
# Using the formula: subarray_sum(0 to 2) = prefix_sum[2] - prefix_sum[i-1]
# where i = 0, so i-1 = -1
# 
# What is prefix_sum[-1]? This doesn't exist!
# 
# SOLUTION: Define prefix_sum[-1] = 0 (sum of no elements)
# So: subarray_sum(0 to 2) = prefix_sum[2] - prefix_sum[-1] = 2 - 0 = 2 ✅

# ============================================================================
# CONCRETE EXAMPLE WITH ALL SUBARRAYS
# ============================================================================

# nums = [2, -1, 1, 2]
# prefix_sums = [2, 1, 2, 4]
# 
# Let's check ALL possible subarrays:
# 
# Subarray [2] (index 0 to 0):
#   Direct sum: 2
#   Formula: prefix_sum[0] - prefix_sum[-1] = 2 - 0 = 2 ✅
# 
# Subarray [2, -1] (index 0 to 1):
#   Direct sum: 2 + (-1) = 1  
#   Formula: prefix_sum[1] - prefix_sum[-1] = 1 - 0 = 1 ✅
# 
# Subarray [-1] (index 1 to 1):
#   Direct sum: -1
#   Formula: prefix_sum[1] - prefix_sum[0] = 1 - 2 = -1 ✅
# 
# Subarray [-1, 1] (index 1 to 2):
#   Direct sum: -1 + 1 = 0
#   Formula: prefix_sum[2] - prefix_sum[0] = 2 - 2 = 0 ✅
# 
# Subarray [1, 2] (index 2 to 3):
#   Direct sum: 1 + 2 = 3
#   Formula: prefix_sum[3] - prefix_sum[1] = 4 - 1 = 3 ✅

# ============================================================================
# THE ALGORITHM INSIGHT
# ============================================================================

# As we process the array left to right, we maintain:
# 1. current_prefix_sum (running total)
# 2. hashmap of {prefix_sum: count} 
# 
# At each position, we ask:
# "How many previous positions had prefix_sum = (current_prefix_sum - k)?"
# 
# Those previous positions are the START points of subarrays ending at current position with sum k!

# Example trace with nums = [2, -1, 1, 2], k = 2:
# 
# Initialize: hashmap = {0: 1}  ← This handles the "prefix_sum[-1] = 0" case!
# 
# Index 0, nums[0] = 2:
#   current_prefix_sum = 2
#   Looking for: 2 - 2 = 0 in hashmap
#   Found {0: 1}, so add 1 to count  ← This finds subarray [2]
#   Update hashmap: {0: 1, 2: 1}
# 
# Index 1, nums[1] = -1:
#   current_prefix_sum = 2 + (-1) = 1  
#   Looking for: 1 - 2 = -1 in hashmap
#   Not found, so add 0 to count
#   Update hashmap: {0: 1, 2: 1, 1: 1}
# 
# Index 2, nums[2] = 1:
#   current_prefix_sum = 1 + 1 = 2
#   Looking for: 2 - 2 = 0 in hashmap  
#   Found {0: 1}, so add 1 to count  ← This finds subarray [2, -1, 1]
#   Update hashmap: {0: 1, 2: 2, 1: 1}  ← Note: 2 appears twice now!
# 
# Index 3, nums[3] = 2:
#   current_prefix_sum = 2 + 2 = 4
#   Looking for: 4 - 2 = 2 in hashmap
#   Found {2: 2}, so add 2 to count  ← This finds subarrays [-1, 1, 2] and [2]
# 
# Total count: 1 + 0 + 1 + 2 = 4 ✅

# ============================================================================
# THINKING THROUGH THE APPROACHES
# ============================================================================

# YOUR O(N²) INTUITION IS CORRECT:
# For each starting position i:
#   For each ending position j >= i:
#     Calculate sum from i to j
#     If sum == k, increment count
#
# def subarraySum(self, nums, k):
#     count = 0
#     for i in range(len(nums)):
#         current_sum = 0
#         for j in range(i, len(nums)):
#             current_sum += nums[j]
#             if current_sum == k:
#                 count += 1
#     return count

# ============================================================================
# THINKING TOWARD THE O(N) SOLUTION
# ============================================================================

# QUESTION: Can we avoid recalculating sums from scratch each time?
# 
# KEY INSIGHT: What if we kept track of "prefix sums"?
# prefix_sum[i] = sum of all elements from index 0 to i
#
# Example: nums = [2, -1, 1, 2]
# prefix_sums = [2, 1, 2, 4]
#                ^  ^  ^  ^
#                |  |  |  sum(0 to 3) = 2-1+1+2 = 4
#                |  |  sum(0 to 2) = 2-1+1 = 2  
#                |  sum(0 to 1) = 2-1 = 1
#                sum(0 to 0) = 2

# NOW THE MAGIC QUESTION:
# If I want subarray from index i to j to have sum = k,
# what relationship must exist between prefix sums?
#
# subarray_sum(i to j) = prefix_sum[j] - prefix_sum[i-1]
# We want: subarray_sum(i to j) = k
# So: prefix_sum[i-1] = prefix_sum[j] - k

# ============================================================================
# THE "AHA!" MOMENT
# ============================================================================

# As I build prefix sums from left to right:
# - At each position j, I have current prefix_sum
# - I want to know: how many previous positions had prefix_sum = (current - k)?
# - Those are exactly the starting positions that create subarrays ending at j with sum k!
#
# HINT: What data structure is great for counting occurrences? 🤔

# ============================================================================
# STEP-BY-STEP TRACE EXAMPLE
# ============================================================================

# nums = [2, -1, 1, 2], k = 2
# 
# As we process:
# Index 0: prefix_sum = 2
#   Looking for previous prefix_sum = 2 - 2 = 0
#   Count of 0? (Need to handle this edge case!)
#   
# Index 1: prefix_sum = 1  
#   Looking for previous prefix_sum = 1 - 2 = -1
#   Count of -1? 
#
# Index 2: prefix_sum = 2
#   Looking for previous prefix_sum = 2 - 2 = 0
#   Count of 0?
#
# Index 3: prefix_sum = 4
#   Looking for previous prefix_sum = 4 - 2 = 2
#   Count of 2? (We've seen this before!)

# ============================================================================
# HINTS FOR IMPLEMENTATION
# ============================================================================

# 🎯 HINT 1: Use a hashmap to count prefix sum frequencies
# 🎯 HINT 2: Initialize with {0: 1} to handle subarrays starting from index 0
# 🎯 HINT 3: For each position, check if (current_prefix_sum - k) exists in map
# 🎯 HINT 4: Add current prefix sum to map after checking
# 🎯 HINT 5: The magic is in "how many times have I seen this prefix sum before?"

class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        # Try implementing the O(n) approach using the hints above!
        # Think: prefix sums + hashmap for counting

        currSum = 0
        prefix_sum = {0: 1}  # Initialize with {0: 1} for subarrays starting from index 0
        count = 0
        
        for num in nums:
            currSum += num
            
            # Check if (currSum - k) exists in our hashmap
            # This means there's a previous prefix sum that creates a subarray with sum = k
            if (currSum - k) in prefix_sum:
                count += prefix_sum[currSum - k]
            
            # Add current prefix sum to hashmap (or increment its count)
            if currSum not in prefix_sum:
                prefix_sum[currSum] = 1
            else:
                prefix_sum[currSum] += 1

        return count
        
        # Let's trace through nums = [2, -1, 1, 2], k = 2:
        # 
        # Initialize: prefix_sum = {0: 1}, count = 0
        # 
        # num = 2: currSum = 2
        #   Check: (2 - 2) = 0 in {0: 1}? YES! count += 1 → count = 1
        #   Update: prefix_sum = {0: 1, 2: 1}
        # 
        # num = -1: currSum = 1  
        #   Check: (1 - 2) = -1 in {0: 1, 2: 1}? NO
        #   Update: prefix_sum = {0: 1, 2: 1, 1: 1}
        # 
        # num = 1: currSum = 2
        #   Check: (2 - 2) = 0 in {0: 1, 2: 1, 1: 1}? YES! count += 1 → count = 2  
        #   Update: prefix_sum = {0: 1, 2: 2, 1: 1}  # Note: 2 appears twice!
        # 
        # num = 2: currSum = 4
        #   Check: (4 - 2) = 2 in {0: 1, 2: 2, 1: 1}? YES! count += 2 → count = 4
        #   Update: prefix_sum = {0: 1, 2: 2, 1: 1, 4: 1}
        # 
        # Return count = 4 ✅
        