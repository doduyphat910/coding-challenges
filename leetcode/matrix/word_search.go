package graph

func exist(board [][]byte, word string) bool {
	var (
		rows    = len(board)
		columns = len(board[0])
		visited = make([][]bool, rows)
	)
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] == word[0] && search(board, i, j, 0, word, visited) {
				return true
			}
		}
	}

	return false
}

func search(board [][]byte, i, j, index int, word string, visited [][]bool) bool {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 ||
		visited[i][j] || board[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}

	visited[i][j] = true
	if search(board, i+1, j, index+1, word, visited) ||
		search(board, i-1, j, index+1, word, visited) ||
		search(board, i, j+1, index+1, word, visited) ||
		search(board, i, j-1, index+1, word, visited) {
		return true
	}
	visited[i][j] = false

	return false
}
