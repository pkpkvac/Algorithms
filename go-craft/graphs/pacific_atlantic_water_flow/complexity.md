# Complexity Analysis: Pacific Atlantic Water Flow

## Problem Approach

Multi-source BFS starting from ocean borders, working backward to find which cells can reach each ocean.

**Key insight**: Instead of checking "can this cell reach both oceans?" from every cell (expensive), we ask "which cells can each ocean reach?" by flowing upward (height ≥ current).

**Algorithm**:
1. BFS from Pacific borders (top row + left column) - mark all reachable cells
2. BFS from Atlantic borders (bottom row + right column) - mark all reachable cells
3. Return intersection of both sets

## Time Complexity

- **Best case**: O(m × n)
- **Average case**: O(m × n)
- **Worst case**: O(m × n)

**Analysis**:
- Each cell is visited at most twice (once per ocean's BFS)
- Adding Pacific borders: O(m + n)
- BFS from Pacific: O(m × n) worst case (visits all cells)
- Adding Atlantic borders: O(m + n)
- BFS from Atlantic: O(m × n) worst case
- Finding intersection: O(m × n) to iterate through set
- Sorting results: O(k log k) where k is number of results, k ≤ m × n

Total: O(m × n)

## Space Complexity

O(m × n)

**Analysis**:
- `pacific` set: O(m × n) worst case (if Pacific reaches all cells)
- `atlantic` set: O(m × n) worst case
- Queue: O(m × n) worst case (all cells in queue)
- Result array: O(m × n) worst case (all cells in result)

Total auxiliary space: O(m × n)

## Optimizations

**Current approach (set-based BFS)** is optimal for general case.

**Alternative approaches**:

1. **DFS instead of BFS**
   - Same time/space complexity
   - Uses recursion stack instead of queue
   - May be slightly more memory efficient in practice (no queue allocation)

2. **Boolean matrices instead of sets**
   ```go
   pacificReachable := make([][]bool, m)
   atlanticReachable := make([][]bool, m)
   ```
   - Same complexity, but may have better cache locality
   - Simpler indexing (`pacific[r][c]` vs `pacific[[2]int{r, c}]`)

3. **Naive approach (check from each cell)**
   - BFS/DFS from every cell to check if it reaches both oceans
   - Time: O(m² × n²) - much worse
   - Not viable for large grids

**Trade-offs**:
- Set approach: Clean code, handles sparse results well
- Matrix approach: Better cache performance, simpler syntax
- Both are O(m × n), choice depends on preference

**Note**: The key optimization is the "work backward from oceans" insight, which reduces time from O(m² × n²) to O(m × n).
