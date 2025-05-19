from collections import deque
import heapq
import math
from typing import *

# Given a 32-bit unsigned integer n, reverse the bits of the binary representation of n and return the result.

# Example 1:

# Input: n = 00000000000000000000000000010101

# Output:    2818572288 (10101000000000000000000000000000)
# Explanation: Reversing 00000000000000000000000000010101, which represents the unsigned integer 21,
# gives us 10101000000000000000000000000000 which represents the unsigned integer 2818572288.

# Let's take a simple example: n = 43 (binary: 00000000000000000000000000101011)
# After reversing: 11010100000000000000000000000000 (decimal: 3,556,769,792)

# Putting It All Together
# In our bit reversal algorithm:
# (n >> i) shifts the i-th bit to the rightmost position
# (n >> i) & 1 extracts just that bit (0 or 1)
# bit << (31 - i) moves the extracted bit to its reversed position
# result |= (bit << (31 - i)) adds this bit to our result
# Visual Example
# For n = 43 and i = 3:
# n = 43:                 00000000000000000000000000101011
# n >> 3:                 00000000000000000000000000000101
# (n >> 3) & 1:           00000000000000000000000000000001 (bit = 1)
# bit << (31 - 3):        10000000000000000000000000000000
# result |= above:        result now has a 1 at position 28


class Solution:
    def reverseBits(self, n: int) -> int:
        result = 0
    
        # Process all 32 bits
        for i in range(32):
            # Extract the i-th bit from the right
            bit = (n >> i) & 1
            
            # Place this bit at position (31-i) from the right
            result |= (bit << (31 - i))
        
        return result

solution = Solution()