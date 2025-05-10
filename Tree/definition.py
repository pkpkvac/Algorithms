# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def arrayToTree(arr):
    if not arr:
        return None
    
    root = TreeNode(arr[0])
    queue = [root]
    i = 1
    
    while queue and i < len(arr):
        node = queue.pop(0)
        
        if i < len(arr) and arr[i] is not None:
            node.left = TreeNode(arr[i])
            queue.append(node.left)
        i += 1
        
        if i < len(arr) and arr[i] is not None:
            node.right = TreeNode(arr[i])
            queue.append(node.right)
        i += 1
    
    return root

def treeToArray(root):
    # use a queue to turn the tree into its level order representation
    if not root:
        return [] 

    queue = []
    level = []
    queue = [root]
    
    while queue:
        node = queue.pop(0)

        if node is not None:
            level.append(node.val)
            queue.append(node.left)
            queue.append(node.right)
        else:
            level.append(None)

    # Remove trailing None values
    while level and level[-1] is None:
        level.pop()

    return level

