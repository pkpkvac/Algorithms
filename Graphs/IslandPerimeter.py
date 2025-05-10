from collections import deque
import heapq
import math
from typing import *

# You are given row x col grid representing a map where grid[i][j] = 1
# represents land and grid[i][j] = 0 represents water.

# Grid cells are connected horizontally/vertically (not diagonally).
# The grid is completely surrounded by water, and there is exactly one 
# island (i.e., one or more connected land cells).

# The island doesn't have "lakes", meaning the water inside isn't 
# connected to the water around the island. One cell is a square 
# with side length 1. The grid is rectangular, width and height
#  don't exceed 100. Determine the perimeter of the island.

class Solution:
    perimeterCount = 0
    visited = Set()
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        # Initialize visited set inside the method, not as class variable
        self.visited = set()
        
        startRow = None
        startCol = None
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == 1:
                    startRow = i
                    startCol = j
                    break
            if startRow is not None:
                break
                    
        if startRow is not None:
            perimeter = self.dfs(grid, startRow, startCol)
            return perimeter
        return 0

    def dfs(self, grid: List[List[int]], i: int, j: int):
        rows = len(grid)
        cols = len(grid[0])

        # Check boundaries and water cells
        if i < 0 or i >= rows or j < 0 or j >= cols or grid[i][j] == 0:
            return 1
            
        # Check if already visited
        if (i, j) in self.visited:
            return 0

        # Mark as visited
        self.visited.add((i, j))
        
        # Calculate perimeter from all four directions
        perimeter = self.dfs(grid, i+1, j) + self.dfs(grid, i-1, j) + \
                    self.dfs(grid, i, j+1) + self.dfs(grid, i, j-1)
                    
        return perimeter

    def iterative(self, grid: List[List[int]]) -> int:
        perimeter = 0
        rows = len(grid)
        cols = len(grid[0])

        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == 1:
                    # Check all 4 directions
                    
                    # Check top neighbor
                    if i == 0 or grid[i-1][j] == 0:
                        perimeter += 1
                    
                    # Check bottom neighbor
                    if i == rows-1 or grid[i+1][j] == 0:
                        perimeter += 1
                    
                    # Check left neighbor
                    if j == 0 or grid[i][j-1] == 0:
                        perimeter += 1
                    
                    # Check right neighbor
                    if j == cols-1 or grid[i][j+1] == 0:
                        perimeter += 1
                        
        return perimeter

# Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
# Output: 16

grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]

solution = Solution()

print(solution.islandPerimeter(grid))