import { arrayToTree, TreeNode } from "./definition";

function maxDepthRecursive(root: TreeNode | null): number {
	if (!root) return 0;

	let depth = dfs(root);

	function dfs(node: TreeNode | null): number {
		if (!node) return 0;

		let left = dfs(node.left);

		let right = dfs(node.right);

		return Math.max(left, right) + 1;
	}

	return depth;
}

function maxDepthIterativeBFS(root: TreeNode | null): number {
	// this utilizes a level order traversal, as it uses a Queue and is a BFS technique

	if (!root) return 0;

	const queue: TreeNode[] = [root];
	let depth = 0;
	while (queue.length) {
		const levelSize = queue.length;
		for (let i = 0; i < levelSize; i++) {
			const node = queue.shift();

			if (node?.left) queue.push(node.left);
			if (node?.right) queue.push(node.right);
		}
		depth++;
	}
	return depth;
}

function maxDepthIterativeDFS(root: TreeNode | null): number {
	if (!root) return 0;

	const stack: [TreeNode, number][] = [[root, 1]];

	let maxDepth = 0;

	while (stack.length) {
		const [node, depth] = stack.pop()!;

		// update maxDepth if currDepth is greater
		maxDepth = Math.max(maxDepth, depth);

		if (node.right) stack.push([node.right, depth + 1]);
		if (node.left) stack.push([node.left, depth + 1]);
	}
	return maxDepth;
}

const root = [3, 9, 20, null, null, 15, 7];

const root_node = arrayToTree(root);

// console.log(maxDepthRecursive(root_node));
// console.log(maxDepthIterativeBFS(root_node));
console.log(maxDepthIterativeDFS(root_node));
