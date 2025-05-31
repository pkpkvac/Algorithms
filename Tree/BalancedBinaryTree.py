from typing import List, Optional

from definition import TreeNode, arrayToTree

# Tree is balanced if:
# for each node, the heights of its left and right subtree differ by at most 1
# and the left and right subtrees are also balanced



class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:

        def checkNode(node : Optional[TreeNode]) -> tuple[bool, int]:
            # base case: node is None
            if node is None:
                return True, 0

            (is_left_balanced, left_height) = checkNode(node=node.left)

            (is_right_balanced, right_height) = checkNode(node=node.right)

            is_current_balanced = False

            if abs(left_height - right_height) <= 1 and is_left_balanced and is_right_balanced:
                is_current_balanced = True

            return (is_current_balanced, max(left_height, right_height) + 1)
        
        (is_balanced, height) = checkNode(root)
        print(is_balanced)
        return is_balanced



# nodes= [1,2,3, None,None,4]
nodes = [1,2,3,None,None,4,None,5]
root = arrayToTree(nodes)

solution = Solution()

solution.isBalanced(root)