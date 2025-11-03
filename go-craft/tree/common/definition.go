package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func NewTreeNodeWithLeftAndRight(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func CreateBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := NewTreeNode(nums[0])
	queue := []*TreeNode{root}
	i := 1
	for i < len(nums) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if i < len(nums) && nums[i] != -999 {
			node.Left = NewTreeNode(nums[i])
			queue = append(queue, node.Left)
		}
		i++
		if i < len(nums) && nums[i] != -999 {
			node.Right = NewTreeNode(nums[i])
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
