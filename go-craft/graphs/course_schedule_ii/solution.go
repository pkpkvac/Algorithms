package courseschedule_ii

// Using KAHN's algorithm.

func findOrder(numCourses int, prerequisites [][]int) []int {

	res := []int{}
	q := []int{}
	// queue of courses that can be taken first
	adjList := make(map[int][]int)
	inDegrees := make(map[int]int)

	for _, prereq := range prerequisites {
		course := prereq[0]
		prerequisite := prereq[1]
		adjList[prerequisite] = append(adjList[prerequisite], course)
		inDegrees[course]++
	}

	for course := 0; course < numCourses; course++ {
		// this means this course can be first in order,
		// these courses do not depend on other courses
		// to start
		if inDegrees[course] == 0 { // Only add if no prerequisites!
			q = append(q, course)
		}
	}

	for len(q) > 0 {

		course := q[0]
		q = q[1:]

		res = append(res, course)
		// for this course, since you're completing it
		// reduce the indegree for all courses that depend on this one
		// from the adjaceny list
		// updatedList := []int{}
		for _, dependency := range adjList[course] {

			inDegrees[dependency]--
			if inDegrees[dependency] == 0 {
				// after we decrement, will be zero, can now add to queue
				q = append(q, dependency)
			}

		}

	}

	if len(res) == numCourses {
		return res
	}

	return []int{}

}
