import { ListNode, createLinkedList, printLinkedList } from "./definition";

function reverseListIterativeNAIVE(head: ListNode | null): ListNode | null {
	if (!head) return null;

	let newHead = new ListNode();

	let node = new ListNode();

	// operate on old list
	let curr = head;
	let last = head;

	while (last.next != null) last = last.next;

	newHead.val = last.val;
	node = newHead;

	while (curr != last) {
		if (curr.next == last) {
			last = curr;
			const newNode = new ListNode(last.val);
			node.next = newNode;
			node = newNode;

			// restart the process
			curr = head;
		} else {
			if (curr.next) curr = curr.next;
		}
	}

	return newHead;
}

function reverseListIterative(head: ListNode | null): ListNode | null {
	// Handle empty list or single node
	if (!head || !head.next) return head;

	let prev: ListNode | null = null;
	let current: ListNode | null = head;
	let next: ListNode | null = null;

	// Iterate through the list
	while (current !== null) {
		// Save the next node
		next = current.next;

		// Reverse the pointer
		current.next = prev;

		// Move prev and current one step forward
		prev = current;
		current = next;
	}

	// prev is the new head
	return prev;
}

function reverseListRecursive(head: ListNode | null): ListNode | null {
	// Base case: empty / single node
	if (!head || !head.next) return head;

	// recursive case: reverse rest of list after head
	const newHead = reverseListRecursive(head.next);

	// reverse pointer b/w head and head.next
	head.next.next = head;
	head.next = null;

	return newHead;
}

const head = [0, 1, 2, 3];

const list: ListNode | null = createLinkedList(head);

// console.log(printLinkedList(list));

// const newHead = reverseListIterativeNAIVE(list);
// const newHead = reverseListIterative(list);
const newHead = reverseListRecursive(list);

console.log(printLinkedList(newHead));
