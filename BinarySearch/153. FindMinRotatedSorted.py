from typing import List

class Solution:
    def findMin(self, nums: List[int]) -> int:
        l = 0
        r = len(nums) - 1
        while l < r:
            # min_val = nums[]
            mid = (l + r ) // 2
            if nums[mid] > nums[r]:
                # must in the right side then
                l = mid + 1
            else:
                r = mid
                
        print(nums[l])
        return nums[l]

solution = Solution()

nums = [3,4,5,1,2]
solution.findMin(nums)
