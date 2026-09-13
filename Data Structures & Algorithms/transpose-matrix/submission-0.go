func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
    result := make([][]int, cols)
    
    for i:= range result {
        result[i] = make([]int, rows)
    }

    for r:=0; r<rows; r++{
        for c:=0; c<cols; c++{
            result[c][r] = matrix[r][c]
        }
    }

    return result
}
