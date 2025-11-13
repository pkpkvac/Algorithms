package minimumstack

// You should aim for a solution with O(1) time for
//  each function call and O(n) space, where n is
//  the maximum number of elements present in the stack.

// Hint 1
// A brute force solution would be
// to always check for the minimum element
//  in the stack for the getMin() function call.
//  This would be an O(n) appraoch. Can you think of a better way?
// Maybe O(n) extra space to store some information.

// Hint 2
// We can use a stack to maintain the elements.
// But how can we find the minimum element at any given time?
// Perhaps we should consider a prefix approach.

// Hint 3
// We use an additional stack to maintain
//  the prefix minimum element.
// When popping elements from the main stack,
// we should also pop from this extra stack.
// However, when pushing onto the extra stack,
// we should push the minimum of the top element of the
// extra stack and the current element onto this extra stack.

type MinStack struct {
	s  []int
	ms []int
}

func Constructor() MinStack {

	return MinStack{
		s:  []int{},
		ms: []int{},
	}

}

func (ms *MinStack) Push(val int) {

	ms.s = append(ms.s, val)

	if len(ms.ms) == 0 {
		ms.ms = append(ms.ms, val)
	} else {
		min := ms.ms[len(ms.ms)-1]
		if val < min {
			ms.ms = append(ms.ms, val)
		} else {
			ms.ms = append(ms.ms, min)
		}
	}

}

func (ms *MinStack) Pop() {
	ms.s = ms.s[:len(ms.s)-1]
	ms.ms = ms.ms[:len(ms.ms)-1]

}

func (ms *MinStack) Top() int {

	return ms.s[len(ms.s)-1]

}

func (ms *MinStack) GetMin() int {

	return ms.ms[len(ms.ms)-1]

}
