func isValidSudoku(board [][]byte) bool {

	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool

	for i:=0; i < 9; i++{
		for j:=0; j < 9; j++{
			val := board[i][j]

			if val == '.' {
				continue
			}

			digit := val - '1'
			boxIndex := (i/3)*3 + (j/3)

			if rows[i][digit] || cols[j][digit] || boxes[boxIndex][digit] {
				return false
			}

			rows[i][digit]=true
			cols[j][digit]=true
			boxes[boxIndex][digit]=true
		}
	}

	return true
}
