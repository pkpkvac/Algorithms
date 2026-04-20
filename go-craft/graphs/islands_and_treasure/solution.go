package islands_and_treasure

import "math"

// type gridItem struct {
// 	x int
// 	y int
// }

func islandsAndTreasure(grid [][]int) {
	queue := [][]int{}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right
	land := math.MaxInt32

	// first pass, just take all the treasure, add it to the queue
	for r := 0; r < len(grid); r++ {

		for c := 0; c < len(grid[0]); c++ {

			if grid[r][c] == 0 {
				queue = append(queue, []int{r, c})
			}

		}

	}

	// step 2, BFS from Queue items
	// rules:
	// pop from queue, explore neighbors
	// if neighbor is INF = unvisited
	// - set to current_distance + 1
	// - add neighbor to queue
	// if neighbor is -1, or visited (not INF), skip
	// works b/c all treasures are expanding simultaneously at distance 0
	// ==> first visit to any cell IS the shortest path (you don't need to compare)

	for len(queue) > 0 {

		curr := queue[0]
		queue = queue[1:]
		r, c := curr[0], curr[1]

		// explore the 4 neighbors
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]

			// check bounds, update IF INF
			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
				continue
			}

			if grid[nr][nc] == land {
				grid[nr][nc] = grid[r][c] + 1
				queue = append(queue, []int{nr, nc})
			}

		}

	}

}
