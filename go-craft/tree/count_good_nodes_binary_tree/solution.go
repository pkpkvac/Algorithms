package countgoodnodesbinarytree

import (
	"algorithms/tree/common"
)

func goodNodes(root *common.TreeNode) int {

	if root == nil {
		return 0
	}

	count := 0

	helper(root, root.Val, &count)

	return count
}

func helper(root *common.TreeNode, maxVal int, count *int) {

	if root == nil {
		return
	}

	if root.Val >= maxVal {
		*count++
		maxVal = root.Val
	}

	helper(root.Left, maxVal, count)
	helper(root.Right, maxVal, count)
}
