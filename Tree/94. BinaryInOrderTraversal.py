# pyright: reportGeneralTypeIssues=false
from typing import Optional, List

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def inorderTraversalRecursive(self, root: Optional[TreeNode]) -> List[int]:
        result = []

        def recursiveHelper(root: Optional[TreeNode], result_list: List[int]):
            if root is not None:
                recursiveHelper(root.left, result_list)
                result.append(root.val)
                recursiveHelper(root.right, result_list)

        recursiveHelper(root, result)
        print(result)        
        return result

    def inorderTraversalStack(self, root: Optional[TreeNode]) -> List[int]:
        result = []
        stack  = []
        curr = root
        
        while stack or curr:
            while curr:
                stack.append(curr)
                curr = curr.left

            if stack:
                parent = stack.pop()
                result.append(parent.val)
                curr = parent.right
        return result
    

# test
root = TreeNode(1, None, TreeNode(2, TreeNode(3, None, None), None))
solution = Solution()
# print(solution.inorderTraversalRecursive(root))
print(solution.inorderTraversalStack(root))


