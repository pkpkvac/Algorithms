package valid_binary_search_tree

import (
	"algorithms/tree/common"
	"math"
)

func isValidBST(root *common.TreeNode) bool {

	if root == nil {
		return true
	}
	return helper(root, math.MinInt, math.MaxInt)

}

func helper(root *common.TreeNode, min int, max int) bool {

	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}
