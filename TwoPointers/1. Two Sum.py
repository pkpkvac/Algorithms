from typing import List

class Solution:
    def twoSumNaive(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]

    def twoSumHashMapNaive(self, nums: List[int], target: int) -> List[int]:
        hash_map = {}
        for i in range(len(nums) ):
            if hash_map.get(nums[i]) is None:
                hash_map[nums[i]] = [i]
            else:
                hash_map[nums[i]].append(i)

        for num_1, indices in  hash_map.items():
            required_value = target - num_1
            if required_value == num_1 and len(indices) > 1:
                return [indices[0], indices[1]]
            if hash_map.get(required_value):
                required_index = hash_map.get(required_value)
                if required_index[0] is not indices[0]:
                    return [indices[0], required_index[0]]

                
    def twoSumHashMapOptimized(self, nums: List[int], target: int) -> List[int]:
        # build a hash map as you go
        seen = {}  # num -> index
        for i, num in enumerate(nums):
            complement = target - num
            
            # Check if complement exists before adding current number
            if complement in seen:
                return [seen[complement], i]
                
            # Add current number to hash map
            seen[num] = i
        
        return []  # No solution found
            

# nums = [2,7,11,15, 11]
# target = 9
nums = [3,2,4]
target = 6 

sol = Solution()
# print(sol.twoSumNaive(nums, target))
# print(sol.twoSumHashMapNaive(nums, target))