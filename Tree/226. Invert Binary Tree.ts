import { TreeNode, arrayToTree } from "./LevelOrderToTree";

// The Core Operation
// The fundamental operation in inverting a binary tree is swapping the left and right children of each node. This is actually quite simple:

// For a given node:
// const temp = node.left;
// node.left = node.right;
// node.right = temp;

// Traversal Approaches

// You can use any traversal method to visit each node and perform this swap:
// 1. Recursive Approaches
// Preorder (Root, Left, Right)
// Swap the children first
// Then recursively process the (new) left and right subtrees
// Postorder (Left, Right, Root)
// Recursively process the left and right subtrees first
// Then swap the children
// Inorder (Left, Root, Right)
// Recursively process the left subtree
// Swap the children
// Recursively process the (new) right subtree (which was originally the left)

// 2. Iterative Approaches
// Using a Stack (DFS)
// Push nodes onto a stack
// Pop each node, swap its children, and push the children onto the stack
// Using a Queue (BFS/Level Order)
// Enqueue nodes level by level
// Dequeue each node, swap its children, and enqueue the children

// Given the root of a binary tree, invert the tree, and return its root.

// Example 1:
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]

// Example 2:
// Input: root = [2,1,3]
// Output: [2,3,1]
// Example 3:

// Input: root = []
// Output: []

function invertTree(root: TreeNode | null): TreeNode | null {}

// const root = [4,2,7,1,3,6,9];
const root = [2, 3, 1];
