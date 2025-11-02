package islandperimeter

import "fmt"

func islandPerimeter(grid [][]int) int {

	perimeter := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {

				// check 4 directions
				for _, direction := range directions {

					newRow := i + direction[0]
					newCol := j + direction[1]

					if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || grid[newRow][newCol] == 0 {
						perimeter++
					}
				}
			}
		}
	}

	fmt.Println("perim", perimeter)

	return perimeter
}
