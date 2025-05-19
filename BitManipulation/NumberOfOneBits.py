from collections import deque
import heapq
import math
from typing import *


# You are given an unsigned integer n. Return the number of 1 bits in its binary representation.

# You may assume n is a non-negative integer which fits within 32-bits.

# Example 1:

# Input: n = 00000000000000000000000000010111

# Output: 4
# Example 2:

# Input: n = 01111111111111111111111111111101

# Output: 30

# Hint 1
# The given integer is a 32-bit integer. Can you think of using bitwise operators to iterate through its bits? Maybe you should consider iterating 32 times.


# Hint 2
# We iterate 32 times (0 to 31) using index i. The expression (1 << i) creates a bitmask with a set bit at the i-th position. How can you check whether the i-th bit is set in the given number? Maybe you should consider using the bitwise-AND ("&").


# Hint 3
# Since the mask has a set bit at the i-th position and all 0s elsewhere, we can perform a bitwise-AND with n. If n has a set bit at the i-th position, the result is positive; otherwise, it is 0. We increment the global count if the result is positive and return it after the iteration.


class Solution:
    def hammingWeightIterative(self, n: int) -> int:

        # Visualization of this approach:
        # For n = 11 (binary: 1011):
        # i=0: mask = 1 (binary: 0001), n & mask = 1 (true) → count = 1
        # i=1: mask = 2 (binary: 0010), n & mask = 2 (true) → count = 2
        # i=2: mask = 4 (binary: 0100), n & mask = 0 (false) → count = 2
        # i=3: mask = 8 (binary: 1000), n & mask = 8 (true) → count = 3
        # i=4-31: all result in 0 → count remains 3

        count = 0

        for i in range(32):
            # create mask w/ only ith bit set to 1
            # this moves with each iteration
            mask = 1 << i

            # determine if ith bit of n is 1, if it is, increment
            if n & mask != 0:
                count += 1

        return count

    def hammingWeightBitManipulation(self, n: int) -> int:
        # more efficient

        # Visualization of this approach:
        # For n = 11 (binary: 1011):
        # n = 1011, n-1 = 1010, n & (n-1) = 1010 → count = 1
        # n = 1010, n-1 = 1001, n & (n-1) = 1000 → count = 2
        # n = 1000, n-1 = 0111, n & (n-1) = 0000 → count = 3
        # n = 0000, loop ends → return count = 3

        count = 0
        while n != 0:
            # clear rightmost bit
            n &= (n-1)
            count += 1

        return count

    def hammingWeight(n: int) -> int:
        return bin(n).count('1')

n = 11 #binary representation is 00000000000000000000000000001011


solution = Solution()

solution.hammingWeightIterative(n)