func setZeroes(matrix [][]int) {
   /*

            00      01      02
            1       1       1

            10      11      12
            1       0       1

            20      21      22
            1       1       1

   */ 

   m:= len(matrix)
   n:= len(matrix[0])


   firstRowZero:=false
   for j:=0; j<n; j++{
        if matrix[0][j]==0{
            firstRowZero=true
            break
        }
   }

   firstColZero:=false
   for i:=0; i<m; i++{
        if matrix[i][0]==0{
            firstColZero=true
            break
        }
   }


    for i:=1; i<m; i++{
        for j:=1; j<n; j++{
            if matrix[i][j]==0{
                matrix[i][0]=0
                matrix[0][j]=0
            }
        }
    }

    for i:=1; i<m; i++{
        for j:=1; j<n; j++{
            if matrix[i][0]==0 || matrix[0][j]==0{
                matrix[i][j]=0
            }
        }
    }

    if firstRowZero{
        for j:=0; j<n; j++{
            matrix[0][j]=0
        }
    }

    if firstColZero{
        for i:=0; i<m; i++{
            matrix[i][0]=0
        }
    }

}
