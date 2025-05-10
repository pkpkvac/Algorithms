from typing import *
from definition import arrayToTree, TreeNode, treeToArray

# You are given the root of a binary tree root. Invert the binary tree and return its root.

class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:

        def swap(node: Optional[TreeNode]):
            if not node:
                return
            
            tempLeft = node.left
            node.left = node.right
            node.right = tempLeft

            # One small note: When you swap the children and then recursively
            #  process them, you're actually processing the original right
            #  child when you call swap(node.left) and the original left
            #  child when you call swap(node.right).
            #  This is correct and works perfectly fine

            swap(node.left)
            swap(node.right)
        
        swap(root)
        return root
            


# Input: root = [1,2,3,4,5,6,7]
# Output: [1,3,2,7,6,5,4]

root = arrayToTree([1,2,3,4,5,6,7])

solution = Solution()

result = solution.invertTree(root)

print(treeToArray(result))