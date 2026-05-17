package graphs

// LeetCode #733 (Easy): https://leetcode.com/problems/flood-fill/

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	rows, cols := len(image), len(image[0])
	startColor := image[sr][sc]

	if startColor == color {
		return image
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if image[r][c] != startColor {
			return
		}

		image[r][c] = color
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	dfs(sr, sc)
	return image

}
