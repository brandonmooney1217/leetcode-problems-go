package dfs

// LeetCode #200 (Medium): https://leetcode.com/problems/number-of-islands/description/

func numIslands(grid [][]byte) int {

	rows, cols := len(grid), len(grid[0])
	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if grid[r][c] != '1' {
			return
		}
		grid[r][c] = '2'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}
