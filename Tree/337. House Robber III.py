from typing import Optional
from Tree.definition import arrayToTree, TreeNode, treeToArray

# Two strategies:
# Strategy 1
# 1. Postorder with DP
# At each node, we track two states:
# - Maximum money if we rob this node
# - Maximum money if we don't rob this node

# Process:
# 1. Get results from left child
# 2. Get results from right child
# 3. At current node:
#    - If we rob it: node_value + (don't rob left) + (don't rob right)
#    - If we skip it: max(rob left, skip left) + max(rob right, skip right)

# ////

# Strategy 2
# 2. Postorder with Memoization
# At each node, we track:
# - Have we seen this state before? (node + can_rob combination)
# - Can we rob this node? (based on if we robbed parent)

# Process:
# 1. Check if we've seen this state
# 2. If we can rob current node:
#    - Try robbing it (must skip children)
#    - Try skipping it (can rob children)
# 3. Store best result for this state



class Solution:
    def rob_postorder_depth_first_search_and_dp(self, root: Optional[TreeNode]) -> int:

        def dfs(node):
            if not node:
            #   base case
                return [0, 0]
        
            left = dfs(node.left)
            right = dfs(node.right)

            # left[1] and right[1] are the skip values, so you rob this node, skip the
            rob = node.val + left[1] + right[1]
            skip = max(left[0], left[1]) + max(right[0], right[1])

            return [rob, skip]

        return max(dfs(root))
        


    def rob_postorder_depth_first_search_and_memoization(self, root: Optional[TreeNode]) -> int:
        # The DP version is:
        # More intuitive
        # Naturally fits the tree structure
        # Already efficient (visits each node once)
        # Doesn't need extra space for memo
        # Memoization would help if:
        # We had overlapping subproblems
        # We were doing top-down recursion
        # We needed to cache intermediate results
       print('test') 

tree_list = [3,2,3,None,3,None,1]

root = arrayToTree(tree_list)

# print(treeToArray(root))


solution = Solution()
print(solution.rob_postorder_depth_first_search_and_dp(root))