from collections import deque
import heapq
import math
from typing import *

# You are given an array of non-overlapping intervals 
# intervals where intervals[i] = [start_i, end_i] represents the 
# start and the end time of the ith interval. intervals is initially 
# sorted in ascending order by start_i.

# You are given another interval newInterval = [start, end].

# Insert newInterval into intervals such that intervals 
# is still sorted in ascending order by start_i and also 
# intervals still does not have any overlapping intervals. 
# You may merge the overlapping intervals if needed.

# Return intervals after adding newInterval.

# Note: Intervals are non-overlapping if they have 
# no common point. For example, [1,2] and [3,4] are 
# non-overlapping, but [1,2] and [2,3] are overlapping.

# Example 1:

# Input: intervals = [[1,3],[4,6]], newInterval = [2,5]

# Output: [[1,6]]
# Example 2:

# Input: intervals = [[1,2],[3,5],[9,10]], newInterval = [6,7]

# Output: [[1,2],[3,5],[6,7],[9,10]]

class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        result = []
        i = 0
        n = len(intervals)
        
        # Part 1: Add all intervals that come before newInterval
        while i < n and intervals[i][1] < newInterval[0]:
            result.append(intervals[i])
            i += 1
        
        # Part 2: Merge overlapping intervals
        while i < n and intervals[i][0] <= newInterval[1]:
            # Update newInterval to include the current overlapping interval
            newInterval[0] = min(newInterval[0], intervals[i][0])
            newInterval[1] = max(newInterval[1], intervals[i][1])
            i += 1
        
        # Add the merged interval
        result.append(newInterval)
        
        # Part 3: Add all intervals that come after newInterval
        while i < n:
            result.append(intervals[i])
            i += 1
        
        return result
    
solution = Solution()

intervals = [[1,2],[3,5],[9,10]]

newInterval = [6,7]

solution.insert(intervals, newInterval)
