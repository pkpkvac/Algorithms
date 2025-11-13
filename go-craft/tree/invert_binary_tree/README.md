Do a "refactor pass" — implement all 4 approaches (post-order recursive/iterative, pre-order recursive/iterative) in one focused session

The Core Operation
The fundamental operation in inverting a binary tree is swapping the left and right children of each node. This is actually quite simple:

For a given node:
const temp = node.left;
node.left = node.right;
node.right = temp;

Traversal Approaches

You can use any traversal method to visit each node and perform this swap:

1. Recursive Approaches
   a. Preorder (Root, Left, Right)
   Swap the children first
   Then recursively process the (new) left and right subtrees

   b. Postorder (Left, Right, Root)
   Recursively process the left and right subtrees first
   Then swap the children

   c. Inorder (Left, Root, Right)
   Recursively process the left subtree
   Swap the children
   Recursively process the (new) right subtree (which was originally the left)

2. Iterative Approaches
   a. Using a Stack (DFS)
   Push nodes onto a stack
   Pop each node, swap its children, and push the children onto the stack

   b. Using a Queue (BFS/Level Order)
   Enqueue nodes level by level
   Dequeue each node, swap its children, and enqueue the children

Given the root of a binary tree, invert the tree, and return its root.
