from typing import List

# Insert Interval
# You are given an array of non-overlapping intervals intervals where intervals[i] = [start_i, end_i] represents the start and the end time of the ith interval. intervals is initially sorted in ascending order by start_i.
# You are given another interval newInterval = [start, end].
# Insert newInterval into intervals such that intervals is still sorted in ascending order by start_i and also intervals still does not have any overlapping intervals. You may merge the overlapping intervals if needed.
# Return intervals after adding newInterval.
# Note: Intervals are non-overlapping if they have no common point. For example, [1,2] and [3,4] are non-overlapping, but [1,2] and [2,3] are overlapping.

# Example 1:

# Input: intervals = [[1,3],[4,6]], newInterval = [2,5]
# Output: [[1,6]]

# Example 2:

# Input: intervals = [[1,2],[3,5],[9,10]], newInterval = [6,7]
# Output: [[1,2],[3,5],[6,7],[9,10]]



class Solution:
    def insertLinear(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        # first iterate through where second element of intervals is less than the first element of newInterval
        # eg intervals [1,2], newInterval [3,4]
        res = []
        i = 0
        n = len(intervals)

        while i < n and intervals[i][1] < newInterval[0]:
            res.append(intervals[i])
            i += 1
        
        # now extend the new interval
        while i < n and intervals[i][0] <= newInterval[1]:
            newInterval[0] = min(newInterval[0], intervals[i][0])
            newInterval[1] = max(newInterval[1], intervals[i][1])
            i += 1

        res.append(newInterval)
        
        #  append the remaining items
        while i < n:
            res.append(intervals[i])
            i += 1

        print(res)
        return res

        return []        
    def insertBinarySearch(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        return []
    def insertGreedy(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        return []



intervals = [[1,3],[4,6]]
newInterval = [2,5]

solution = Solution()
solution.insertLinear(intervals, newInterval)
solution.insertBinarySearch(intervals, newInterval)
solution.insertGreedy(intervals, newInterval)