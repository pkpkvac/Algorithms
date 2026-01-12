package copyrandompointer

import (
	"algorithms/linked_list/common"
)

func copyRandomList(head *common.ListNodeWithRandom) *common.ListNodeWithRandom {

	// head =: [[val, random_index], [val, random_index] ... ]
	// KEY POINT it is hm[OLDNODE][NEWNODE]
	hm := make(map[*common.ListNodeWithRandom]*common.ListNodeWithRandom)
	// iterate head, create a new node

	curr := head
	for curr != nil {

		newNode := common.ListNodeWithRandom{Val: curr.Val}

		hm[curr] = &newNode
		// hm[OLDNODE] = NEWNODE

		curr = curr.Next
	}

	curr = head
	for curr != nil {
		// now we need to assign the random and next pointers
		newNode := hm[curr]

		newNode.Next = hm[curr.Next]

		newNode.Random = hm[curr.Random]

		curr = curr.Next
	}

	return hm[head]
}
