from collections import deque
import heapq
import math
from typing import *

# Given an array of meeting time interval objects consisting of
#  start and end times [[start_1,end_1],[start_2,end_2],...] (start_i < end_i),
#  determine if a person could add all meetings to their schedule without any conflicts.

# Example 1:

# Input: intervals = [(0,30),(5,10),(15,20)]

# Output: false
# Explanation:

# (0,30) and (5,10) will conflict
# (0,30) and (15,20) will conflict
# Example 2:

# Input: intervals = [(5,8),(9,15)]

# Output: true
# Note:

# (0,8),(8,10) is not considered a conflict at 8

class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end

class Solution:
    def canAttendMeetings(self, intervals: List[Interval]) -> bool:
        
        if len(intervals) == 0: return True
        
        intervals.sort(key=lambda x: x.start)
        
        end = intervals[0].end
        # naive would be a double for loop
        for i in range(1, len(intervals)):
            start = intervals[i].start

            if end > start: return False

            end = intervals[i].end
        
        return True

solution = Solution()

# intervals = [Interval(0,30),Interval(5,10),Interval(15,20)]
intervals = [Interval(0,8),Interval(8,10)]

print(solution.canAttendMeetings(intervals))