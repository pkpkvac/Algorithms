from typing import Optional

from definition import TreeNode, arrayToTree, treeToArray

# Time complexity
# best case (balanced) - O(logn)
# worst (skewed) - O(n)
# general O(h), h = height

# Space complexity
# Recursive is O(h) due to call stack -> best case O(logn) -> worst case O(n)
# Iterative is O(1) constant space, no recursion stack

class Solution:
    def insertIntoBSTRecursive(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        # Base case: create and RETURN the new node

        if root is None:
            return TreeNode(val)
    
        # Recursive case: which direction to go?
        if val > root.val:
            root.right = self.insertIntoBST(root.right, val)
        else:
            root.left = self.insertIntoBST(root.left, val)
        
        return root
    
    def insertIntoBSTIterative(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:

        if root is None:
            return TreeNode(val)
        
        current = root
        while True:
            if val > current.val:
                if current.right is None:
                    # insert here
                    current.right = TreeNode(val)
                    break
                else:
                    current = current.right
            else:
                if current.left is None:
                    # insert here
                    current.left = TreeNode(val)
                    break
                else: 
                    current = current.left
        
        return root
    


nodes = [5,3,6,None,4,None,10,None,None,7]
root = arrayToTree(nodes)

val = 9

solution = Solution()
result = solution.insertIntoBST(root, val)

print(treeToArray(result))
