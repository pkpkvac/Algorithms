# Complexity Analysis: Permutations

## Problem Approach

Backtracking using the **Choose → Explore → Unchoose** pattern to generate all permutations of distinct integers.

**Key mechanism**: Use a `used []bool` array to track which elements are already included in the current permutation path, preventing re-use within a single permutation.

**Algorithm**:
1. Base case: When `len(current) == len(nums)`, we have a complete permutation → copy and add to results
2. Recursive case: Try adding each unused element
   - **Choose**: Append element to current path, mark as used
   - **Explore**: Recursively build remaining permutation
   - **Unchoose**: Remove element, mark as unused (backtrack to explore other paths)

## Time Complexity

- **Best case**: O(n! × n)
- **Average case**: O(n! × n)
- **Worst case**: O(n! × n)

**Analysis**:

The algorithm generates **all n! permutations**, and each permutation requires **O(n) work** to copy it to the result.

**Why n! permutations?**
- First position: n choices
- Second position: n-1 choices (one element used)
- Third position: n-2 choices
- ...
- Last position: 1 choice

Total permutations = n × (n-1) × (n-2) × ... × 1 = **n!**

**Why × n?**
- Each complete permutation requires O(n) to copy the slice to results
- Total: n! permutations × O(n) copy = **O(n! × n)**

**Note**: Some analyses write this as O(n · n!) or simply O(n!), but O(n! × n) is more precise.

## Space Complexity

O(n)

**Analysis** (not counting output):
- **Recursion depth**: O(n) - one level per element position
- **`current` array**: O(n) - holds current permutation being built
- **`used` array**: O(n) - tracks which indices are used
- **Stack frames**: O(n) - one frame per recursion level

**Output space** (not included in auxiliary space):
- Result contains n! permutations, each of size n
- Total output: O(n! × n)

## Optimizations

**Current approach** is optimal for generating all permutations - you must visit all n! states.

**Alternative approaches**:

1. **Heap's Algorithm**
   - Generates permutations via in-place swapping
   - Time: O(n!) - slightly better constant (no copying)
   - Space: O(n) for recursion only
   - More complex to understand, same asymptotic complexity

2. **Swap-based backtracking** (instead of `used` array)
   ```go
   // Instead of maintaining used[], swap elements
   for i := startIdx; i < len(nums); i++ {
       swap(nums, startIdx, i)
       backtrack(nums, startIdx+1)
       swap(nums, startIdx, i) // restore
   }
   ```
   - Saves space from `used` array
   - Less intuitive - modifies input during recursion
   - Same time complexity

3. **Iterative with Johnson-Trotter**
   - Generates permutations iteratively
   - Time: O(n!)
   - More complex implementation
   - Good for generating permutations one-by-one

**Trade-offs**:
- **Choose-Explore-Unchoose pattern** (current): Most intuitive, clearest code
- **Swap-based**: Saves O(n) space, but muddier logic
- **Heap's/Johnson-Trotter**: Better constants, but complexity to understand

For interview/learning: Current approach is best - it clearly demonstrates the backtracking paradigm and is easy to explain.

**Note**: Cannot do better than O(n!) since there are n! permutations to generate. Any "optimization" only improves constant factors, not asymptotic complexity.
