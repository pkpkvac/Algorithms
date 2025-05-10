// Definition for binary tree node
class TreeNode {
	val: number;
	left: TreeNode | null;
	right: TreeNode | null;

	constructor(
		val: number = 0,
		left: TreeNode | null = null,
		right: TreeNode | null = null
	) {
		this.val = val;
		this.left = left;
		this.right = right;
	}
}

function arrayToTree(arr: (number | null)[]): TreeNode | null {
	if (!arr.length) return null;

	const root = new TreeNode(arr[0]!); // We know first element exists
	const queue: TreeNode[] = [root];
	let i = 1;

	while (queue.length && i < arr.length) {
		const node = queue.shift()!; // We know queue has elements

		// Handle left child
		if (i < arr.length && arr[i] !== null) {
			node.left = new TreeNode(arr[i]!);
			queue.push(node.left);
		}
		i++;

		// Handle right child
		if (i < arr.length && arr[i] !== null) {
			node.right = new TreeNode(arr[i]!);
			queue.push(node.right);
		}
		i++;
	}

	return root;
}

function treeToArray(root: TreeNode | null): (number | null)[] {
	if (!root) return [];

	const queue: (TreeNode | null)[] = [root];
	const level: (number | null)[] = [];

	while (queue.length) {
		const node = queue.shift()!;

		if (node !== null) {
			level.push(node.val);
			queue.push(node.left);
			queue.push(node.right);
		} else {
			level.push(null);
		}
	}

	// Remove trailing nulls
	while (level.length && level[level.length - 1] === null) {
		level.pop();
	}

	return level;
}

export { TreeNode, arrayToTree, treeToArray };
