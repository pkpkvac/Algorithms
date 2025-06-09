from typing import List

# Given an integer array arr, return the 
# length of a maximum size turbulent subarray of arr.

# A subarray is turbulent if the comparison sign
#  flips between each adjacent pair of elements in the subarray.

# More formally, a subarray [arr[i], arr[i + 1], ...,
#  arr[j]] of arr is said to be turbulent if and only if:

# For i <= k < j:
# arr[k] > arr[k + 1] when k is odd, and
# arr[k] < arr[k + 1] when k is even.
# Or, for i <= k < j:
# arr[k] > arr[k + 1] when k is even, and
# arr[k] < arr[k + 1] when k is odd.


# Example 1:

# Input: arr = [9,4,2,10,7,8,8,1,9]
# Output: 5
# Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]
# Example 2:

# Input: arr = [4,8,12,16]
# Output: 2
# Example 3:

# Input: arr = [100]
# Output: 1

class Solution:

   
    def maxTurbulenceSize(self, arr: List[int]) -> int:

        def compare(a, b):
            if a < b: return -1
            if a > b: return 1
            return 0

        if len(arr) < 2:
            return len(arr)
        
        prev_dir = compare(arr[0], arr[1])

        if prev_dir != 0:
            curr_length = 2
        else:
            curr_length = 1

        max_turbulence_size = curr_length

        for i in range(2, len(arr)):

            curr_dir = compare(arr[i-1], arr[i])

            if curr_dir != 0 and curr_dir != prev_dir:
                # turbulence continues
                curr_length += 1
            elif curr_dir == 0:
                max_turbulence_size = max(max_turbulence_size, curr_length)
                curr_length = 1
            else:
                max_turbulence_size = max(max_turbulence_size, curr_length)
                curr_length = 2
                # reset
            prev_dir = curr_dir

        max_turbulence_size = max(max_turbulence_size, curr_length)
        return max_turbulence_size

# arr = [9,4,2,10,7,8,8,1,9]
arr = [4,8,12,16]
solution = Solution()
print(solution.maxTurbulenceSize(arr))