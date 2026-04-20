package deletenodebst

import (
	"algorithms/tree/common"
)

func deleteNode(root *common.TreeNode, key int) *common.TreeNode {

	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	// there are 3 deletion cases here:
	// 1. no children, return nil

	if root.Right == nil && root.Left == nil {
		return nil
	}

	// 2. one child, return that child (replacing deleted node)
	if root.Left == nil {

		return root.Right

	}

	if root.Right == nil {

		return root.Left
	}

	// 3. two children:
	// replace with inorder successor (smallest value in right subtree),
	// then delete that successor from the right subtree.
	successor := root.Right
	for successor.Left != nil {
		successor = successor.Left
	}
	root.Val = successor.Val
	root.Right = deleteNode(root.Right, successor.Val)

	// Alternative (equivalent) approach using inorder predecessor:
	//
	// predecessor := root.Left
	// for predecessor.Right != nil {
	// 	predecessor = predecessor.Right
	// }
	// root.Val = predecessor.Val
	// root.Left = deleteNode(root.Left, predecessor.Val)

	return root
}
