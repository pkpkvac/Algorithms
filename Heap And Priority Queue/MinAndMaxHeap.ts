// Heaps are often forgotten about in the coding interview

// 			parent = (index-1) /2
// 				|
// 			index
// 		 /     \
// left        right
// index*2 + 1   index*2 + 2

// the heap is implemented using an ARRAY
// it is far more compact than using a node class

// getLeftChildIndex(int parentIndex) return parentIndex * 2 + 1
// getRightChildIndex(int parentIndex) return parentIndex * 2 + 2
// getParentIndex(int childIndex) return (childIndex -1) / 2

// hasLeftChild(int index) { return getLeftChild(index) < size}
// hasRightChild(int index) { return getRightChild(index) < size}
// hasParent(int index) { return getParentIndex(index) >= 0}

class MinIntHeap {
	private capacity: number = 10;
	private size: number = 0;

	items: number[] = new Array(this.capacity).fill(0);

	private swap(indexOne: number, indexTwo: number): void {
		const temp = this.items[indexOne];
		this.items[indexOne] = this.items[indexTwo];
		this.items[indexTwo] = temp;
	}

	public getSize(): number {
		return this.size;
	}

	private ensureExtraCapacity(): void {
		if (this.size === this.capacity) {
			// Create new array with double capacity
			const newItems = new Array(this.capacity * 2).fill(0);

			// Copy old items to new array
			for (let i = 0; i < this.items.length; i++) {
				newItems[i] = this.items[i];
			}

			// Update items and capacity
			this.items = newItems;
			this.capacity *= 2;
		}
	}

	private getLeftChildIndex(parentIndex: number): number {
		return parentIndex * 2 + 1;
	}
	private getRightChildIndex(parentIndex: number): number {
		return parentIndex * 2 + 2;
	}
	private getParentIndex(childIndex: number): number {
		return (childIndex - 1) / 2;
	}

	private hasLeftChild(index: number): boolean {
		return this.getLeftChildIndex(index) < this.size;
	}

	private hasRightChild(index: number): boolean {
		return this.getRightChildIndex(index) < this.size;
	}

	private hasParent(index: number): boolean {
		return this.getParentIndex(index) >= 0;
	}

	private leftChild(index: number): number {
		return this.items[this.getLeftChildIndex(index)];
	}

	private rightChild(index: number): number {
		return this.items[this.getRightChildIndex(index)];
	}

	private parent(index: number): number {
		return this.items[this.getParentIndex(index)];
	}

	public peek(): number | undefined {
		if (this.size === 0) {
			return undefined;
		}
		return this.items[0];
	}

	public poll(): number | undefined {
		if (this.size == 0) return undefined;

		const item = this.items[0];
		this.items[0] = this.items[this.size - 1];
		this.size--;

		this.heapifyDown();

		return item;
	}

	public add(item: number) {
		this.ensureExtraCapacity();

		// add item to last spot
		this.items[this.size] = item;
		this.size++;

		this.heapifyUp();
	}

	public heapifyUp() {
		let index = this.size - 1;

		// things are out of order if the parent item is larger
		// than the current index position (for a min heap)
		while (this.hasParent(index) && this.parent(index) > this.items[index]) {
			// swap
			this.swap(this.getParentIndex(index), index);
			// then move upwards
			index = this.getParentIndex(index);
		}
	}

	public heapifyDown() {
		let index = 0;
		while (this.hasLeftChild(index)) {
			// only need to check if there's a left child, because if there's no left
			// then there's definitely no right child

			// set to smaller of the children
			let smallerChildIndex = this.getLeftChildIndex(index);
			if (
				this.hasRightChild(index) &&
				this.rightChild(index) < this.leftChild(index)
			) {
				smallerChildIndex = this.getRightChildIndex(index);
			}

			// ensure that the ordering is such that smaller
			if (this.items[index] < this.items[smallerChildIndex]) {
				// exit, because properly sorted
				break;
			} else {
				// the heap is out of order
				this.swap(index, smallerChildIndex);
				// move down to smaller child
			}
			index = smallerChildIndex;
		}
	}
}

class MaxIntHeap {
	private capacity: number = 10;
	private size: number = 0;
	items: number[] = new Array(this.capacity).fill(0);

	private swap(indexOne: number, indexTwo: number): void {
		const temp = this.items[indexOne];
		this.items[indexOne] = this.items[indexTwo];
		this.items[indexTwo] = temp;
	}

	private ensureExtraCapacity(): void {
		if (this.size === this.capacity) {
			const newItems = new Array(this.capacity * 2).fill(0);
			for (let i = 0; i < this.items.length; i++) {
				newItems[i] = this.items[i];
			}
			this.items = newItems;
			this.capacity *= 2;
		}
	}

	private getLeftChildIndex(parentIndex: number): number {
		return parentIndex * 2 + 1;
	}
	private getRightChildIndex(parentIndex: number): number {
		return parentIndex * 2 + 2;
	}
	private getParentIndex(childIndex: number): number {
		return (childIndex - 1) / 2;
	}

	private hasLeftChild(index: number): boolean {
		return this.getLeftChildIndex(index) < this.size;
	}

	private hasRightChild(index: number): boolean {
		return this.getRightChildIndex(index) < this.size;
	}

	private hasParent(index: number): boolean {
		return this.getParentIndex(index) >= 0;
	}

	private leftChild(index: number): number {
		return this.items[this.getLeftChildIndex(index)];
	}

	private rightChild(index: number): number {
		return this.items[this.getRightChildIndex(index)];
	}

	private parent(index: number): number {
		return this.items[this.getParentIndex(index)];
	}

	public peek(): number | undefined {
		if (this.size === 0) return undefined;
		return this.items[0];
	}

	public poll(): number | undefined {
		if (this.size === 0) return undefined;
		const item = this.items[0];
		this.items[0] = this.items[this.size - 1];
		this.size--;
		this.heapifyDown();
		return item;
	}

	public getSize(): number {
		return this.size;
	}

	public add(item: number): void {
		this.ensureExtraCapacity();
		this.items[this.size] = item;
		this.size++;
		this.heapifyUp();
	}

	public heapifyUp(): void {
		let index = this.size - 1;
		while (this.hasParent(index) && this.parent(index) < this.items[index]) {
			this.swap(this.getParentIndex(index), index);
			index = this.getParentIndex(index);
		}
	}

	public heapifyDown(): void {
		let index = 0;
		while (this.hasLeftChild(index)) {
			let largerChildIndex = this.getLeftChildIndex(index);
			if (
				this.hasRightChild(index) &&
				this.rightChild(index) > this.leftChild(index)
			) {
				largerChildIndex = this.getRightChildIndex(index);
			}

			if (this.items[index] > this.items[largerChildIndex]) {
				break;
			} else {
				this.swap(index, largerChildIndex);
			}
			index = largerChildIndex;
		}
	}
}

export { MinIntHeap, MaxIntHeap };
