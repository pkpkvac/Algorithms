package clonegraph

import "algorithms/graphs/common"

func cloneGraph(node *common.GraphNode) *common.GraphNode {

	if node == nil {
		return nil
	}
	// visited is for cloned nodes
	visited := map[int]*common.GraphNode{}
	// node stack is for the original node processing
	nodeStack := []*common.GraphNode{}

	nodeStack = append(nodeStack, node)

	for len(nodeStack) > 0 {

		// pop the node off the stack
		n := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		// check the node exists in visited, if not, create it
		// at its position

		_, ok := visited[n.Val]
		if !ok {
			// if the node does not exist
			// there is a process of new node creation
			new_n := common.NewGraphNode(n.Val)
			visited[n.Val] = new_n
		}

		// new node
		// there is a process attaching
		// and creating the node's neighbors
		for _, neighbor := range n.Neighbors {

			_, exists := visited[neighbor.Val]

			if !exists {
				// if DNE
				// create new clone node in visited map
				new_nei := common.NewGraphNode(neighbor.Val)
				visited[neighbor.Val] = new_nei

				// only append if it does not exist
				// push the ORIGINAL node's neighbors onto the stack for processing
				// but NOT the cloned ones
				nodeStack = append(nodeStack, neighbor)
			}

			clone := visited[n.Val]
			clonedNeighbor := visited[neighbor.Val]
			clone.Neighbors = append(clone.Neighbors, clonedNeighbor)

		}

	}

	return visited[node.Val]
}
