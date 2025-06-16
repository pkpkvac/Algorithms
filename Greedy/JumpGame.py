from typing import List
# You are given an integer array nums. You are initially positioned
# at the array's first index, and each element in the
# array represents your maximum jump length at that position.

# Return true if you can reach the last index, or false otherwise.

# Example 1:

# Input: nums = [2,3,1,1,4]
# Output: true
# Explanation: Jump 1 step from index 0 to 1,
# then 3 steps to the last index.
# Example 2:

# Input: nums = [3,2,1,0,4]
# Output: false
# Explanation: You will always arrive at index 3 no matter what.
# Its maximum jump length is 0, which makes it
# impossible to reach the last index.

# ============================================================================
# ADDRESSING THE SPECIFIC "STUCK" CONCERN
# ============================================================================

# QUESTION: "What if from position 0 with value 3:
# - Jump 3 steps → land on position 3 (value 0) → STUCK!
# - Jump 2 steps → land on position 2 (value 1) → can continue
# How does greedy handle this?"

# ANSWER: GREEDY DOESN'T CHOOSE SPECIFIC JUMPS!
# 
# CRUCIAL MISUNDERSTANDING: Greedy algorithm doesn't simulate actual jumps
# It computes REACHABILITY - what positions are theoretically reachable
#
# Let's trace [3,3,1,0,4]:
#
# Position 0: value = 3
# - Greedy doesn't pick "jump 3" or "jump 2"
# - It asks: "What's the FARTHEST I can reach from here?"
# - From pos 0: can reach pos 1, 2, OR 3
# - max_reach = max(0, 0+3) = 3
# - ALL positions 0,1,2,3 are now reachable!
#
# Position 1: value = 3 (and 1 <= 3, so we can reach pos 1)
# - From pos 1: can reach pos 2, 3, OR 4  
# - max_reach = max(3, 1+3) = 4
# - Position 4 is now reachable!
#
# RESULT: Can reach the end! ✅

# ============================================================================
# WHY THIS IS GREEDY, NOT SLIDING WINDOW
# ============================================================================

# SLIDING WINDOW thinking (❌ Wrong approach):
# "Look at a window of possible jumps and find maximum"
# - This focuses on examining ranges/windows
# - Typically involves two pointers that expand/contract a window
# - Example: longest substring, maximum sum subarray

# GREEDY thinking (✅ Correct approach):
# "At each step, make the locally optimal choice"
# - Keep track of FARTHEST position reachable so far
# - The greedy choice: "Can I reach further from here?"
# - Never look back - only care about maximum reach forward

# ============================================================================
# THE GREEDY INSIGHT
# ============================================================================

# Key question: "What's the farthest I can possibly reach?"
# 
# At each position i:
# - Can I even GET to position i? (i <= max_reach?)
# - If yes: What's the farthest I can reach FROM position i?
# - Update max_reach = max(max_reach, i + nums[i])
# 
# The greedy choice: Always track the maximum reachable position
# We don't need to consider all possible jump combinations!

# ============================================================================
# WHY BACKTRACKING DOESN'T WORK WELL (OVERKILL)
# ============================================================================

# BACKTRACKING would try every possible jump sequence:
# 
# def canJumpBacktrack(nums, pos=0):
#     if pos >= len(nums) - 1: return True
#     for jump in range(1, nums[pos] + 1):
#         if canJumpBacktrack(nums, pos + jump):
#             return True
#     return False
#
# PROBLEMS:
# 1. EXPONENTIAL TIME: At each position, try all possible jumps
#    - Position 0: try 1,2,3... jumps
#    - Each leads to more branching
#    - Time: O(k^n) where k = max jump size
#
# 2. REDUNDANT WORK: Same positions visited multiple times
#    - Example: [2,3,1,1,4]
#    - Can reach position 2 from position 0 (jump 2) OR position 1 (jump 1)
#    - Backtracking explores BOTH paths unnecessarily
#
# 3. OVERKILL: We only need to know IF reachable, not HOW MANY WAYS

# ============================================================================
# WHY DP DOESN'T ADD VALUE (UNNECESSARY)
# ============================================================================

# DP would track: "Can I reach position i?"
# 
# def canJumpDP(nums):
#     dp = [False] * len(nums)
#     dp[0] = True
#     for i in range(len(nums)):
#         if dp[i]:  # If position i is reachable
#             for jump in range(1, nums[i] + 1):
#                 if i + jump < len(nums):
#                     dp[i + jump] = True
#     return dp[-1]
#
# PROBLEMS:
# 1. TIME: O(n * k) where k = average jump size (potentially O(n²))
# 2. SPACE: O(n) for DP array
# 3. REDUNDANT: If we can reach position 5, we implicitly know we can reach 1,2,3,4
#    DP calculates all these individually!

# ============================================================================
# WHY GREEDY WORKS - THE MAGIC PROPERTY
# ============================================================================

# KEY INSIGHT: This problem has MONOTONIC REACHABILITY
#
# Property: "If you can reach position A, you can reach every position ≤ A"
# 
# Example: nums = [2,3,1,1,4]
# - From position 0: Can jump 1 or 2 steps → can reach positions 1,2
# - From position 1: Can jump 1,2,3 steps → from here can reach 2,3,4
# - CRUCIAL: If max_reach = 4, we can reach positions 0,1,2,3,4
#
# This means: We don't need to track individual positions!
# Just track the FARTHEST reachable position

# ============================================================================
# DETAILED TRACE: [3,3,1,0,4]
# ============================================================================

# Position 0: nums[0] = 3, max_reach = 0
# - Can we reach pos 0? Yes (0 <= 0)
# - From pos 0, farthest we can reach: 0 + 3 = 3
# - max_reach = max(0, 3) = 3
# - Reachable positions: {0, 1, 2, 3}

# Position 1: nums[1] = 3, max_reach = 3  
# - Can we reach pos 1? Yes (1 <= 3)
# - From pos 1, farthest we can reach: 1 + 3 = 4
# - max_reach = max(3, 4) = 4
# - Reachable positions: {0, 1, 2, 3, 4} ← INCLUDES THE END!

# Position 2: nums[2] = 1, max_reach = 4
# - Can we reach pos 2? Yes (2 <= 4)  
# - From pos 2, farthest we can reach: 2 + 1 = 3
# - max_reach = max(4, 3) = 4 (no change)
# - Still reachable: {0, 1, 2, 3, 4}

# Position 3: nums[3] = 0, max_reach = 4
# - Can we reach pos 3? Yes (3 <= 4)
# - From pos 3, farthest we can reach: 3 + 0 = 3  
# - max_reach = max(4, 3) = 4 (no change)
# - Still reachable: {0, 1, 2, 3, 4}

# Position 4: Our goal! Can we reach it?
# - Is 4 <= max_reach? Yes! (4 <= 4)
# - ANSWER: TRUE ✅

# KEY INSIGHT: Greedy never gets "stuck" at position 3 because:
# 1. It doesn't simulate actual jumps
# 2. It tracks ALL reachable positions simultaneously  
# 3. Position 4 became reachable from position 1, not position 3!

# ============================================================================
# ADDRESSING THE "STUCK" CONCERN
# ============================================================================

# CONCERN: "What if greedy gets stuck and misses a better path?"
#
# ANSWER: Impossible due to monotonic property!
#
# Example: nums = [3,2,1,0,4]
# Position 0: max_reach = 0 + 3 = 3
# Position 1: max_reach = max(3, 1+2) = 3  
# Position 2: max_reach = max(3, 2+1) = 3
# Position 3: max_reach = max(3, 3+0) = 3
# Position 4: Can't reach! (4 > 3)
#
# Could ANY other strategy reach position 4?
# NO! Because:
# - To reach pos 4, you must first reach pos 3 (or closer)
# - From pos 3, you can only jump 0 steps
# - From pos 2, you can only reach pos 3 (2+1=3)
# - From pos 1, you can only reach pos 3 (1+2=3)
# - From pos 0, you can only reach pos 3 (0+3=3)
#
# If greedy can't reach pos 4, NOTHING can reach pos 4!

# GREEDY GUARANTEE: If ANY path exists to the end, greedy will find it
# Why? Because greedy always maximizes reachable positions

# ============================================================================
# GREEDY VS OTHER APPROACHES
# ============================================================================

# Why not DP? 
# - DP would ask: "How many ways to reach each position?"
# - We only care: "CAN we reach the end?" (boolean)
# - Greedy is more efficient: O(n) vs O(n²)

# Why not backtracking?
# - Would try all possible jump sequences
# - Exponential time complexity
# - Greedy insight: We only need maximum reachable position

# The beauty of greedy: Local optimal choices lead to global optimum
# "Always extend our reach as far as possible" → correct answer

# 🎯 Hint 1: Track ONE variable - the farthest you can reach
# 🎯 Hint 2: At each position, ask "Can I get here?" and "How far can I go from here?"
# 🎯 Hint 3: If you can't reach a position, you're stuck
# 🎯 Hint 4: You only need to reach the last index, not necessarily land on it
# 🎯 Hint 5: The locally optimal choice leads to the globally optimal solution

class Solution:
    def canJumpWithTrace(self, nums: List[int]) -> bool:
        print(f"Array: {nums}")
        print(f"Goal: Reach index {len(nums) - 1}")
        print()
        
        max_reach = 0
        
        for i in range(len(nums)):
            print(f"Position {i}: value={nums[i]}, max_reach={max_reach}")
            
            if i > max_reach:
                print(f"  ❌ Can't reach position {i} (beyond max_reach={max_reach})")
                return False
            
            new_reach = i + nums[i]
            if new_reach > max_reach:
                print(f"  ✅ Updated max_reach: {max_reach} → {new_reach}")
                max_reach = new_reach
            else:
                print(f"  → No improvement (can reach up to {new_reach})")
            
            if max_reach >= len(nums) - 1:
                print(f"  🎯 Can reach the end! (max_reach={max_reach})")
                return True
            
            print()
        
        return True

# nums = [3,3,1,0,4]  # Your specific example!
# nums = [2,3,1,1,4]
nums = [3,2,1,0,4]

solution = Solution()

solution.canJumpWithTrace(nums)