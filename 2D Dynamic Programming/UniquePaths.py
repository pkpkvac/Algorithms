# There is an m x n grid where you are allowed to move either down
# or to the right at any point in time.

# Given the two integers m and n, return the number of
# possible unique paths that can be taken from the
# top-left corner of the grid (grid[0][0]) 
# to the bottom-right corner (grid[m - 1][n - 1]).

# Input: m = 3, n = 6
# Output: 21
# Example 2:

# Input: m = 3, n = 3
# Output: 6

class Solution:
    def uniquePathsMine(self, m: int, n: int) -> int:

        # Time Complexity: O(2^(m+n))
        # Why? At each cell, you make 2 recursive calls (down + right):
        # Apply to UniquePaths....
        # calls
        # This creates a binary tree of depth m + n - 2 (distance from start to end).
        # Number of calls ≈ 2^(m+n-2) = O(2^(m+n))
        # Space Complexity: O(m+n) - recursion stack depth
        # For m=3, n=6: About 2^7 = 128+ calls (many duplicates!)

        # for 3x6 grid, example:
        # this solution is bad, because it has a lot
        # of duplicate work being done
        # to make it better, you can use
        # memoization, or tabulation
        # dp(0,0)
        # ├── dp(1,0)
        # │   ├── dp(2,0)
        # │   │   ├── dp(3,0) → 0
        # │   │   └── dp(2,1)
        # │   │       ├── dp(3,1) → 0
        # │   │       └── dp(2,2) ← CALCULATED AGAIN LATER!
        # │   └── dp(1,1)
        # │       ├── dp(2,1) ← DUPLICATE! Already calculated above
        # │       └── dp(1,2) ← Will be calculated many times
        # └── dp(0,1)
        #     ├── dp(1,1) ← DUPLICATE! Already calculated above
        #     └── dp(0,2)
        #         └── ... more duplicates
        #   dp(2,2) gets calculated hundreds of times!

        
        def dp(curr_m, curr_n) -> int:

            # Bases cases:
            if curr_m > m - 1 or curr_n > n - 1:
                # out of bounds, not a potential route
                return 0
            if curr_m == m - 1 and curr_n == n - 1:
                # arrived
                return 1
            
            paths_down = dp(curr_m + 1, curr_n)
            paths_right = dp(curr_m, curr_n + 1)

            return paths_down + paths_right

        return dp(0, 0)

    def uniquePathsMemoized(self, m: int, n: int) -> int:

        # Time Complexity: O(m×n)
        # Why? There are exactly m×n unique subproblems (each grid position). Each is calculated exactly once thanks to memo.
        # Space Complexity: O(m×n) - memo table + O(m+n) recursion stack = O(m×n)
        # For m=3, n=6: Exactly 18 unique calculations!

        memo = {}
        
        def dp(curr_m, curr_n) -> int:

            if(curr_m, curr_n) in memo:
                return memo[(curr_m, curr_n)]
            # Bases cases:
            if curr_m > m - 1 or curr_n > n - 1:
                # out of bounds, not a potential route
                return 0
            if curr_m == m - 1 and curr_n == n - 1:
                # arrived
                return 1
            
            paths_down = dp(curr_m + 1, curr_n)
            paths_right = dp(curr_m, curr_n + 1)

            memo[(curr_m, curr_n)] = paths_down + paths_right
            return memo[(curr_m, curr_n)]

        return dp(0, 0)

    def uniquePathsTabulation(self, m: int, n: int) -> int:

        # Tabulation: O(m×n) - great, and more space efficient in practice!

        # Step 1: Create DP table - dp[i][j] = paths from (i,j) to destination
        dp = [[0] * n for _ in range(m)]
        
        # Step 2: Base case - destination has 1 path (to itself)
        dp[m-1][n-1] = 1
        
        # Step 3: Fill table from bottom-right to top-left
        # Why this order? Because dp[i][j] depends on dp[i+1][j] and dp[i][j+1]
        for i in range(m-1, -1, -1):  # From bottom row to top
            for j in range(n-1, -1, -1):  # From right col to left
                if i == m-1 and j == n-1:
                    continue  # Skip destination (already set)
                
                # Current cell paths = paths from cell below + paths from cell right
                paths_down = dp[i+1][j] if i+1 < m else 0
                paths_right = dp[i][j+1] if j+1 < n else 0
                dp[i][j] = paths_down + paths_right
        
        # Step 4: Answer is at starting position
        return dp[0][0]

    def uniquePathsTabulationVisual(self, m: int, n: int) -> int:
        print(f"\n=== Tabulation for {m}x{n} grid ===")
        dp = [[0] * n for _ in range(m)]
        dp[m-1][n-1] = 1
        
        print("Step 1 - Initialize with base case:")
        self.print_grid(dp)
        
        step = 2
        for i in range(m-1, -1, -1):
            for j in range(n-1, -1, -1):
                if i == m-1 and j == n-1:
                    continue
                
                paths_down = dp[i+1][j] if i+1 < m else 0
                paths_right = dp[i][j+1] if j+1 < n else 0
                dp[i][j] = paths_down + paths_right
                
                print(f"Step {step} - Fill dp[{i}][{j}] = {paths_down} + {paths_right} = {dp[i][j]}")
                self.print_grid(dp)
                step += 1
        
        return dp[0][0]
    
    def print_grid(self, grid):
        for row in grid:
            print([f"{x:2}" for x in row])
        print()


# Test all approaches
m = 3
n = 3  # Smaller grid for clearer visualization
solution = Solution()

print("Original (no memo):", solution.uniquePathsMine(m, n))
print("Memoized:", solution.uniquePathsMemoized(m, n))
print("Tabulation:", solution.uniquePathsTabulation(m, n))

# Visual demonstration
solution.uniquePathsTabulationVisual(m, n)