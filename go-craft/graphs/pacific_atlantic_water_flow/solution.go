package pacific_atlantic_water_flow

import "sort"

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])

	// Sets to track which cells each ocean can reach
	pacific := make(map[[2]int]bool)
	atlantic := make(map[[2]int]bool)

	// BFS from Pacific borders (top row + left column)
	queue := [][]int{}

	// Add top row to Pacific
	for c := 0; c < n; c++ {
		queue = append(queue, []int{0, c})
		pacific[[2]int{0, c}] = true
	}

	// Add left column to Pacific
	for r := 1; r < m; r++ {
		queue = append(queue, []int{r, 0})
		pacific[[2]int{r, 0}] = true
	}

	// BFS to find all cells Pacific can reach
	bfs(heights, queue, pacific)

	// BFS from Atlantic borders (bottom row + right column)
	queue = [][]int{}

	// Add bottom row to Atlantic
	for c := 0; c < n; c++ {
		queue = append(queue, []int{m - 1, c})
		atlantic[[2]int{m - 1, c}] = true
	}

	// Add right column to Atlantic
	for r := 0; r < m-1; r++ {
		queue = append(queue, []int{r, n - 1})
		atlantic[[2]int{r, n - 1}] = true
	}

	// BFS to find all cells Atlantic can reach
	bfs(heights, queue, atlantic)

	// Find intersection: cells both oceans can reach
	result := [][]int{}
	for coord := range pacific {
		if atlantic[coord] {
			result = append(result, []int{coord[0], coord[1]})
		}
	}

	// Sort results for consistent ordering
	sort.Slice(result, func(i, j int) bool {
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		}
		return result[i][1] < result[j][1]
	})

	return result
}

func bfs(heights [][]int, queue [][]int, reachable map[[2]int]bool) {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(heights), len(heights[0])

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		r, c := curr[0], curr[1]

		// Explore neighbors
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]

			// Check bounds
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			coord := [2]int{nr, nc}

			// If not visited AND height allows flow (water can flow "up")
			if !reachable[coord] && heights[nr][nc] >= heights[r][c] {
				reachable[coord] = true
				queue = append(queue, []int{nr, nc})
			}
		}
	}
}
