package postordertraversal

import (
	"algorithms/tree/common"
)

func PostorderTraversalRecursive(root *common.TreeNode) []int {

	// left, right, root
	if root == nil {
		return []int{}
	}

	result := PostorderTraversalRecursive(root.Left)
	result = append(result, PostorderTraversalRecursive(root.Right)...)
	result = append(result, root.Val)

	return result

}

func PostorderTraversalIterative(root *common.TreeNode) []int {

	stack := []*common.TreeNode{}
	current := root
	res := []int{}
	// Tricky part: we need to track the last visited node to avoid infinite loop
	// if the last visited node is the right child of the current node,
	// then we can visit the current node
	var lastVisited *common.TreeNode

	for current != nil || len(stack) > 0 {
		// go left as far as possible
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		top := stack[len(stack)-1]
		// lastVisited is used to track the last visited node
		// if the last visited node is the right child of the current node,
		// then we can visit the current node

		if top.Right != nil && top.Right != lastVisited {
			// move to right subtree
			current = top.Right
		} else {
			// pop and visit
			res = append(res, top.Val)
			lastVisited = top
			stack = stack[:len(stack)-1]
			current = nil
		}

	}
	return res

}
