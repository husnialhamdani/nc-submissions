func spiralOrder(matrix [][]int) []int {
    
    /*
        00  01  02
        1   2   3   
        
        10  11  12
        4   5   6   
        
        20  21  22
        7   8   9 
    
        1, 2, 3, 6, 9, 8
    */
    m := len(matrix)
    n := len(matrix[0])

    top, bottom := 0, m-1
    left, right := 0, n-1

    result := make([]int, 0, m*n)

    for top <= bottom && left <= right{
        
        for i:=left; i<=right; i++{
            result = append(result, matrix[top][i])
        }
        top++

        for j:=top; j<=bottom; j++{
            result = append(result, matrix[j][right])
        }
        right--

        if top <= bottom{
            for i:=right; i>=left; i--{
                result = append(result, matrix[bottom][i])
            }
            bottom--
        }

        if left <= right{
            for j:=bottom; j>=top; j--{
                result = append(result, matrix[j][left])
            }
            left++
        }

    }

    return result
}
