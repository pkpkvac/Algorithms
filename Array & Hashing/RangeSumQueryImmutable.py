from typing import List
# Given an integer array nums, handle
# multiple queries of the following type:

# Calculate the sum of the elements of nums
# between indices left and right inclusive
#  where left <= right.

# Implement the NumArray class:

# NumArray(int[] nums) Initializes
#  the object with the integer array nums.
# int sumRange(int left, int right) Returns
# the sum of the elements of nums between
# indices left and right inclusive
# (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

# Example 1:

# Input
# ["NumArray", "sumRange", "sumRange", "sumRange"]
# [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
# Output
# [null, 1, -1, -3]

# Explanation
# NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
# numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
# numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
# numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

class NumArray:

    nums = []
    # def __init__(self, nums: List[int]):
    #     # Naive init
    #     self.nums = nums

    def __init__(self, nums: List[int]):
        # Prefix sum init
        sum = 0
        for i in range(0, len(nums)):
            sum += nums[i]
            self.nums.append(sum)
            # [-2, 0, 3, -5, 2, -1] -> [-2, -2, 1, -4, -2, -3]
        print(self.nums)


    # def sumRangeNaive(self, left: int, right: int) -> int:
    #     # complexity is 
    #     sum = 0 
    #     for i in range(left, right + 1):
    #         sum += self.nums[i]

    #     print(sum)
    #     return sum

    def sumRange(self, left: int, right: int) -> int:
        if left == 0:
            return self.nums[right]
        else:
            return self.nums[right] - self.nums[left - 1]
    

# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)

nums = [-2, 0, 3, -5, 2, -1]
left = 2
right = 5

obj = NumArray(nums)

# param_1 = obj.sumRangeNaive(left,right)
param_1 = obj.sumRange(left,right)
print(param_1)
