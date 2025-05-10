import { arrayToTree } from "./LevelOrderToTree";

// For determining if a tree is balanced (heights of left and right subtrees differ by at most 1),

// Example trees:
//    3                  1
//  /   \              /   \
// 9    20            2     2
// /     \          /       \
//15      7        3         3
// 							 /           \
// 							4             4
// ✓ Balanced       ✗ Not Balanced

class TreeNode {
	val: number;
	left: TreeNode | null;
	right: TreeNode | null;
	constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
		this.val = val === undefined ? 0 : val;
		this.left = left === undefined ? null : left;
		this.right = right === undefined ? null : right;
	}
}

function isBalanced(root: TreeNode | null): boolean {
	const height = getHeight(root);

	return height !== -1;

	function getHeight(node: TreeNode | null): number {
		if (!node) return 0;

		const leftHeight = getHeight(node.left);

		const rightHeight = getHeight(node.right);

		// if abs(leftHeight - rightHeight) > 1, imbalanced, return -1
		if (
			Math.abs(leftHeight - rightHeight) > 1 ||
			leftHeight === -1 ||
			rightHeight === -1
		)
			return -1;

		return Math.max(leftHeight, rightHeight) + 1;
	}
}

const root = [3, 9, 20, null, null, 15, 7];

const root_node = arrayToTree(root);

console.log(isBalanced(root_node));
