from definition import TreeNode,arrayToTree


class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        pVal = p.val
        qVal = q.val
        currNode = root

        while currNode is not None:
            currVal = currNode.val
            if pVal > currVal and qVal > currVal:
                currNode = currNode.right
            elif pVal < currVal and qVal < currVal:
                currNode = currNode.left
            else:
                return currNode
        
        return None  # This shouldn't happen in a valid BST with existing nodes

    def findNode(self, root: TreeNode, val: int) -> TreeNode:
        """Helper function to find a node with given value in the tree"""
        if not root:
            return None
        if root.val == val:
            return root
        elif val < root.val:
            return self.findNode(root.left, val)
        else:
            return self.findNode(root.right, val)

    
solution = Solution()

root = [5,3,8,1,4,7,9,None,2]
rootNode = arrayToTree(root)

# Find the actual nodes in the tree
p = solution.findNode(rootNode, 3)
q = solution.findNode(rootNode, 8)

result = solution.lowestCommonAncestor(rootNode, p, q)
print(f"LCA of {p.val} and {q.val} is: {result.val}")