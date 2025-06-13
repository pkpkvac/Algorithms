from typing import List, Optional

from definition import TreeNode, arrayToTree

# You are given the root of a binary tree.
#  Return only the values of the nodes that are
#  visible from the right side of the tree,
#  ordered from top to bottom.



class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        
        q = [root]
        res = []
        while q:
            level_size = len(q)

            for i in range(level_size):

                node = q.pop(0)

                if i == level_size - 1:
                    res.append(node.val)
                
                if node.left: q.append(node.left)
                if node.right: q.append(node.right)
        print(res)
        return res

# root = [1,2,3]
# Output: [1,3]

root_arr = [1,2,3,4,5,6,7]
# Output: [1,3,7]

root = arrayToTree(root_arr)

solution = Solution()

solution.rightSideView(root=root)