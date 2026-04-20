# Complexity Analysis: Islands and Treasure

## Problem Approach

Multi-source BFS starting from all treasure chests simultaneously, spreading outward to fill land cells with distances.

**Key insight**: Instead of BFS from each land cell to find nearest treasure (expensive), start from all treasures at once and expand outward. First visit to any cell = shortest distance.

**Algorithm**:
1. First pass: Collect all treasure positions (value = 0) into queue
2. BFS: Process queue, expanding to unvisited land cells (INF)
3. Update each land cell with `current_distance + 1`

## Time Complexity

- **Best case**: O(m × n)
- **Average case**: O(m × n)
- **Worst case**: O(m × n)

**Analysis**:
- First pass to collect treasures: O(m × n)
- BFS processes each cell at most once: O(m × n)
- Each cell explores 4 neighbors: O(4 × m × n) = O(m × n)

Total: O(m × n)

## Space Complexity

O(m × n)

**Analysis**:
- Queue: O(m × n) worst case (all cells could be in queue)
- No additional data structures needed (grid is modified in-place)

Total auxiliary space: O(m × n) for the queue

## Optimizations

**Current approach (multi-source BFS)** is optimal.

**Alternative approaches**:

1. **Naive approach (BFS from each land cell)**
   - For each land cell, BFS to find nearest treasure
   - Time: O(k × m × n) where k = number of land cells, potentially O(m² × n²)
   - Space: O(m × n) for BFS queue
   - Much slower - would timeout on large grids

2. **BFS from each treasure individually**
   - BFS from each treasure, update cells if new distance < existing
   - Time: O(t × m × n) where t = number of treasures
   - Space: O(m × n) for BFS queue
   - Better than naive but still worse than multi-source

3. **DFS (not recommended)**
   - DFS doesn't guarantee shortest paths in first visit
   - Would need to try all paths and take minimum
   - Much less efficient than BFS for shortest path problems

**Trade-offs**:
- Multi-source BFS visits each cell exactly once (optimal)
- No need to track/compare distances - first visit is guaranteed shortest
- Works efficiently even with multiple treasures

**Key optimization**: The "all sources at once" approach reduces time from O(k × m × n) or O(t × m × n) to O(m × n), where each cell is processed exactly once.
