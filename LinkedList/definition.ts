export class ListNode {
	val: number;
	next: ListNode | null;
	constructor(val?: number, next?: ListNode | null) {
		this.val = val === undefined ? 0 : val;
		this.next = next === undefined ? null : next;
	}
}

export function createLinkedList(inputArr: number[]): ListNode | null {
	if (inputArr.length === 0) return null;

	// head node first element
	const head = new ListNode(inputArr[0]);

	let current = head;

	for (let i = 1; i < inputArr.length; i++) {
		const newNode = new ListNode(inputArr[i]);

		current.next = newNode;

		current = newNode;
	}

	return head;
}

export function printLinkedList(node: ListNode | null): number[] {
	const result: number[] = [];

	// Traverse the list and collect values
	while (node !== null) {
		result.push(node.val);
		node = node.next;
	}

	return result;
}
