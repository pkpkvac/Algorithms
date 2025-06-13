

# You are given a binary tree root and an integer target,
# delete all the leaf nodes with value target.

# Note that once you delete a leaf node with value target,
#  if its parent node becomes a leaf node and has the value
#   target, it should also be deleted (you need to continue doing that until you cannot).

from typing import Optional

from definition import TreeNode, arrayToTree, treeToArray


class Solution:

    def removeLeafNodes(self, root: Optional[TreeNode], target: int) -> Optional[TreeNode]:
        # post order DFS is the best approach 

        # Why Post-Order is Perfect:
        # Post-order = Left → Right → Current
        # The "cascading deletion" pattern requires bottom-up processing:
        # Process children first → Remove any target leaf nodes
        # Then check parent → After children are gone, is parent now a leaf with target value?
        # Repeat upward → This naturally handles the cascading effect

        # By the time you process a parent node, 
        # you already know the final state of its children. 
        # This lets you accurately determine:
        # Is this node now a leaf?
        # Should this node be deleted?
        # Post-order ensures you always have complete 
        # information about the subtree below before making 
        # deletion decisions. Perfect for cascading deletions!

        if not root:
            return None
        
        # process children first (post order)
        root.left = self.removeLeafNodes(root.left, target)
        root.right = self.removeLeafNodes(root.right, target)

        if not root.left and not root.right and root.val == target:
            # deletes the node
            return None 
        
        return root


# Input:
root_list = [1,2,3,5,2,2,5]
target = 2

# Output:
# [1,2,3,5,null,null,5]


# Input:
# root_list = [3,null,3,3]
# target = 3

# Output: []

root = arrayToTree(root_list)

solution = Solution()

ans = solution.removeLeafNodes(root=root, target=target)

print(treeToArray(ans))

