from collections import deque
from typing import *
from definition import arrayToTree, TreeNode

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None: return []

        queue = deque()
        queue.append(root)
        result = []
        while len(queue) > 0:
            levelSize =len(queue) 
            tempArr = []
            for i in range(levelSize):
                node = queue.popleft()
                tempArr.append(node.val)
                if(node.left):
                    queue.append(node.left)
                if(node.right):
                    queue.append(node.right)
            result.append(tempArr)
        
        return result

            

    

# Input: root = [1,2,3,4,5,6,7]
# Output: [[1],[2,3],[4,5,6,7]]
root = arrayToTree([1,2,3,4,5,6,7])

solution = Solution()

result = solution.levelOrder(root)

print(result)