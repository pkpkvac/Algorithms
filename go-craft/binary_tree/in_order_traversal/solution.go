package inordertraversal

import (
	"algorithms/binary_tree/common"
)

func InorderTraversalRecursive(root *common.TreeNode) []int {

	// left, current, right
	if root == nil {
		return []int{}
	}

	result := InorderTraversalRecursive(root.Left)
	result = append(result, root.Val)
	result = append(result, InorderTraversalRecursive(root.Right)...)

	return result

}

func InorderTraversalIterative(root *common.TreeNode) []int {

	stack := []*common.TreeNode{}
	current := root
	res := []int{}

	for current != nil || len(stack) > 0 {
		// go left as far as possible
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		// pop
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// visit
		res = append(res, top.Val)

		// move to right subtree
		current = top.Right
	}
	return res

}
