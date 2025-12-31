package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {

	// build adjacency list (graph)
	adjList := make(map[int][]int)

	for _, prereq := range prerequisites {
		course := prereq[0]
		prerequisite := prereq[1]
		adjList[prerequisite] = append(adjList[prerequisite], course)
	}

	visited := make(map[int]bool)
	visiting := make(map[int]bool)

	// initialize state tracking

	for course := 0; course < numCourses; course++ {

		if !visited[course] {
			if dfsCycle(course, adjList, visited, visiting) {
				return false
			}
		}
	}

	return true
}

func dfsCycle(course int, adjList map[int][]int, visited map[int]bool, visiting map[int]bool) bool {

	// base case
	if visiting[course] {
		// cycle exists
		return true
	}

	if visited[course] {
		// already checked, no cycle
		return false
	}

	neighbors := adjList[course]

	visiting[course] = true
	// visit all neighbors of current course
	for _, neighbor := range neighbors {
		if dfsCycle(neighbor, adjList, visited, visiting) {
			return true // Exists a cycle
		}
	}

	visiting[course] = false
	visited[course] = true

	return false
}
