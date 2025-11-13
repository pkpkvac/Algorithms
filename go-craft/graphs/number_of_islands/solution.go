package numberofislands

func numIslandsRecursive(grid [][]byte) int {

	islands := 0

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == '1' {
				islands++
				dfs(grid, i, j)
			}
		}
	}

	return islands
}

func numIslandsIterative(grid [][]byte) int {
	islands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				islands++
				dfsIterative(grid, i, j)
			}
		}
	}

	return islands
}

// Recursive DFS approach
func dfs(grid [][]byte, start_row int, start_col int) {
	// job of this function is to zero the island if it's land,
	// call in all 4 directions

	if grid[start_row][start_col] == '0' {
		// this is water, just return
		return
	}
	// mark current cell as visited (water)
	grid[start_row][start_col] = '0'

	// check 4 directions
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, direction := range directions {

		newRow := start_row + direction[0]
		newCol := start_col + direction[1]

		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || grid[newRow][newCol] == '0' {
			// nothing to do, out of range, or water
			continue
		}

		dfs(grid, newRow, newCol)
	}

}

// Iterative DFS approach using a stack
func dfsIterative(grid [][]byte, start_row int, start_col int) {
	stack := [][]int{{start_row, start_col}}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(stack) > 0 {
		// pop from stack
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		row, col := cell[0], cell[1]

		// skip if already visited or out of bounds
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' {
			continue
		}

		// mark as visited
		grid[row][col] = '0'

		// push all valid neighbors onto stack
		for _, direction := range directions {
			newRow := row + direction[0]
			newCol := col + direction[1]

			if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] == '1' {
				stack = append(stack, []int{newRow, newCol})
			}
		}
	}
}
