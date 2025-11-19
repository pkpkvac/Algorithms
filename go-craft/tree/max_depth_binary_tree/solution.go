package maxdepthbinarytree

import (
	"algorithms/tree/common"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxDepthBinaryTreeRecursive(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(MaxDepthBinaryTreeRecursive(root.Left), MaxDepthBinaryTreeRecursive(root.Right))
}

func MaxDepthBinaryTreeIterativeBFSA(root *common.TreeNode) int {

	if root == nil {
		return 0
	}
	q := []*common.TreeNode{}
	q = append(q, root)
	depth := 0
	for len(q) > 0 {
		levelSize := len(q)

		for i := 0; i < levelSize; i++ {

			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

		}
		depth++
	}

	return depth
}

type NodeDepth struct {
	node  *common.TreeNode
	depth int
}

func MaxDepthBinaryTreeIterativeBFSB(root *common.TreeNode) int {

	if root == nil {
		return 0
	}

	q := []NodeDepth{{node: root, depth: 1}}

	maxDepth := 0

	for len(q) > 0 {

		node := q[0]
		depth := node.depth
		q = q[1:]

		if node.node.Left != nil {
			q = append(q, NodeDepth{node: node.node.Left, depth: node.depth + 1})
		}
		if node.node.Right != nil {
			q = append(q, NodeDepth{node: node.node.Right, depth: node.depth + 1})
		}

		maxDepth = max(maxDepth, depth)
	}

	return maxDepth

}
