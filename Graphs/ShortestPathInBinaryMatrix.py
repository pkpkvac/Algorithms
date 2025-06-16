from typing import List
from collections import deque
# Given an n x n binary matrix grid, return the 
# length of the shortest clear path in the matrix.
# If there is no clear path, return -1.

# A clear path in a binary matrix is a path
# from the top-left cell (i.e., (0, 0)) to the
# bottom-right cell (i.e., (n - 1, n - 1)) such that:

# All the visited cells of the path are 0.
# All the adjacent cells of the path are 8-directionally 
# connected (i.e., they are different and
#             they share an edge or a corner).
# The length of a clear path is the number of
# visited cells of this path.


# ============================================================================
# BFS MENTAL FRAMEWORK - STEP BY STEP
# ============================================================================

# CHUNK 1: BFS IS ITERATIVE WITH A QUEUE (NOT RECURSIVE)
# ═══════════════════════════════════════════════════════════════════════════
# Think of BFS like ripples in a pond:
# - Drop a stone at START position
# - Wave 1: All positions 1 step away  
# - Wave 2: All positions 2 steps away
# - Wave 3: All positions 3 steps away...
# - STOP when wave hits the TARGET

# We use a QUEUE because:
# - Process wave 1 completely before wave 2
# - FIFO (First In, First Out) = process in distance order
# - NOT recursive - we iterate through waves

# CHUNK 2: WHAT DO WE TRACK/MEMOIZE?
# ═══════════════════════════════════════════════════════════════════════════
# We DON'T memoize like DP problems!
# Instead we track:
# 1. VISITED SET: "Have I been to (x,y) before?"
# 2. QUEUE: "What positions should I explore next?"
# 3. DISTANCE PARAMETER: "How many steps to get here?"

# CHUNK 3: HOW DO WE GUARANTEE SHORTEST PATH?
# ═══════════════════════════════════════════════════════════════════════════
# The magic of BFS ordering:
# - We explore distance 1, then 2, then 3...
# - FIRST time we reach target = guaranteed shortest!
# - Why? We've already explored ALL shorter paths

# Example: Start at (0,0), target at (2,2)
# Wave 1 (distance 1): Explore all neighbors of (0,0)
# Wave 2 (distance 2): Explore all neighbors of wave 1 positions  
# Wave 3 (distance 3): If we reach (2,2) here, we know it's shortest
#   because we already checked if (2,2) was reachable in 1 or 2 steps

# CHUNK 4: THE CORE ALGORITHM PATTERN
# ═══════════════════════════════════════════════════════════════════════════
# 1. START: Put starting position in queue with distance = 1
# 2. VISITED: Mark starting position as visited
# 3. LOOP while queue not empty:
#    a. Get next position from queue (oldest first - FIFO)
#    b. If this is target → return distance (DONE!)
#    c. Look at all neighbors
#    d. For each valid, unvisited neighbor:
#       - Mark as visited (prevent revisiting)
#       - Add to queue with distance + 1
# 4. If queue empty and no target found → return -1

# KEY INSIGHT: The distance parameter "travels" with each position
# We don't calculate distance - we carry it forward: current_distance + 1

# Shortest Path in Binary Matrix is NOT a good
# candidate for Dijkstra - it's perfect for BFS.


class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        n = len(grid)
        
        # Quick edge cases
        if grid[0][0] == 1 or grid[n-1][n-1] == 1:
            return -1
        if n == 1:
            return 1
        
        # Helper function - clean way to avoid "dumb bounds checking"
        def is_valid(row, col):
            return (0 <= row < n and 
                    0 <= col < n and 
                    grid[row][col] == 0)
        
        # BFS setup
        # initialize queue using the (row, col, distance tuple)
        queue = deque([(0, 0, 1)])  # (row, col, distance)

        # initialize a visited set as an array of tuples
        visited = set([(0, 0)])
        
        # 8 directions: all surrounding cells (including diagonals)
        directions = [(-1,-1), (-1,0), (-1,1), (0,-1), (0,1), (1,-1), (1,0), (1,1)]
        
        # Main BFS loop - NO double for loop needed!
        # We process the queue (which grows as we explore)
        while queue:

            # pop queue, extract row, col
            row, col, distance = queue.popleft()
            
            # Check if we reached destination
            if row == n - 1 and col == n - 1:
                # return distance
                return distance
            
            # Single for loop: explore all 8 neighbors
            for dr, dc in directions:

                # find new_row, new_col by adjusting for dr, dc
                new_row = row + dr
                new_col = col + dc

                # Clean condition using helper function
                # ensure new_row, new_col hasn't been visited
                if is_valid(new_row, new_col) and (new_row, new_col) not in visited:
                    # add new_row, new_col to visited,
                    visited.add((new_row, new_col))
                    # add new_row, new_col, new distance to queue
                    queue.append((new_row, new_col, distance + 1))
        
        return -1  # No path found

# ============================================================================
# LOOP STRUCTURE EXPLANATION
# ============================================================================
# NO double for loop needed! Here's why:
# 
# ❌ WRONG thinking: "I need to visit every cell, so double for loop"
# ✅ CORRECT thinking: "I need to explore neighbors of reachable cells"
#
# Structure:
# - WHILE loop: Process queue until empty (BFS main loop)
# - FOR loop: Check 8 neighbors of current cell
# 
# The queue grows as we discover new reachable cells
# We don't visit every cell - only reachable ones!


# Input: grid = [[0,1],[1,0]]
# Output: 2
# Example 2:


# Input:
#  grid = [[0,0,0],[1,1,0],[1,1,0]]
# Output: 4

# Input: grid = [[1,0,0],[1,1,0],[1,1,0]]
# Output: -1


solution = Solution()

# grid = [[0,1],[1,0]]
grid = [[0,0,0],[1,1,0],[1,1,0]]
# grid = [[1,0,0],[1,1,0],[1,1,0]]

print(solution.shortestPathBinaryMatrix(grid=grid))