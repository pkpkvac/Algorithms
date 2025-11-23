package common

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func NewGraphNode(val int) *GraphNode {
	return &GraphNode{Val: val, Neighbors: []*GraphNode{}}
}

func NewGraphNodeWithNeighbors(val int, neighbors []*GraphNode) *GraphNode {
	return &GraphNode{Val: val, Neighbors: neighbors}
}

func BuildGraphFromAdjList(adjList [][]int) *GraphNode {
	if len(adjList) == 0 {
		return nil
	}

	// Create all nodes first (1-indexed, so node i has value i+1)
	nodes := make(map[int]*GraphNode)
	for i := 0; i < len(adjList); i++ {
		val := i + 1 // 1-indexed
		nodes[val] = NewGraphNode(val)
	}

	// Connect nodes based on adjacency list
	for i := 0; i < len(adjList); i++ {
		nodeVal := i + 1
		node := nodes[nodeVal]
		for _, neighborVal := range adjList[i] {
			node.Neighbors = append(node.Neighbors, nodes[neighborVal])
		}
	}

	return nodes[1] // Return node 1 (the input node)
}

// Compare graphs by structure, not memory addresses
func GraphsEqual(node1, node2 *GraphNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	visited1 := make(map[*GraphNode]*GraphNode)
	visited2 := make(map[*GraphNode]bool)

	var dfs func(*GraphNode, *GraphNode) bool
	dfs = func(n1, n2 *GraphNode) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		if len(n1.Neighbors) != len(n2.Neighbors) {
			return false
		}

		// Map original node to cloned node
		visited1[n1] = n2
		visited2[n2] = true

		// Check neighbors
		for i := range n1.Neighbors {
			n1Neighbor := n1.Neighbors[i]
			n2Neighbor := n2.Neighbors[i]

			// If we've seen n1 before, check it maps to same n2
			if mapped, exists := visited1[n1Neighbor]; exists {
				if mapped != n2Neighbor {
					return false
				}
			} else if !dfs(n1Neighbor, n2Neighbor) {
				return false
			}
		}
		return true
	}

	return dfs(node1, node2)
}
