package maxpathsum

import (
	"algorithms/tree/common"
	"math"
)

func maxPathSum(root *common.TreeNode) int {

	maxSum := math.MinInt32

	var dfs func(node *common.TreeNode) int

	dfs = func(node *common.TreeNode) int {

		// base case
		if node == nil {
			return 0
		}

		// postorder traversal
		// get contribution from children, ignore negatives
		// i.e., worst you take from subtree is nothing.
		left := max(0, dfs(node.Left))

		right := max(0, dfs(node.Right))

		// update global max (path that peaks here)
		maxSum = max(maxSum, left+right+node.Val)

		// return extendable path (only one direction)
		// what can be offered upward
		return node.Val + max(left, right)
	}

	dfs(root)

	return maxSum

}

func maxPathSumBruteForce(root *common.TreeNode) int {

	// important technique. Fundamental.

	return 0
}
