func islandPerimeter(grid [][]int) int {
	result := 0
	rows := len(grid)
	cols := len(grid[0])

	for r:=0; r<rows; r++{
		for c:=0; c<cols; c++{
			if grid[r][c] == 1{
				result += 4
				if r > 0 && grid[r-1][c] == 1 {
					result -= 2
				}
				if c > 0 && grid[r][c-1] == 1 {
					result -= 2
				}
			}
		}
	}
	return result
}
