from typing import *
from definition import arrayToTree, TreeNode, treeToArray


# The diameter of a binary tree is defined as the length of the longest path between any two nodes within the tree. The path does not necessarily have to pass through the root.

# The length of a path between two nodes in a binary tree is the number of edges between the nodes.

# Given the root of a binary tree root, return the diameter of the tree.

# Input: root = [1,null,2,3,4,5]

# Output: 3
# Explanation: 3 is the length of the path [1,2,3,5] or [5,3,2,4].

# Key insight: The diameter passing through any node equals the sum of
# the heights of its left and right subtrees (the longest path from left to right through that node).

class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        
        # we need to define a dfs function that returns the SUM
        # of the left and right heights for each node, and
        # retain a maximum to return
        self.maxDiameter = 0
        def dfs(node: Optional[TreeNode]) -> int:
            if not node:
                return 0
            
            left = dfs(node.left)
            right = dfs(node.right)

            self.maxDiameter = max(self.maxDiameter, left + right)
            # it is tempting to do this, but it's wrong, we 
            # want to return the height of the subtree
            # return left + right + 1
            return max(left, right) + 1
        
        dfs(root)
        return self.maxDiameter





root = [1,None,2,3,4,5]

tree = arrayToTree(root)

print(tree)

solution = Solution()

print(solution.diameterOfBinaryTree(tree))
        

