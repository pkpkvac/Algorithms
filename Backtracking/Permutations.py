from typing import List

# Given an array nums of unique integers, return all the possible permutations. You may return the answer in any order.

# Example 1:

# Input: nums = [1,2,3]

# Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
# Example 2:

# Input: nums = [7]

# Output: [[7]]

# def backtrack(state):
#     if is_valid_solution(state):
#         # 1. Base Case
#         solutions.append(state.copy())  # Save solution
#         return

#     for choice in get_possible_choices(state):
#         # 2. Make a choice
#         state.add(choice)

#         # 3. Recurse with new state
#         backtrack(state)

#         # 4. Undo choice (backtrack)
#         state.remove(choice)

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        

        permutations = []
        current_permutation = []

        def backtrack(current_permutation):

            # base case, index == len(nums)
            # implies you've reached a permutation
            if len(current_permutation) == len(nums):
                permutations.append(current_permutation.copy())
                return
            
            # This differs from generating the subsets, or power set.
            # here, you will need to try each unused element at this current
            # position

            for num in nums:
                #skip used elements, or
                if num not in current_permutation: 
                    current_permutation.append(num) # choose 
                    backtrack(current_permutation) # recurse
                    current_permutation.pop() # undo
            
        backtrack(current_permutation)
        return permutations
    
nums = [1,2,3]

solution = Solution()
print(solution.permute(nums))