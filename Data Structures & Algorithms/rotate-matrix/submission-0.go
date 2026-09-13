func rotate(matrix [][]int)  {
    
    n := len(matrix)
    //transpose
    for i:=0; i<n; i++{
        for j:=i+1; j<n; j++{
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    //reverse
    for i:=0; i<n; i++{
        for j:=0; j<n/2; j++{
            matrix[i][j], matrix[i][n-1-j] =  matrix[i][n-1-j], matrix[i][j]
        }
    }

    /*
            [1,2,3],
            [4,5,6],
            [7,8,9]

            00  01  02
            1   4   7
            
            2   5   8
            
            3   6   9

    */

}
