package maxareaofisland

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAreaOfIslandRecursive(grid [][]int) int {

	maxArea := 0

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == 1 {
				islandArea := dfs(grid, i, j)
				maxArea = max(islandArea, maxArea)
			}

		}
	}

	return maxArea
}
func maxAreaOfIslandIterative(grid [][]int) int {

	maxArea := 0

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == 1 {
				islandArea := dfsIterative(grid, i, j)
				maxArea = max(islandArea, maxArea)
			}

		}
	}

	return maxArea
}

func dfs(grid [][]int, start_row int, start_col int) int {

	// check 4 directions, if they're inbound, and set them to 0 if they are 1
	if start_row < 0 || start_row >= len(grid) || start_col < 0 || start_col >= len(grid[0]) || grid[start_row][start_col] == 0 {
		return 0
	}

	grid[start_row][start_col] = 0

	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	area := 1

	for _, direction := range directions {

		newRow := start_row + direction[0]
		newCol := start_col + direction[1]

		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || grid[newRow][newCol] == 0 {
			// nothing to do, out of range, or water
			continue
		}

		area += dfs(grid, newRow, newCol)

	}

	return area
}

// // Iterative DFS approach using a stack
func dfsIterative(grid [][]int, start_row int, start_col int) int {
	stack := [][]int{{start_row, start_col}}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	area := 0

	for len(stack) > 0 {
		// pop from stack
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		row, col := cell[0], cell[1]

		// skip if already visited or out of bounds
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == 0 {
			continue
		}

		// mark as visited
		grid[row][col] = 0
		area++

		// push all valid neighbors onto stack
		for _, direction := range directions {
			newRow := row + direction[0]
			newCol := col + direction[1]

			if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] == 1 {
				stack = append(stack, []int{newRow, newCol})
			}
		}
	}

	return area
}
