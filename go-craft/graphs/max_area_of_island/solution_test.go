package maxareaofisland

import (
	"testing"
)

func TestNumIslands(t *testing.T) {

	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "leetcode example 1",
			grid: [][]int{
				{0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1},
				{0, 1, 1, 0, 1},
				{0, 1, 0, 0, 1},
			},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the grid for recursive test
			gridCopy1 := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy1[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy1[i], tt.grid[i])
			}
			result := maxAreaOfIslandRecursive(gridCopy1)
			if result != tt.expected {
				t.Errorf("maxAreaOfIsland(%v) = %v, want %v", tt.grid, result, tt.expected)
			}

			// Make a copy of the grid for iterative test
			gridCopy2 := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy2[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy2[i], tt.grid[i])
			}
			result = maxAreaOfIslandIterative(gridCopy2)
			if result != tt.expected {
				t.Errorf("numIslandsIterative(%v) = %v, want %v", tt.grid, result, tt.expected)
			}
		})
	}
}

// func BenchmarkNumIslandsRecursive(b *testing.B) {
// 	originalGrid := [][]byte{{'1', '1', '0', '0'},
// 		{'1', '0', '0', '0'},
// 		{'1', '1', '1', '0'},
// 		{'0', '0', '1', '1'}}
// 	b.ResetTimer()
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		// Make a copy for each iteration since the function modifies the grid
// 		grid := make([][]byte, len(originalGrid))
// 		for j := range originalGrid {
// 			grid[j] = make([]byte, len(originalGrid[j]))
// 			copy(grid[j], originalGrid[j])
// 		}
// 		numIslandsRecursive(grid)
// 	}
// }

// func BenchmarkNumIslandsIterative(b *testing.B) {
// 	originalGrid := [][]byte{{'1', '1', '0', '0'},
// 		{'1', '0', '0', '0'},
// 		{'1', '1', '1', '0'},
// 		{'0', '0', '1', '1'}}
// 	b.ResetTimer()
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		// Make a copy for each iteration since the function modifies the grid
// 		grid := make([][]byte, len(originalGrid))
// 		for j := range originalGrid {
// 			grid[j] = make([]byte, len(originalGrid[j]))
// 			copy(grid[j], originalGrid[j])
// 		}
// 		numIslandsIterative(grid)
// 	}
// }
