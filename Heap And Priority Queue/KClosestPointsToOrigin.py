from collections import deque
import heapq
import math
from typing import *

# Input: points = [[0,2],[2,2]], k = 1

# Output: [[0,2]]
# Explanation : The distance between (0, 2) and the origin (0, 0) is 2. The distance between (2, 2) and the origin is sqrt(2^2 + 2^2) = 2.82842. So the closest point to the origin is (0, 2).

# Example 2:

# Input: points = [[0,2],[2,0],[2,2]], k = 2

# Output: [[0,2],[2,0]]
# Explanation: The output [2,0],[0,2] would also be accepted.

class Solution:

    def distance(self, x: int, y: int):
        return x**2 + y**2

    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        #  max heap
        min_heap = []

        # heapq.heappush
        for point in points:
            dist = self.distance(point[0], point[1])
            
            heapq.heappush(min_heap, (dist,point))

        # extract k closest points
        result = []
        for _ in range(k):
            _, point = heapq.heappop(min_heap)
            result.append(point)

        return result        

points = [[0,2],[2,0],[2,2]]
k = 2

solution = Solution()

print(solution.kClosest(points, k))