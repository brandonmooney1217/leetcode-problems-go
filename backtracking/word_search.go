package backtracking

// LeetCode #79 (Medium): https://leetcode.com/problems/word-search/

func exist(board [][]byte, word string) bool {

	rows, cols := len(board), len(board[0])
	wordBytes := []byte(word)

	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		if index == len(wordBytes) {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}

		if wordBytes[index] != board[r][c] {
			return false
		}

		tmp := board[r][c]
		board[r][c] = '#'

		left := dfs(r+1, c, index+1)
		right := dfs(r-1, c, index+1)
		up := dfs(r, c+1, index+1)
		down := dfs(r, c-1, index+1)
		board[r][c] = tmp

		return left || right || up || down

	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tmp := dfs(i, j, 0)
			if tmp {
				return true
			}

		}
	}
	return false
}
