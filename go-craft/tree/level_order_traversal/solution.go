package levelordertraversal

import "algorithms/tree/common"

func levelOrderTraversal(root *common.TreeNode) [][]int {

	res := [][]int{}

	if root == nil {
		return res
	}

	q := []*common.TreeNode{}

	q = append(q, root)

	for len(q) > 0 {
		arr := []int{}

		sizeLevel := len(q)

		for i := 0; i < sizeLevel; i++ {

			front := q[0]

			arr = append(arr, front.Val)

			q = q[1:]

			if front.Left != nil {
				q = append(q, front.Left)
			}

			if front.Right != nil {
				q = append(q, front.Right)
			}

		}

		res = append(res, arr)

	}

	return res
}
