from typing import List
# Given an array of intervals intervals where 
# intervals[i] = [start_i, end_i], 
# return the minimum number of intervals
# you need to remove to make the rest of the 
# intervals non-overlapping.

# Note: Intervals are non-overlapping
# even if they have a common point. 
# For example, [1, 3] and [2, 4] are overlapping,
# but [1, 2] and [2, 3] are non-overlapping.

# Example 1:

# Input: intervals = [[1,2],[2,4],[1,4]]
# Output: 1

# Explanation: After [1,4] is removed, the rest
#  of the intervals are non-overlapping.

# Example 2:

# Input: intervals = [[1,2],[2,4]]
# Output: 0



class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        # CRUCIAL: Sort by END time, not start time!
        # Why? Greedy choice: keep intervals that end earlier
        intervals.sort(key=lambda x: x[1])

        removed = 0
        last_end = intervals[0][1]  # End time of last kept interval
        
        # Start from second interval
        for i in range(1, len(intervals)):
            start, end = intervals[i]
            
            # Check for overlap: current starts before last one ended
            if start < last_end:
                # OVERLAP! Remove current interval (increment count)
                removed += 1
                # Keep the one with earlier end time (already sorted, so keep last_end)
                # Don't update last_end
            else:
                # NO OVERLAP! Keep this interval
                last_end = end
        
        return removed
    
    def eraseOverlapIntervalsVisual(self, intervals: List[List[int]]) -> int:
        print(f"Original: {intervals}")
        intervals.sort(key=lambda x: x[1])
        print(f"Sorted by end time: {intervals}")
        
        removed = 0
        last_end = intervals[0][1]
        kept = [intervals[0]]
        
        print(f"Keep first interval: {intervals[0]}, last_end = {last_end}")
        
        for i in range(1, len(intervals)):
            start, end = intervals[i]
            print(f"\nCheck interval {intervals[i]}:")
            print(f"  Does {start} < {last_end}? ", end="")
            
            if start < last_end:
                print("YES - OVERLAP!")
                print(f"  Remove {intervals[i]}")
                removed += 1
            else:
                print("NO - Keep it!")
                last_end = end
                kept.append(intervals[i])
                print(f"  Update last_end = {last_end}")
        
        print(f"\nFinal kept intervals: {kept}")
        print(f"Removed count: {removed}")
        return removed

# Test with examples
solution = Solution()

print("=== Example 1: [[1,2],[2,4],[1,4]] ===")
intervals1 = [[1,2],[2,4],[1,4]]
result1 = solution.eraseOverlapIntervalsVisual(intervals1.copy())
print(f"Answer: {result1}\n")

print("=== Example 2: [[1,2],[2,4]] ===")
intervals2 = [[1,2],[2,4]]
result2 = solution.eraseOverlapIntervalsVisual(intervals2.copy())
print(f"Answer: {result2}\n")

print("=== More Complex Example: [[1,4],[2,3],[3,4],[1,2]] ===")
intervals3 = [[1,4],[2,3],[3,4],[1,2]]
result3 = solution.eraseOverlapIntervalsVisual(intervals3.copy())
print(f"Answer: {result3}")

