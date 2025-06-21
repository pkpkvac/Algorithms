# Given an m x n matrix of integers matrix, 
# if an element is 0, set its entire row and column to 0's.

# You must update the matrix in-place.

# Follow up: Could you solve it using O(1) space?

# Input: matrix = [
#   [0,1],
#   [1,0]
# ]

# Output: [
#   [0,0],
#   [0,0]
# ]

# Input: matrix = [
#   [1,2,3],
#   [4,0,5],
#   [6,7,8]
# ]

# Output: [
#   [1,0,3],
#   [0,0,0],
#   [6,0,8]
# ]

# ============================================================================
# MULTIPLE APPROACHES TO SOLVE THIS PROBLEM
# ============================================================================

# YOUR INSIGHT IS CORRECT! 
# ❌ Can't set zeros immediately → cascade effect makes entire matrix zero
# ✅ Must record locations first, then apply changes

# APPROACH 1: STORE COORDINATES (What you suggested)
# ═══════════════════════════════════════════════════════════════════════════
# 1. Find all zero positions and store them: [(1,1), (0,2), ...]
# 2. For each stored position, set entire row and column to zero
# 3. Time: O(mn), Space: O(mn) worst case (if many zeros)

# APPROACH 2: BOOLEAN ARRAYS (More efficient storage)
# ═══════════════════════════════════════════════════════════════════════════
# 1. Create zero_rows[m] and zero_cols[n] boolean arrays
# 2. Scan matrix: if matrix[i][j] == 0, mark zero_rows[i] = zero_cols[j] = True
# 3. Second pass: if zero_rows[i] OR zero_cols[j], set matrix[i][j] = 0
# 4. Time: O(mn), Space: O(m + n)

# APPROACH 3: USE MATRIX ITSELF AS STORAGE (O(1) space!) 
# ═══════════════════════════════════════════════════════════════════════════
# BRILLIANT INSIGHT: Use first row and first column as the "boolean arrays"!
# 
# Strategy:
# - matrix[0][j] stores whether column j should be zero
# - matrix[i][0] stores whether row i should be zero
# - Special handling for first row/column since they're used as markers
#
# This is the O(1) space solution!

# ============================================================================
# WHY THE O(1) APPROACH IS CLEVER
# ============================================================================

# The key insight: "Where can I store information without using extra space?"
# Answer: Use the matrix itself!
#
# Example: matrix = [[1,2,3], [4,0,5], [6,7,8]]
#
# Step 1: Use first row/column as markers
# If matrix[i][j] == 0, set:
# - matrix[i][0] = 0  (mark row i for zeroing)
# - matrix[0][j] = 0  (mark column j for zeroing)
#
# Step 2: Handle first row/column specially
# Since we're using them as markers, we need separate flags:
# - first_row_zero: should first row be zeroed?
# - first_col_zero: should first column be zeroed?
#
# Step 3: Apply the zeros based on markers
# - If matrix[i][0] == 0, zero out row i
# - If matrix[0][j] == 0, zero out column j
# - Apply first_row_zero and first_col_zero last

# ============================================================================
# STEP-BY-STEP TRACE: O(1) APPROACH
# ============================================================================

# Original: [[1,2,3], [4,0,5], [6,7,8]]
#
# Step 1: Scan and mark
# - Found 0 at (1,1) 
# - Set matrix[1][0] = 0 (mark row 1)
# - Set matrix[0][1] = 0 (mark column 1)
# - Matrix becomes: [[1,0,3], [0,0,5], [6,7,8]]
#
# Step 2: Apply based on markers
# - matrix[1][0] == 0 → zero row 1: [[1,0,3], [0,0,0], [6,7,8]]
# - matrix[0][1] == 0 → zero col 1: [[1,0,3], [0,0,0], [6,0,8]]
#
# Final: [[1,0,3], [0,0,0], [6,0,8]] ✅

# ============================================================================
# IMPLEMENTATION HINTS (Choose your approach!)
# ============================================================================

# APPROACH 1 (Store coordinates):
# 1. zeros = [(i,j) for i in range(m) for j in range(n) if matrix[i][j] == 0]
# 2. For each (i,j) in zeros: zero out row i and column j

# APPROACH 2 (Boolean arrays):  
# 1. zero_rows = [False] * m, zero_cols = [False] * n
# 2. Find zeros and mark: zero_rows[i] = zero_cols[j] = True
# 3. Apply: if zero_rows[i] or zero_cols[j]: matrix[i][j] = 0

# APPROACH 3 (O(1) space):
# 1. Check if first row/col need zeroing (separate flags)
# 2. Use first row/col as markers for other rows/cols
# 3. Apply zeros based on markers
# 4. Handle first row/col last

from typing import List


class Solution:
    def setZeroesNaive(self, matrix: List[List[int]]) -> None:
        
        zeros = []

        #  SPACE O(m*n) in worst case
        for r in range (0, len(matrix)):

            for c in range(0, len(matrix[r])):

                if matrix[r][c] == 0:

                    zeros.append((r,c))

        # TIME O(m * n)
        for r, c in zeros:
            print(r, c)

            # zero across the row, by going across the columns
            for col in range(0, len(matrix[r])):
                matrix[r][col] = 0

            # zero across column, by going vertically through rows
            for row in range(0, len(matrix)):
                matrix[row][c] = 0

        print(matrix)

        return 

# ============================================================================
# ANALYSIS OF YOUR NAIVE SOLUTION
# ============================================================================

# ✅ YOUR SOLUTION IS ABSOLUTELY CORRECT! 
#
# WHY IT WORKS:
# 1. You correctly identified the cascade problem
# 2. You avoid it by collecting ALL zero positions FIRST
# 3. Then you apply the zeroing operations
# 4. No cascade because original zeros are already recorded
#
# COMPLEXITY ANALYSIS:
# Time: O(m*n) to find zeros + O(k*(m+n)) to zero rows/cols
#       where k = number of original zeros
#       Worst case: O(m*n + m*n*(m+n)) = O(m²n + mn²)
#       Best case (few zeros): Close to O(m*n)
#
# Space: O(k) where k = number of zeros
#        Worst case: O(m*n) if entire matrix is zeros
#        Best case: O(1) if no zeros
#
# CORRECTNESS: Perfect! Handles all edge cases correctly.

# ============================================================================
# HINTS FOR THE O(1) SPACE SOLUTION
# ============================================================================

# Now for the "brain" solution! Key insights:

# 🧠 INSIGHT 1: Where can you store "row i needs zeroing" without extra space?
# Answer: Use matrix[i][0] (first column) as a flag!

# 🧠 INSIGHT 2: Where can you store "column j needs zeroing" without extra space?  
# Answer: Use matrix[0][j] (first row) as a flag!

# 🧠 INSIGHT 3: But what about the first row and first column themselves?
# Problem: You're using them as flags, but they might need zeroing too!
# Answer: Use separate boolean variables for first row/column

# 🧠 INSIGHT 4: Order matters! 
# You must process the first row/column LAST
# Why? Because you're using them as flags until the very end

# The algorithm becomes:
# 1. Check: Does first row have any zeros? (first_row_zero flag)
# 2. Check: Does first column have any zeros? (first_col_zero flag)  
# 3. Scan rest of matrix: Use first row/col as flags
# 4. Apply zeros based on flags (skip first row/col for now)
# 5. Finally: Apply first_row_zero and first_col_zero

    def setZeroesOptimal(self, matrix: List[List[int]]) -> None:
        """
        O(1) space solution using first row and first column as markers
        """
        rows, cols = len(matrix), len(matrix[0])
        
        # STEP 1: Check if first row and first column need to be zeroed
        # We need these flags because we'll use first row/col as markers
        first_row_zero = any(matrix[0][j] == 0 for j in range(cols))
        first_col_zero = any(matrix[i][0] == 0 for i in range(rows))
        
        # STEP 2: Use first row and first column as markers
        # Scan the matrix (excluding first row/column) 
        for i in range(1, rows):
            for j in range(1, cols):
                if matrix[i][j] == 0:
                    # Mark this row and column for zeroing
                    matrix[i][0] = 0  # Mark row i (using first column)
                    matrix[0][j] = 0  # Mark column j (using first row)
        
        # STEP 3: Zero out marked rows and columns
        # Process everything EXCEPT first row and first column
        for i in range(1, rows):
            for j in range(1, cols):
                if matrix[i][0] == 0 or matrix[0][j] == 0:
                    matrix[i][j] = 0
        
        # STEP 4: Handle first row and first column last
        # (We saved this for last because we were using them as markers)
        
        # Zero out first row if needed
        if first_row_zero:
            for j in range(cols):
                matrix[0][j] = 0
        
        # Zero out first column if needed
        if first_col_zero:
            for i in range(rows):
                matrix[i][0] = 0

    def setZeroesOptimalWithTrace(self, matrix: List[List[int]]) -> None:
        """
        Same algorithm but with detailed tracing for learning
        """
        rows, cols = len(matrix), len(matrix[0])
        print(f"Original matrix: {matrix}")
        
        # STEP 1: Check first row and column
        first_row_zero = any(matrix[0][j] == 0 for j in range(cols))
        first_col_zero = any(matrix[i][0] == 0 for i in range(rows))
        print(f"First row needs zeroing: {first_row_zero}")
        print(f"First column needs zeroing: {first_col_zero}")
        
        # STEP 2: Mark using first row/column
        print("\nStep 2: Marking phase...")
        for i in range(1, rows):
            for j in range(1, cols):
                if matrix[i][j] == 0:
                    print(f"  Found 0 at ({i},{j}) - marking row {i} and column {j}")
                    matrix[i][0] = 0  # Mark row
                    matrix[0][j] = 0  # Mark column
        
        print(f"After marking: {matrix}")
        
        # STEP 3: Apply zeros based on markers
        print("\nStep 3: Applying zeros...")
        for i in range(1, rows):
            for j in range(1, cols):
                if matrix[i][0] == 0 or matrix[0][j] == 0:
                    print(f"  Zeroing ({i},{j}) because row {i} or column {j} is marked")
                    matrix[i][j] = 0
        
        print(f"After applying markers: {matrix}")
        
        # STEP 4: Handle first row/column
        print("\nStep 4: Handling first row and column...")
        if first_row_zero:
            print("  Zeroing first row")
            for j in range(cols):
                matrix[0][j] = 0
        
        if first_col_zero:
            print("  Zeroing first column")
            for i in range(rows):
                matrix[i][0] = 0
        
        print(f"Final result: {matrix}")

solution = Solution()

# Test with your example
matrix1 = [
    [1,2,3],
    [4,0,5],
    [6,7,8]
]

print("=== TESTING OPTIMAL SOLUTION ===")
solution.setZeroesOptimalWithTrace(matrix1)

# Test with another example
matrix2 = [
    [0,1,2,0],
    [3,4,5,2],
    [1,3,1,5]
]

print("\n=== TESTING ANOTHER EXAMPLE ===")
solution.setZeroesOptimalWithTrace(matrix2)