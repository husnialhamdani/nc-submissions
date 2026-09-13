func numIslands(grid [][]byte) int {
    rows := len(grid)
	cols := len(grid[0])

	islands := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		// Out of bounds or water
        if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
            return
        }

		// Mark as visited
        grid[r][c] = '0'

        // Explore 4 directions
        dfs(r+1, c) // down
        dfs(r-1, c) // up
        dfs(r, c+1) // right
        dfs(r, c-1) // left
	}


	for r:=0; r<rows; r++{
		for c:=0; c<cols; c++{
			if grid[r][c] == '1'{
				islands++
			}

			dfs(r,c)
		}
	}

	return islands
}
