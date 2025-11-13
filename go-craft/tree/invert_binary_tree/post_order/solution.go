package invertbinarytree

import (
	"algorithms/tree/common"
)

func invertTreeRecursive(root *common.TreeNode) *common.TreeNode {
	// post order inversion

	if root == nil {
		return nil
	}

	invertTreeRecursive(root.Left)
	invertTreeRecursive(root.Right)

	root.Left, root.Right = root.Right, root.Left
	return root

}

func invertTreeIterative(root *common.TreeNode) *common.TreeNode {

}
