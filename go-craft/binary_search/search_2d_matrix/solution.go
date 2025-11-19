package search2dmatrix

import "fmt"

func Search2DMatrix(matrix [][]int, target int) bool {

	var searchSpace []int

	i := 0

	for i < len(matrix) {
		if matrix[i][0] > target {
			return false
		}

		if matrix[i][len(matrix[i])-1] < target {
			i++
		} else {
			searchSpace = matrix[i]
			break
		}
	}

	fmt.Println(searchSpace)
	if len(searchSpace) > 0 {

		resSearch := BinarySearch(searchSpace, target)

		if resSearch != -1 {
			return true
		}

	}
	return false
}

func BinarySearch(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	// fmt.Println("target ", target)

	for l <= r {

		mid := l + (r-l)/2

		// fmt.Println("mid ", mid, " nums[mid] ", nums[mid])

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			// in the left
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return -1
}

// Alternative: O(log(m*n)) solution
// Treats the matrix as a flat sorted array and uses a single binary search
func Search2DMatrixOptimized(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)    // number of rows
	n := len(matrix[0]) // number of columns

	// Treat the matrix as a flat sorted array
	// Binary search on indices 0 to (m*n - 1)
	left := 0
	right := m*n - 1

	for left <= right {
		mid := left + (right-left)/2

		// Convert 1D index to 2D coordinates
		// Example: For a 3x4 matrix (3 rows, 4 cols), n=4
		// Index 0 -> row 0/4=0, col 0%4=0 -> matrix[0][0]
		// Index 5 -> row 5/4=1, col 5%4=1 -> matrix[1][1]
		// Index 10 -> row 10/4=2, col 10%4=2 -> matrix[2][2]
		//
		// Division (/) gives us which row we're in
		// Modulo (%) gives us which column within that row
		row := mid / n // Integer division: how many complete rows we've passed
		col := mid % n // Remainder: position within the current row
		value := matrix[row][col]

		if value == target {
			return true
		} else if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
