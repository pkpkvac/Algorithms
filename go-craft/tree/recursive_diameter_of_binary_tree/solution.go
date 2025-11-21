package diameterofbinarytree

import (
	"algorithms/tree/common"
)

func DiameterOfBinaryTreeRecursive(root *common.TreeNode) int {
	// base case
	diameter := 0
	if root == nil {
		return diameter
	}
	// diameter is the longest path between any two nodes, which may or may not pass through the root.
	// at each node, consider diameter passing through this node
	// note that it is only the height that is contributing to the diameter in
	// a given subtree
	// = leftHeight + rightHeight
	helper(root, &diameter)

	return diameter

}

func helper(root *common.TreeNode, diameter *int) int {

	if root == nil {
		return 0
	}

	// get the leftHeight
	leftHeight := helper(root.Left, diameter)
	rightHeight := helper(root.Right, diameter)

	*diameter = max(*diameter, leftHeight+rightHeight)

	return 1 + max(leftHeight, rightHeight)

}
