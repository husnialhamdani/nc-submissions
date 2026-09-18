func orangesRotting(grid [][]int) int {
    rows := len(grid)
	cols := len(grid[0])

	queue := make([][2]int, 0)

	fresh := 0

	for r:=0; r<rows; r++{
		for c:=0; c<cols; c++{
			if grid[r][c]==2 {
				queue = append(queue, [2]int{r,c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	minutes :=0

	directions := [][2]int{
		{-1, 0},
		{1,0},
		{0, -1},
		{0, 1},
	}
	
	for len(queue) > 0 && fresh > 0 {
		levelSize := len(queue)

		// Process everything that is rotten
        // at the beginning of this minute.
        for i := 0; i < levelSize; i++ {
            r := queue[0][0]
            c := queue[0][1]
            queue = queue[1:]

            for _, dir := range directions {
                nr := r + dir[0]
                nc := c + dir[1]

                // Out of bounds
                if nr < 0 || nr >= rows ||
                    nc < 0 || nc >= cols {
                    continue
                }

                // Only fresh oranges can become rotten.
                if grid[nr][nc] != 1 {
                    continue
                }

                // Fresh -> rotten
                grid[nr][nc] = 2
                fresh--

                // It will spread rot during the next minute.
                queue = append(queue, [2]int{nr, nc})
            }
	   }

	   minutes++
	}

	if fresh > 0 {
		return -1
	}

	return minutes
}
