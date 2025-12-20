package lowestcommonancestorinbst

import (
	"algorithms/tree/common"
)

func lowestCommonAncestor(root *common.TreeNode, p *common.TreeNode, q *common.TreeNode) *common.TreeNode {

	node := root

	// node will be the tracking node for the LCA

	for node != nil {

		if p.Val > node.Val && q.Val > node.Val {
			// go right
			node = node.Right
		} else if p.Val < node.Val && q.Val < node.Val {
			// go left
			node = node.Left
		} else {
			// there's a separation, can return node
			return node
		}

	}

	return node
}
