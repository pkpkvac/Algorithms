package same_binary_tree

import (
	"algorithms/tree/common"
)

func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {

	// base cases
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
