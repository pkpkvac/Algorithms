from typing import List

class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        hash_map = {}
        
        for num in nums:
            if num in hash_map: return True
            hash_map[num] = 1
        
        return False

    def hasDuplicateBruteForce(self, nums: List[int]) -> bool:
    #   1.
    #   for i in range(len(nums) ):
    #         j = i + 1
    #         while j < len(nums):
    #             if nums[i] == nums[j]: return True
    #             j += 1
    #     return False
    #   2.
    #   for i in range(len(nums)):
    #       for j in range(i+1, len(nums)):
    #         if nums[i] == nums[j]:
    #             return True
    #   return False
    #   3.
        for i, num in enumerate(nums):
            for j in range(i + 1, len(nums)):
                if num == nums[j]: return True
        return False

    def hasDuplicateUsingSet(self, nums: List[int]) -> bool:
        return len(set(nums)) != len(nums)

    

nums = [1,2,3,3]
# nums = [1, 2, 3, 4]

solution = Solution()
# result = solution.hasDuplicate(nums)
# result = solution.hasDuplicateBruteForce(nums)
result = solution.hasDuplicateUsingSet(nums)
print(result)