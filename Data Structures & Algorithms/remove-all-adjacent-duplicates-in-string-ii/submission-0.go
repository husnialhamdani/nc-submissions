type Pair struct{
    char rune
    count int
}

func removeDuplicates(s string, k int) string {
    stack := []Pair{}

    for _, char := range s{
        if len(stack)>0 && stack[len(stack)-1].char == char{

            stack[len(stack)-1].count++

            if stack[len(stack)-1].count == k{
                stack = stack[:len(stack)-1]
            }
        }else{
            stack = append(stack, Pair{char, 1})
        }
    }

    var result strings.Builder
    for _, p := range stack{
        for i:=0; i<p.count; i++{
            result.WriteRune(p.char)
        }
    }

    return result.String()
}
