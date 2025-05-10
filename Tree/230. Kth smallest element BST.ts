// Given the root of a binary search tree, and an
// integer k, return the kth smallest value
// (1-indexed) of all the values of the nodes in the tree.

import { TreeNode, arrayToTree } from "./LevelOrderToTree";

function kthSmallest(root: TreeNode, k: number): number {
	let count = 0; // Global count to track total nodes seen

	function dfs(node: TreeNode | null): number | null {
		if (!node) return null;

		// Process left
		const left = dfs(node.left);
		if (left !== null) return left;

		// Process current
		count++; // Increment count when we process a node
		if (count === k) return node.val;

		// Process right
		return dfs(node.right);
	}

	return dfs(root) ?? 0;
}

function kthSmallestYours(root: TreeNode, k: number): number {
	function dfs(node: TreeNode | null): {
		count: number;
		result: number | null;
	} {
		if (!node) return { count: 0, result: null };
		// in order

		// 1. process left subtree
		const left = dfs(node.left);

		// if found in left, propagate solution
		if (left.result !== null) return left;

		// 2. process current node

		const newCount = left.count + 1;

		if (newCount === k) {
			return { count: newCount, result: node.val };
		}

		// 3. process right subtree
		const right = dfs(node.right);

		if (right.result !== null) {
			return { count: left.count + 1 + right.count, result: right.result };
		}

		// 4. combine counts if no result found yet
		return {
			count: left.count + 1 + right.count,
			result: null,
		};
	}

	const { result } = dfs(root);

	return result ?? 0;
}

// Input: root = [3,1,4,null,2], k = 1
// Output: 1

// Input: root = [5,3,6,2,4,null,null,1], k = 3
// Output: 3

const root = [5, 3, 6, 2, 4, null, null, 1];
const k = 3;

const node = arrayToTree(root);

console.log(node);

console.log(kthSmallest(node, k));
