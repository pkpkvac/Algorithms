package word_search

func exist(board [][]byte, word string) bool {
  // Try starting from every cell
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if dfs(board, word, i, j, 0) {
                return true
            }
        }
    }
    return false
}
func dfs(board [][]byte, word string, row, col, index int) bool {
    // Matched entire word
    if index == len(word) {
        return true
    }
    
    // Out of bounds or wrong character
    if row < 0 || row >= len(board) || 
       col < 0 || col >= len(board[0]) || 
       board[row][col] != word[index] {
        return false
    }
    
    // DO: Mark this cell as visited
    temp := board[row][col]
    board[row][col] = '#'  // Mark visited
    
    // EXPLORE: Try all 4 directions (up, down, left, right)
    found := dfs(board, word, row+1, col, index+1) ||  // down
             dfs(board, word, row-1, col, index+1) ||  // up
             dfs(board, word, row, col+1, index+1) ||  // right
             dfs(board, word, row, col-1, index+1)     // left
    
    // UNDO: Restore the cell (backtrack)
    board[row][col] = temp
    
    return found
}