package rottingfruit

type Coord struct {
	Row, Col int
}

func rottingFruit(grid [][]byte) int {
	rotters := []Coord{}
	t := 0
	// check 4 directions
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for row := 0; row < len(grid); row++ {

		for col := 0; col < len(grid[row]); col++ {

			if grid[row][col] == 2 {

				c := Coord{Row: row, Col: col}

				rotters = append(rotters, c)

			}

		}
	}

	for len(rotters) > 0 {

		// the number of rotters in this round
		layerSize := len(rotters)

		for layerSize > 0 {
			layerSize--
			r := rotters[0]
			rotters = rotters[1:]

			for _, direction := range directions {

				newRow := r.Row + direction[0]
				newCol := r.Col + direction[1]

				if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || grid[newRow][newCol] == 0 {
					continue
				}

				if grid[newRow][newCol] == 1 {
					c := Coord{Row: newRow, Col: newCol}
					grid[newRow][newCol] = 2
					rotters = append(rotters, c)
				}
			}
		}

		// only count a minute when we actually spread to at least one new orange
		if len(rotters) > 0 {
			t++
		}
	}

	for row := 0; row < len(grid); row++ {

		for col := 0; col < len(grid[row]); col++ {

			if grid[row][col] == 1 {

				return -1
			}

		}
	}

	return t
}
