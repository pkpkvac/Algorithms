from collections import deque
import heapq
import math
from typing import *


# Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],
# [0,1,1,0,1,0,0,0,0,0,0,0,0],
# [0,1,0,0,1,1,0,0,1,0,1,0,0],
# [0,1,0,0,1,1,0,0,1,1,1,0,0],
# [0,0,0,0,0,0,0,0,0,0,1,0,0],
# [0,0,0,0,0,0,0,1,1,1,0,0,0],
# [0,0,0,0,0,0,0,1,1,0,0,0,0]]
# Output: 6
# Explanation: The answer is not 11, because the island must be connected 4-directionally.

class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        # Initialize visited set inside the method
        self.visited = set()
        max_area = 0
        
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == 1 and (i, j) not in self.visited:
                    # Start a new DFS for each unvisited land cell
                    current_area = self.dfs(grid, i, j)
                    max_area = max(max_area, current_area)
        
        return max_area

    def dfs(self, grid: List[List[int]], i: int, j: int) -> int:
        rows, cols = len(grid), len(grid[0])
        
        # Check boundaries and if it's water or already visited
        if (i < 0 or i >= rows or 
            j < 0 or j >= cols or 
            grid[i][j] == 0 or 
            (i, j) in self.visited):
            return 0
        
        # Mark as visited
        self.visited.add((i, j))
        
        # Current cell contributes 1 to area
        area = 1
        
        # Add areas from all four directions
        area += self.dfs(grid, i+1, j)
        area += self.dfs(grid, i-1, j)
        area += self.dfs(grid, i, j+1)
        area += self.dfs(grid, i, j-1)
                    
        return area
    


solution = Solution()
grid = [
[0,1,1,0,1],
[1,0,1,0,1],
[0,1,1,0,1],
[0,1,0,0,1]
]

print(solution.maxAreaOfIsland(grid))



# grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]

# Input: grid = [
#   [0,1,1,0,1],
#   [1,0,1,0,1],
#   [0,1,1,0,1],
#   [0,1,0,0,1]
# ]

# Output: 6