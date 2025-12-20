package subtreeofanothertree

import (
	"algorithms/tree/common"
)

func isSubtree(root *common.TreeNode, subRoot *common.TreeNode) bool {

	// base case, tree has been exhausted
	if root == nil {
		return false
	}

	isSubroot := false

	// case in which it makes sense to recurse on the subtrees
	if isSameTree(root, subRoot) {
		return true
	}

	// otherwise, explore the tree
	isSubroot = isSubtree(root.Left, subRoot)

	if isSubroot {
		return true
	}

	isSubroot = isSubtree(root.Right, subRoot)

	return isSubroot

}

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
